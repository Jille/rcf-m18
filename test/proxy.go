package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/Jille/go-osc/osc"
)

var ErrNoHandler = errors.New("no handler installed")

var (
	hideTypes = flag.String("hide_types", "", "Don't print for these types")
)

func toAddr(addr string) net.Addr {
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	return a
}

func shouldHide(addr string) bool {
	typ := addr[7:9]
	for _, i := range strings.Split(*hideTypes, ",") {
		if typ == i {
			return true
		}
	}
	return false
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	flag.Parse()
	c2s := &Proxy{
		Prefix: "[c2s]",
		dst:    toAddr("rcfm18.skynet.quis.cx:8000"),
	}
	s2c := &Proxy{
		Prefix: "[s2c]",
		dst:    toAddr("192.168.2.103:9000"),
	}

	pc, err := net.ListenPacket("udp", "0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
	s2c.rconn = pc
	c2s.rconn = pc
	go func() {
		log.Fatal(c2s.Serve(pc))
	}()
	log.Fatal(s2c.ListenAndServe("0.0.0.0:9000"))
}

type Proxy struct {
	Prefix string
	rconn   net.PacketConn
	dst    net.Addr
}

func (s *Proxy) ListenAndServe(addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	return s.Serve(pc)
}

func (s *Proxy) Serve(pc net.PacketConn) error {
	buf := make([]byte, 8192)
	for {
		if err := func() error {
			n, src, err := pc.ReadFrom(buf)
			if err != nil {
				return err
			}
			p, err := osc.ParsePacket(string(buf[:n]))
			if err != nil {
				return err
			}
			return s.HandlePacket(pc, p, src)
		}(); err != nil {
			log.Printf("%s Error: %v", s.Prefix, err)
		}
	}
}

func (s *Proxy) HandlePacket(pc net.PacketConn, p osc.Packet, src net.Addr) error {
	var err error
	silent := true
	switch m := p.(type) {
	case *osc.Message:
		err = s.HandleMessage(m, &silent)
	case *osc.Bundle:
		err = s.HandleBundle(m, &silent)
	default:
		return errors.New("invalid Packet")
	}
	s.Relay(p, silent)
	return err
}

func (s *Proxy) HandleBundle(b *osc.Bundle, silent *bool) error {
	// TODO(quis): handle b.Timetag
	for _, m := range b.Messages {
		if err := s.HandleMessage(m, silent); err != nil {
			return fmt.Errorf("Packet %s: %v", m, err)
		}
	}
	for _, sb := range b.Bundles {
		if err := s.HandleBundle(sb, silent); err != nil {
			return err
		}
	}
	return nil
}

func (s *Proxy) HandleMessage(m *osc.Message, silent *bool) error {
	if !shouldHide(m.Address) {
		log.Printf("%s Packet %s", s.Prefix, m)
		*silent = false
	}
	return nil
}

func (s *Proxy) Relay(p osc.Packet, silent bool) error {
	if !silent {
		log.Printf("%s [-> %s] Packet %s", s.Prefix, s.dst, p)
	}
	data, err := p.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err = s.rconn.WriteTo(data, s.dst); err != nil {
		return err
	}
	return nil
}
