// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

// LowerCamelCase  lowerCamelCase comment
type LowerCamelCase struct {
	LowerCamelCaseField int32
}

// Pack Serialize Message
func (m *LowerCamelCase) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *LowerCamelCase) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write lowerCamelCaseField
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.LowerCamelCaseField)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *LowerCamelCase) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *LowerCamelCase) Read(r protopack.Reader) error {
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
		case 1: // Read lowerCamelCaseField
			m.LowerCamelCaseField, err = r.ReadInt32();
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
