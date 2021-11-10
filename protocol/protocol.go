package protocol

import (
	"encoding/binary"
	"io"
)

var (
	lenSize     = 4
)

type Packet struct {
	Body []byte `json:"data"`
}

type ReadWriteCloser interface {
	Read() (*Packet, error)
	Write(msg Packet) error
	Close() error
}

func ReadUint32(reader io.Reader) (uint32, error) {
	data := make([]byte, 4)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data), nil
}

func ReadUint64(reader io.Reader) (uint64, error) {
	data := make([]byte, 8)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(data), nil
}

func Encode(msg Packet) []byte {
	lenBytes := make([]byte, lenSize)
	binary.BigEndian.PutUint32(lenBytes, uint32(len(msg.Body)))
	data := make([]byte, 0)
	data = append(data, lenBytes...)
	return append(data, msg.Body...)
}

func Decode(r io.Reader) (*Packet, error) {
	l, err := ReadUint32(r)
	if err != nil {
		return nil, err
	}
	dataBytes := make([]byte, l)
	_, err = io.ReadFull(r, dataBytes)
	if err != nil {
		return nil, err
	}
	return &Packet{
		Body: dataBytes,
	}, nil
}
