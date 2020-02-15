// sample/proto/Forum.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

type Forum_PostForumReq struct {
	// data Data
	Data *Forum_PostData
}

// Pack Serialize Message
func (m *Forum_PostForumReq) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_PostForumReq) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}

	// Write data
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.Data)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_PostForumReq) Unpack(buf []byte) error {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_PostForumReq) Read(r protopack.Reader) error {
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
		case 1: // Read data
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.Data = nil
				err = r.ReadNil()
			} else {
				m.Data = &Forum_PostData{}
				err = m.Data.Read(r)
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
