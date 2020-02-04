// tests/proto/import.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type DependMessage struct {
	Text string
}

// Pack Serialize Message
func (m *DependMessage) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *DependMessage) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write text
	err = w.WriteTag(500)
	if err != nil {
		return err
	}
	err = w.WriteString(m.Text)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *DependMessage) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *DependMessage) Read(r protopack.Reader) error {
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
		case 500: // Read text
			m.Text, err = r.ReadString();
			if err != nil {
				return err
			}
			break
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
