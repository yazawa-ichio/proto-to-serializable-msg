// sample/proto/Forum.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type Forum_User struct {
	ID   int32
	Name string
	Roll Forum_Roll
}

// Pack Serialize Message
func (m *Forum_User) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_User) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(3)
	if err != nil {
		return err
	}
	
	// Write Id
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.ID)
	if err != nil {
		return err
	}
	
	// Write Name
	err = w.WriteTag(2)
	if err != nil {
		return err
	}
	err = w.WriteString(m.Name)
	if err != nil {
		return err
	}
	
	// Write Roll
	err = w.WriteTag(3)
	if err != nil {
		return err
	}
	err = w.WriteInt32(int32(m.Roll))
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_User) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_User) Read(r protopack.Reader) error {
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
		case 1: // Read Id
			m.ID, err = r.ReadInt32();
			if err != nil {
				return err
			}
			break
		case 2: // Read Name
			m.Name, err = r.ReadString();
			if err != nil {
				return err
			}
			break
		case 3: // Read Roll
			val, err := r.ReadInt32()
			m.Roll = Forum_Roll(val)
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
