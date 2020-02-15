// tests/proto/packagetest/package.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

type MyPackage_MyMessage struct {
}

// Pack Serialize Message
func (m *MyPackage_MyMessage) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *MyPackage_MyMessage) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(0)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *MyPackage_MyMessage) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *MyPackage_MyMessage) Read(r protopack.Reader) error {
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
