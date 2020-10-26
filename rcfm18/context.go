package rcfm18

import (
	"github.com/Jille/go-osc/osc"
)

type Context struct {
	conn *Connection

	msg                    *osc.Message
	page, channel, control int
	addressType            string

	reply  *osc.Message
	broadcast bool
}

func (c *Context) Connection() *Connection {
	return c.conn
}

func (c *Context) Reply(m osc.Packet) error {
	return c.conn.Send(m)
}

func (c *Context) DontReply() {
	c.reply = nil
}

func (c *Context) DontBroadcast() {
	c.broadcast = false
}
