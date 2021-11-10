package network

import "C"
import (
	"github.com/ratel-online/core/protocol"
)

type Conn struct {
	ID int64 `json:"id"`

	state   int
	conn    protocol.ReadWriteCloser
	streams []byte
}

func Wrapper(id int64, conn protocol.ReadWriteCloser) *Conn {
	return &Conn{
		ID:      id,
		conn:    conn,
		streams: make([]byte, 0),
	}
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
		msg, err := c.conn.Read()
		if err != nil {
			return err
		}
		apply(*msg, c)
	}
}

func (c *Conn) Write(msg protocol.Packet) error {
	return c.conn.Write(msg)
}

