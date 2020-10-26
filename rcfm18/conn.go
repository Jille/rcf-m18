package rcfm18

import (
	"log"
	"net"

	"github.com/Jille/go-osc/osc"
)

type Connection struct {
	conn net.PacketConn
	dst  net.Addr

	OptimizedMeterSending bool
}

func (c *Connection) Send(p osc.Packet) error {
	log.Printf("[%s] Sending %s", c.dst, p)
	data, err := p.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err = c.conn.WriteTo(data, c.dst); err != nil {
		return err
	}
	return nil
}
