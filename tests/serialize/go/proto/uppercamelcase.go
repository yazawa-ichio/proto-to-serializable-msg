// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

// UpperCamelCase  UpperCamelCase comment
type UpperCamelCase struct {
	UpperCamelCaseField int32
}

// Pack Serialize Message
func (m *UpperCamelCase) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *UpperCamelCase) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write UpperCamelCaseField
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.UpperCamelCaseField)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *UpperCamelCase) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *UpperCamelCase) Read(r protopack.Reader) error {
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
		case 1: // Read UpperCamelCaseField
			m.UpperCamelCaseField, err = r.ReadInt32();
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
