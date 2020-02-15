// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

type SnakeCase struct {
	SnakeCaseValue int32
}

type SnakeCase_NestSnakeCase struct {
	NestSnakeCaseValue int32
}

// Pack Serialize Message
func (m *SnakeCase_NestSnakeCase) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *SnakeCase_NestSnakeCase) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write nest_snake_case_value
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.NestSnakeCaseValue)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *SnakeCase_NestSnakeCase) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *SnakeCase_NestSnakeCase) Read(r protopack.Reader) error {
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
		case 1: // Read nest_snake_case_value
			m.NestSnakeCaseValue, err = r.ReadInt32();
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

// Pack Serialize Message
func (m *SnakeCase) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *SnakeCase) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}
	
	// Write snake_case_value
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.SnakeCaseValue)
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *SnakeCase) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *SnakeCase) Read(r protopack.Reader) error {
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
		case 1: // Read snake_case_value
			m.SnakeCaseValue, err = r.ReadInt32();
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
