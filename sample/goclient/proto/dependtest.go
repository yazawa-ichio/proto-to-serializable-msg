// sample/proto/Depend.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type DependTest struct {
	Dep *Forum_PostData
}

// Pack Serialize Message
func (m *DependTest) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *DependTest) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write dep
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.Dep)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *DependTest) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *DependTest) Read(r protopack.Reader) error {
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
		case 1: // Read dep
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.Dep = nil
				err = r.ReadNil()
			} else {
				m.Dep = &Forum_PostData{}
				err = m.Dep.Read(r)
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
