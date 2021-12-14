package protocol

import (
	"encoding/binary"
	"errors"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/util/json"
	"io"
	"strconv"
)

var (
	lenSize = 4
)

type Packet struct {
	Body []byte `json:"data"`
}

func (p Packet) Int() (int, error) {
	v, err := strconv.ParseInt(p.String(), 10, 64)
	return int(v), err
}

func (p Packet) Int64() (int64, error) {
	v, _ := strconv.ParseInt(p.String(), 10, 64)
	return v, nil
}

func (p Packet) String() string {
	return string(p.Body)
}

func (p Packet) Unmarshal(v interface{}) error {
	return json.Unmarshal(p.Body, v)
}

func StringPacket(msg string) Packet {
	return Packet{
		Body: []byte(msg),
	}
}

func ErrorPacket(err error) Packet {
	return Packet{
		Body: []byte(err.Error()),
	}
}

func ObjectPacket(obj interface{}) Packet {
	return Packet{
		Body: json.Marshal(obj),
	}
}

type ReadWriteCloser interface {
	Read() (*Packet, error)
	Write(msg Packet) error
	Close() error
}

func readUint32(reader io.Reader) (uint32, error) {
	data := make([]byte, 4)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data), nil
}

func readUint64(reader io.Reader) (uint64, error) {
	data := make([]byte, 8)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(data), nil
}

func encode(msg Packet) []byte {
	lenBytes := make([]byte, lenSize)
	binary.BigEndian.PutUint32(lenBytes, uint32(len(msg.Body)))
	data := make([]byte, 0)
	data = append(data, lenBytes...)
	return append(data, msg.Body...)
}

func decode(r io.Reader) (*Packet, error) {
	l, err := readUint32(r)
	if err != nil {
		return nil, err
	}
	if l > consts.MaxPacketSize {
		return nil, errors.New("Overflow max packet size " + strconv.Itoa(consts.MaxPacketSize))
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
