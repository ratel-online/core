package protocol

import (
	"net"
)

type TcpReadWriteCloser struct {
	conn net.Conn
}

func NewTcpReadWriteCloser(conn net.Conn) TcpReadWriteCloser {
	return TcpReadWriteCloser{conn: conn}
}

func (t TcpReadWriteCloser) Read() (*Packet, error) {
	return Decode(t.conn)
}

func (t TcpReadWriteCloser) Write(msg Packet) error {
	_, err := t.conn.Write(Encode(msg))
	return err
}

func (t TcpReadWriteCloser) Close() error {
	return t.conn.Close()
}
