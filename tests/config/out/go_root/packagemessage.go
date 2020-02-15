// tests/proto/import.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
	mypackage "github.com/yazawa-ichio/proto-to-serializable-msg/tests/config/out/go_root/mypackage"
)

type PackageMessage struct {
	Message *mypackage.MyMessage
	MyEnum  mypackage.MyEnum
}

// Pack Serialize Message
func (m *PackageMessage) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *PackageMessage) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(2)
	if err != nil {
		return err
	}
	
	// Write message
	err = w.WriteTag(100)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.Message)
	if err != nil {
		return err
	}
	
	// Write myEnum
	err = w.WriteTag(101)
	if err != nil {
		return err
	}
	err = w.WriteInt32(int32(m.MyEnum))
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *PackageMessage) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *PackageMessage) Read(r protopack.Reader) error {
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
		case 100: // Read message
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.Message = nil
				err = r.ReadNil()
			} else {
				m.Message = &mypackage.MyMessage{}
				err = m.Message.Read(r)
			}
			if err != nil {
				return err
			}
			break
		case 101: // Read myEnum
			val, err := r.ReadInt32()
			m.MyEnum = mypackage.MyEnum(val)
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
