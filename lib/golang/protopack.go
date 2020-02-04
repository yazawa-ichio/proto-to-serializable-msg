package protopack

import (
	"bytes"
	"io"
)

type Message interface {
	Write(w Writer) error
	Read(r Reader) error
}

type Writer interface {
	Length() int
	Bytes() []byte
	Reset()
	WriteMapHeader(length int) error
	WriteArrayHeader(length int) error
	WriteTag(tag uint32) error
	WriteNil() error
	WriteFloat64(val float64) error
	WriteFloat32(val float32) error
	WriteInt64(val int64) error
	WriteUint64(val uint64) error
	WriteInt32(val int32) error
	WriteUint32(val uint32) error
	WriteBool(val bool) error
	WriteString(val string) error
	WriteBytes(val []byte) error
	WriteMessage(val Message) error
}

type Reader interface {
	Reset(r io.Reader)
	Skip() error
	NextFormatIsNull() (bool, error)
	ReadTag() (uint32, error)
	ReadNil() error
	ReadBytes() ([]byte, error)
	ReadBool() (bool, error)
	ReadInt32() (int32, error)
	ReadUint32() (uint32, error)
	ReadInt64() (int64, error)
	ReadUint64() (uint64, error)
	ReadFloat64() (float64, error)
	ReadFloat32() (float32, error)
	ReadString() (string, error)
	ReadMapHeader() (uint, error)
	ReadArrayHeader() (uint, error)
	ReadMessage(msg Message) error
}

func NewWriter(buf bytes.Buffer) Writer {
	return &protoWriter{newMsgWriter(buf)}
}

func NewReader(r io.Reader) Reader {
	return &protoReader{newMsgReader(r)}
}

func Pack(m Message) ([]byte, error) {
	w := NewWriter(bytes.Buffer{})
	if err := m.Write(w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func Unpack(m Message, buf []byte) error {
	r := NewReader(bytes.NewBuffer(buf))
	return m.Read(r)
}
