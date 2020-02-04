// sample/proto/Forum.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type Forum_PostData struct {
	ID      int32
	Message string
	User    *Forum_User
}

// Pack Serialize Message
func (m *Forum_PostData) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_PostData) Write(w protopack.Writer) error {
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
	
	// Write Message
	err = w.WriteTag(2)
	if err != nil {
		return err
	}
	err = w.WriteString(m.Message)
	if err != nil {
		return err
	}
	
	// Write User
	err = w.WriteTag(3)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.User)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_PostData) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_PostData) Read(r protopack.Reader) error {
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
		case 2: // Read Message
			m.Message, err = r.ReadString();
			if err != nil {
				return err
			}
			break
		case 3: // Read User
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.User = nil
				err = r.ReadNil()
			} else {
				m.User = &Forum_User{}
				err = m.User.Read(r)
			}
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
