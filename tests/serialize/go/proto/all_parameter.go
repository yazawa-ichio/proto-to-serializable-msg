// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/protoc-gen-msgpack/lib/golang"
)

type AllParameter struct {
	ValueDouble          float64
	ValueFloat           float32
	ValueInt32           int32
	ValueInt64           int64
	ValueUint32          uint32
	ValueUint64          uint64
	ValueSint32          int32
	ValueSint64          int64
	ValueFixed32         uint32
	ValueFixed64         uint64
	ValueSfixed32        int32
	ValueSfixed64        int64
	ValueBool            bool
	ValueString          string
	ValueBytes           []byte
	ValueMapString       map[int32]string
	ValueMapInt          map[string]int32
	ValueMessage         *EmptyMessage
	ValueMapValueMessage map[int32]*DependTest
	ValueTestEnum        TestEnum
}

// Pack Serialize Message
func (m *AllParameter) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *AllParameter) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(20)
	if err != nil {
		return err
	}
	
	// Write value_double
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	err = w.WriteFloat64(m.ValueDouble)
	if err != nil {
		return err
	}
	
	// Write value_float
	err = w.WriteTag(2)
	if err != nil {
		return err
	}
	err = w.WriteFloat32(m.ValueFloat)
	if err != nil {
		return err
	}
	
	// Write value_int32
	err = w.WriteTag(3)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.ValueInt32)
	if err != nil {
		return err
	}
	
	// Write value_int64
	err = w.WriteTag(4)
	if err != nil {
		return err
	}
	err = w.WriteInt64(m.ValueInt64)
	if err != nil {
		return err
	}
	
	// Write value_uint32
	err = w.WriteTag(5)
	if err != nil {
		return err
	}
	err = w.WriteUint32(m.ValueUint32)
	if err != nil {
		return err
	}
	
	// Write value_uint64
	err = w.WriteTag(6)
	if err != nil {
		return err
	}
	err = w.WriteUint64(m.ValueUint64)
	if err != nil {
		return err
	}
	
	// Write value_sint32
	err = w.WriteTag(7)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.ValueSint32)
	if err != nil {
		return err
	}
	
	// Write value_sint64
	err = w.WriteTag(8)
	if err != nil {
		return err
	}
	err = w.WriteInt64(m.ValueSint64)
	if err != nil {
		return err
	}
	
	// Write value_fixed32
	err = w.WriteTag(9)
	if err != nil {
		return err
	}
	err = w.WriteUint32(m.ValueFixed32)
	if err != nil {
		return err
	}
	
	// Write value_fixed64
	err = w.WriteTag(10)
	if err != nil {
		return err
	}
	err = w.WriteUint64(m.ValueFixed64)
	if err != nil {
		return err
	}
	
	// Write value_sfixed32
	err = w.WriteTag(11)
	if err != nil {
		return err
	}
	err = w.WriteInt32(m.ValueSfixed32)
	if err != nil {
		return err
	}
	
	// Write value_sfixed64
	err = w.WriteTag(12)
	if err != nil {
		return err
	}
	err = w.WriteInt64(m.ValueSfixed64)
	if err != nil {
		return err
	}
	
	// Write value_bool
	err = w.WriteTag(13)
	if err != nil {
		return err
	}
	err = w.WriteBool(m.ValueBool)
	if err != nil {
		return err
	}
	
	// Write value_string
	err = w.WriteTag(14)
	if err != nil {
		return err
	}
	err = w.WriteString(m.ValueString)
	if err != nil {
		return err
	}
	
	// Write value_bytes
	err = w.WriteTag(15)
	if err != nil {
		return err
	}
	err = w.WriteBytes(m.ValueBytes)
	if err != nil {
		return err
	}
	
	// Write value_map_string
	err = w.WriteTag(16)
	if err != nil {
		return err
	}
	if m.ValueMapString == nil {
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		mapLen := len(m.ValueMapString)
		err = w.WriteMapHeader(mapLen)
		if err != nil {
			return err
		}
		for mapValueMapStringKey, mapValueMapStringValue := range m.ValueMapString {
			err = w.WriteInt32(mapValueMapStringKey)
			if err != nil {
				return err
			}
			err = w.WriteString(mapValueMapStringValue)
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_map_int
	err = w.WriteTag(17)
	if err != nil {
		return err
	}
	if m.ValueMapInt == nil {
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		mapLen := len(m.ValueMapInt)
		err = w.WriteMapHeader(mapLen)
		if err != nil {
			return err
		}
		for mapValueMapIntKey, mapValueMapIntValue := range m.ValueMapInt {
			err = w.WriteString(mapValueMapIntKey)
			if err != nil {
				return err
			}
			err = w.WriteInt32(mapValueMapIntValue)
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_message
	err = w.WriteTag(18)
	if err != nil {
		return err
	}
	err = w.WriteMessage(m.ValueMessage)
	if err != nil {
		return err
	}
	
	// Write value_map_value_message
	err = w.WriteTag(19)
	if err != nil {
		return err
	}
	if m.ValueMapValueMessage == nil {
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		mapLen := len(m.ValueMapValueMessage)
		err = w.WriteMapHeader(mapLen)
		if err != nil {
			return err
		}
		for mapValueMapValueMessageKey, mapValueMapValueMessageValue := range m.ValueMapValueMessage {
			err = w.WriteInt32(mapValueMapValueMessageKey)
			if err != nil {
				return err
			}
			err = w.WriteMessage(mapValueMapValueMessageValue)
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_testEnum
	err = w.WriteTag(20)
	if err != nil {
		return err
	}
	err = w.WriteInt32(int32(m.ValueTestEnum))
	if err != nil {
		return err
	}
	return err
}

// Unpack Serialize Message
func (m *AllParameter) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *AllParameter) Read(r protopack.Reader) error {
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
		case 1: // Read value_double
			m.ValueDouble, err = r.ReadFloat64();
			if err != nil {
				return err
			}
			break
		case 2: // Read value_float
			m.ValueFloat, err = r.ReadFloat32();
			if err != nil {
				return err
			}
			break
		case 3: // Read value_int32
			m.ValueInt32, err = r.ReadInt32();
			if err != nil {
				return err
			}
			break
		case 4: // Read value_int64
			m.ValueInt64, err = r.ReadInt64();
			if err != nil {
				return err
			}
			break
		case 5: // Read value_uint32
			m.ValueUint32, err = r.ReadUint32();
			if err != nil {
				return err
			}
			break
		case 6: // Read value_uint64
			m.ValueUint64, err = r.ReadUint64();
			if err != nil {
				return err
			}
			break
		case 7: // Read value_sint32
			m.ValueSint32, err = r.ReadInt32();
			if err != nil {
				return err
			}
			break
		case 8: // Read value_sint64
			m.ValueSint64, err = r.ReadInt64();
			if err != nil {
				return err
			}
			break
		case 9: // Read value_fixed32
			m.ValueFixed32, err = r.ReadUint32();
			if err != nil {
				return err
			}
			break
		case 10: // Read value_fixed64
			m.ValueFixed64, err = r.ReadUint64();
			if err != nil {
				return err
			}
			break
		case 11: // Read value_sfixed32
			m.ValueSfixed32, err = r.ReadInt32();
			if err != nil {
				return err
			}
			break
		case 12: // Read value_sfixed64
			m.ValueSfixed64, err = r.ReadInt64();
			if err != nil {
				return err
			}
			break
		case 13: // Read value_bool
			m.ValueBool, err = r.ReadBool();
			if err != nil {
				return err
			}
			break
		case 14: // Read value_string
			m.ValueString, err = r.ReadString();
			if err != nil {
				return err
			}
			break
		case 15: // Read value_bytes
			m.ValueBytes, err = r.ReadBytes();
			if err != nil {
				return err
			}
			break
		case 16: // Read value_map_string
			isMapNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isMapNil {
				r.ReadNil()
				m.ValueMapString = nil
				continue
			}
			mapValueMapStringLen, err := r.ReadMapHeader();
			if err != nil {
				return err
			}
			m.ValueMapString = make(map[int32]string, mapValueMapStringLen)
			for mapIndex := uint(0); mapIndex < mapValueMapStringLen; mapIndex++ {
				var mapValueMapStringKey int32
				var mapValueMapStringValue string
				mapValueMapStringKey, err = r.ReadInt32();
				if err != nil {
					return err
				}
				mapValueMapStringValue, err = r.ReadString();
				if err != nil {
					return err
				}
				m.ValueMapString[mapValueMapStringKey] = mapValueMapStringValue;
			}
			break
		case 17: // Read value_map_int
			isMapNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isMapNil {
				r.ReadNil()
				m.ValueMapInt = nil
				continue
			}
			mapValueMapIntLen, err := r.ReadMapHeader();
			if err != nil {
				return err
			}
			m.ValueMapInt = make(map[string]int32, mapValueMapIntLen)
			for mapIndex := uint(0); mapIndex < mapValueMapIntLen; mapIndex++ {
				var mapValueMapIntKey string
				var mapValueMapIntValue int32
				mapValueMapIntKey, err = r.ReadString();
				if err != nil {
					return err
				}
				mapValueMapIntValue, err = r.ReadInt32();
				if err != nil {
					return err
				}
				m.ValueMapInt[mapValueMapIntKey] = mapValueMapIntValue;
			}
			break
		case 18: // Read value_message
			isNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isNil {
				m.ValueMessage = nil
				err = r.ReadNil()
			} else {
				m.ValueMessage = &EmptyMessage{}
				err = m.ValueMessage.Read(r)
			}
			if err != nil {
				return err
			}
			break
		case 19: // Read value_map_value_message
			isMapNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isMapNil {
				r.ReadNil()
				m.ValueMapValueMessage = nil
				continue
			}
			mapValueMapValueMessageLen, err := r.ReadMapHeader();
			if err != nil {
				return err
			}
			m.ValueMapValueMessage = make(map[int32]*DependTest, mapValueMapValueMessageLen)
			for mapIndex := uint(0); mapIndex < mapValueMapValueMessageLen; mapIndex++ {
				var mapValueMapValueMessageKey int32
				var mapValueMapValueMessageValue *DependTest
				mapValueMapValueMessageKey, err = r.ReadInt32();
				if err != nil {
					return err
				}
				isNil, err := r.NextFormatIsNull()
				if err != nil {
					return err
				}
				if isNil {
					mapValueMapValueMessageValue = nil
					err = r.ReadNil()
				} else {
					mapValueMapValueMessageValue = &DependTest{}
					err = mapValueMapValueMessageValue.Read(r)
				}
				if err != nil {
					return err
				}
				m.ValueMapValueMessage[mapValueMapValueMessageKey] = mapValueMapValueMessageValue;
			}
			break
		case 20: // Read value_testEnum
			val, err := r.ReadInt32()
			m.ValueTestEnum = TestEnum(val)
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
