// sample/proto/Forum.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type Forum_ForumData struct {
	Data []*Forum_PostData
}

// Pack Serialize Message
func (m *Forum_ForumData) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_ForumData) Write(w protopack.Writer) error {
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
	if m.Data == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.Data)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.Data {
			err = w.WriteMessage(m.Data[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_ForumData) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_ForumData) Read(r protopack.Reader) error {
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
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.Data = nil
				continue
			}
			_DataLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.Data = make([]*Forum_PostData, _DataLen);
			for arrayIndex := uint(0); arrayIndex < _DataLen; arrayIndex++ {
				isNil, err := r.NextFormatIsNull()
				if err != nil {
					return err
				}
				if isNil {
					m.Data[arrayIndex] = nil
					err = r.ReadNil()
				} else {
					m.Data[arrayIndex] = &Forum_PostData{}
					err = m.Data[arrayIndex].Read(r)
				}
				if err != nil {
					return err
				}
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
