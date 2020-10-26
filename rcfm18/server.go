package rcfm18

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/Jille/go-osc/osc"
)

func (s *RCFM18Server) ListenAndServe(addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	return s.Serve(pc)
}

func (s *RCFM18Server) Serve(pc net.PacketConn) error {
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
			log.Printf("Error: %v", err)
		}
	}
}

func (s *RCFM18Server) HandlePacket(pc net.PacketConn, p osc.Packet, src net.Addr) error {
	mangled := s.MangleReturnAddress(src)
	conn, found := s.connections[mangled.String()]
	if !found {
		conn = &Connection{
			conn: pc,
			dst:  mangled,
		}
		s.connections[mangled.String()] = conn
	}
	switch m := p.(type) {
	case *osc.Message:
		return s.HandleMessage(conn, m)
	case *osc.Bundle:
		return s.HandleBundle(conn, m)
	default:
		return errors.New("invalid Packet")
	}
}

func (s *RCFM18Server) HandleBundle(conn *Connection, b *osc.Bundle) error {
	// TODO(quis): handle b.Timetag
	for _, m := range b.Messages {
		if err := s.HandleMessage(conn, m); err != nil {
			return fmt.Errorf("Packet %s: %v", m, err)
		}
	}
	for _, sb := range b.Bundles {
		if err := s.HandleBundle(conn, sb); err != nil {
			return err
		}
	}
	return nil
}

func (s *RCFM18Server) HandleMessage(conn *Connection, m *osc.Message) error {
	if m.Address == "/00/00/pl_000" {
		// TODO(quis): Remove hack.
		qr := osc.NewMessage("/00/00/vmeter_000")
		for i := 0; 41 > i; i++ {
			qr.Append(rand.Float32())
		}
		data, err := qr.MarshalBinary()
		if err != nil {
			return err
		}

		if _, err = conn.conn.WriteTo(data, conn.dst); err != nil {
			return err
		}
		return nil
	}
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
	c := &Context{
		conn:        conn,
		msg:         m,
		page:        page,
		channel:     channel,
		addressType: typ,
		control:     control,
		reply:       osc.NewMessage(m.Address),
	}
	if err := s.Dispatch(c); err != nil {
		return err
	}
	if c.broadcast {
		s.Broadcast(m, conn)
	}
	if c.reply == nil {
		return nil
	}
	return conn.Send(c.reply)
}

func (s *RCFM18Server) Broadcast(m osc.Packet, except *Connection) {
	for _, c := range s.connections {
		if c == except {
			continue
		}
		if err := c.Send(m); err != nil {
			log.Printf("Failed to broadcast to %s: %v", c.dst, err)
		}
	}
}

func (s *RCFM18Server) BroadcastMeters(v []float32) {
	switch len(v) {
	case 20, 26, 28, 31, 37, 39, 41:
		// 1-20: 18 inputs, 2 usb player
		// 21-26: FX 1-3 returns (L+R)
		// 27-28: Main (L+R)
		// 29-31: FX 1-3 sends
		// 32-37: AUX 1-6
		// 38-39: Phones out (L+R)
		// 40-41: Recording (L+R)
	default:
		panic("invalid set of meters given")
	}
	m := osc.NewMessage("/00/00/vmeter_000")
	for _, f := range v {
		m.Append(f)
	}
	s.Broadcast(m, nil)
}
