// tests/proto/depend/depend.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type DepDep_DependTestMessage struct {
	Message *PackageMessage
	DepDep  *DependMessage
}

// Pack Serialize Message
func (m *DepDep_DependTestMessage) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *DepDep_DependTestMessage) Write(w protopack.Writer) error {
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
	
	// Write dep_dep
	err = w.WriteTag(101)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.DepDep)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *DepDep_DependTestMessage) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *DepDep_DependTestMessage) Read(r protopack.Reader) error {
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
				m.Message = &PackageMessage{}
				err = m.Message.Read(r)
			}
			if err != nil {
				return err
			}
			break
		case 101: // Read dep_dep
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.DepDep = nil
				err = r.ReadNil()
			} else {
				m.DepDep = &DependMessage{}
				err = m.DepDep.Read(r)
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
