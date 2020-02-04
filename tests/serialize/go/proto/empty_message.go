// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type EmptyMessage struct {
}

// Pack Serialize Message
func (m *EmptyMessage) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *EmptyMessage) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(0)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *EmptyMessage) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *EmptyMessage) Read(r protopack.Reader) error {
	// Read Map Length
	len, err := r.ReadMapHeader()
	if err != nil {
		return err
	}

	for i := uint(0); i < len; i++ {

		// Read Tag
		tag, err := r.ReadTag()
		if err != nil {
			return err
		}

		switch tag {
		default:
			err = r.Skip()
			if err != nil {
				return err
			}
			break
		}
	}
	return err
}
