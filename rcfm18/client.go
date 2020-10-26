package rcfm18

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/Jille/go-osc/osc"
)

func Dial(addr string) (*RCFM18Client, error) {
	dst, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	// conn, err := net.DialUDP(dst.Network(), nil, dst)
	conn, err := net.ListenUDP(dst.Network(), nil)
	if err != nil {
		return nil, err
	}
	c := &Connection{
		conn: conn,
		dst:  dst,
	}
	cl := &RCFM18Client{
		conn: c,
		Send: c.Send,
	}
	cl.setupCallbacks()
	return cl, nil
}

// TODO(quis): Find a way to cleanly deduplicate this code with server.go.

func (c *RCFM18Client) ListenAndServe() error {
	pc, err := net.ListenPacket("udp", "0.0.0.0:9000")
	if err != nil {
		return err
	}
	return c.Serve(pc)
}

func (c *RCFM18Client) Serve(pc net.PacketConn) error {
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
			return c.HandlePacket(pc, p, src)
		}(); err != nil {
			log.Printf("Error: %v", err)
		}
	}
}

func (c *RCFM18Client) HandlePacket(pc net.PacketConn, p osc.Packet, src net.Addr) error {
	switch m := p.(type) {
	case *osc.Message:
		return c.HandleMessage(m)
	case *osc.Bundle:
		return c.HandleBundle(m)
	default:
		return errors.New("invalid Packet")
	}
}

func (c *RCFM18Client) HandleBundle(b *osc.Bundle) error {
	// TODO(quis): handle b.Timetag
	for _, m := range b.Messages {
		if err := c.HandleMessage(m); err != nil {
			return fmt.Errorf("Packet %s: %v", m, err)
		}
	}
	for _, sb := range b.Bundles {
		if err := c.HandleBundle(sb); err != nil {
			return err
		}
	}
	return nil
}

func (c *RCFM18Client) HandleMessage(m *osc.Message) error {
	log.Printf("Packet %s", m)
	g := addrRe.FindStringSubmatch(m.Address)
	if g == nil {
		return fmt.Errorf("address %q didn't match pattern", m.Address)
	}
	page, err := strconv.Atoi(g[1])
	if err != nil {
		return err
	}
	channel, err := strconv.Atoi(g[2])
	if err != nil {
		return err
	}
	typ := g[3]
	control, err := strconv.Atoi(g[4])
	if err != nil {
		return err
	}
	cx := &Context{
		conn:        c.conn,
		msg:         m,
		page:        page,
		channel:     channel,
		addressType: typ,
		control:     control,
	}
	return c.Dispatch(cx)
}
