package network

import (
	"github.com/ratel-online/core/protocol"
	"sync/atomic"
)

var connId int64

type Conn struct {
	id    int64
	state int
	conn  protocol.ReadWriteCloser
}

func Wrapper(conn protocol.ReadWriteCloser) *Conn {
	return &Conn{
		id:   atomic.AddInt64(&connId, 1),
		conn: conn,
	}
}

func (c *Conn) ID() int64 {
	return c.id
}

func (c *Conn) Close() error {
	c.state = 1
	return c.conn.Close()
}

func (c *Conn) State() int {
	return c.state
}

func (c *Conn) Accept(apply func(msg protocol.Packet, c *Conn)) error {
	for {
		packet, err := c.conn.Read()
		if err != nil {
			return err
		}
		apply(*packet, c)
	}
}

func (c *Conn) Write(packet protocol.Packet) error {
	return c.conn.Write(packet)
}

func (c *Conn) Read() (*protocol.Packet, error) {
	return c.conn.Read()
}
