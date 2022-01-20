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
	return decode(t.conn)
}

func (t TcpReadWriteCloser) Write(msg Packet) error {
	_, err := t.conn.Write(encode(msg))
	return err
}

func (t TcpReadWriteCloser) Close() error {
	return t.conn.Close()
}

func (t TcpReadWriteCloser) IP() string {
	return t.conn.RemoteAddr().String()
}
