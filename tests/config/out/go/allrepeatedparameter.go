// tests/proto/test.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

type AllRepeatedParameter struct {
	ValueDouble      []float64
	ValueFloat       []float32
	ValueInt32       []int32
	ValueInt64       []int64
	ValueUint32      []uint32
	ValueUint64      []uint64
	ValueSint32      []int32
	ValueSint64      []int64
	ValueFixed32     []uint32
	ValueFixed64     []uint64
	ValueSfixed32    []int32
	ValueSfixed64    []int64
	ValueBool        []bool
	ValueString      []string
	ValueBytes       [][]byte
	ValueNestMessage []*DependTest
	ValueTestEnum    []TestEnum
}

// Pack Serialize Message
func (m *AllRepeatedParameter) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *AllRepeatedParameter) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(17)
	if err != nil {
		return err
	}
	
	// Write value_double
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	if m.ValueDouble == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueDouble)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueDouble {
			err = w.WriteFloat64(m.ValueDouble[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_float
	err = w.WriteTag(2)
	if err != nil {
		return err
	}
	if m.ValueFloat == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueFloat)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueFloat {
			err = w.WriteFloat32(m.ValueFloat[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_int32
	err = w.WriteTag(3)
	if err != nil {
		return err
	}
	if m.ValueInt32 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueInt32)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueInt32 {
			err = w.WriteInt32(m.ValueInt32[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_int64
	err = w.WriteTag(4)
	if err != nil {
		return err
	}
	if m.ValueInt64 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueInt64)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueInt64 {
			err = w.WriteInt64(m.ValueInt64[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_uint32
	err = w.WriteTag(5)
	if err != nil {
		return err
	}
	if m.ValueUint32 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueUint32)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueUint32 {
			err = w.WriteUint32(m.ValueUint32[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_uint64
	err = w.WriteTag(6)
	if err != nil {
		return err
	}
	if m.ValueUint64 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueUint64)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueUint64 {
			err = w.WriteUint64(m.ValueUint64[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_sint32
	err = w.WriteTag(7)
	if err != nil {
		return err
	}
	if m.ValueSint32 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueSint32)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueSint32 {
			err = w.WriteInt32(m.ValueSint32[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_sint64
	err = w.WriteTag(8)
	if err != nil {
		return err
	}
	if m.ValueSint64 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueSint64)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueSint64 {
			err = w.WriteInt64(m.ValueSint64[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_fixed32
	err = w.WriteTag(9)
	if err != nil {
		return err
	}
	if m.ValueFixed32 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueFixed32)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueFixed32 {
			err = w.WriteUint32(m.ValueFixed32[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_fixed64
	err = w.WriteTag(10)
	if err != nil {
		return err
	}
	if m.ValueFixed64 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueFixed64)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueFixed64 {
			err = w.WriteUint64(m.ValueFixed64[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_sfixed32
	err = w.WriteTag(11)
	if err != nil {
		return err
	}
	if m.ValueSfixed32 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueSfixed32)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueSfixed32 {
			err = w.WriteInt32(m.ValueSfixed32[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_sfixed64
	err = w.WriteTag(12)
	if err != nil {
		return err
	}
	if m.ValueSfixed64 == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueSfixed64)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueSfixed64 {
			err = w.WriteInt64(m.ValueSfixed64[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_bool
	err = w.WriteTag(13)
	if err != nil {
		return err
	}
	if m.ValueBool == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueBool)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueBool {
			err = w.WriteBool(m.ValueBool[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_string
	err = w.WriteTag(14)
	if err != nil {
		return err
	}
	if m.ValueString == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueString)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueString {
			err = w.WriteString(m.ValueString[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write value_bytes
	err = w.WriteTag(15)
	if err != nil {
		return err
	}
	if m.ValueBytes == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueBytes)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueBytes {
			err = w.WriteBytes(m.ValueBytes[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write ValueNestMessage
	err = w.WriteTag(18)
	if err != nil {
		return err
	}
	if m.ValueNestMessage == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueNestMessage)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueNestMessage {
			err = w.WriteMessage(m.ValueNestMessage[arrayIndex])
			if err != nil {
				return err
			}
		}
	}
	
	// Write ValueTestEnum
	err = w.WriteTag(20)
	if err != nil {
		return err
	}
	if m.ValueTestEnum == nil{
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		arrayLen := len(m.ValueTestEnum)
		err = w.WriteArrayHeader(arrayLen)
		if err != nil {
			return err
		}
		for arrayIndex, _ := range m.ValueTestEnum {
			err = w.WriteInt32(int32(m.ValueTestEnum[arrayIndex]))
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Unpack Serialize Message
func (m *AllRepeatedParameter) Unpack(buf []byte) (error) {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *AllRepeatedParameter) Read(r protopack.Reader) error {
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
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueDouble = nil
				continue
			}
			_ValueDoubleLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueDouble = make([]float64, _ValueDoubleLen);
			for arrayIndex := uint(0); arrayIndex < _ValueDoubleLen; arrayIndex++ {
				m.ValueDouble[arrayIndex], err = r.ReadFloat64();
				if err != nil {
					return err
				}
			}
			break
		case 2: // Read value_float
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueFloat = nil
				continue
			}
			_ValueFloatLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueFloat = make([]float32, _ValueFloatLen);
			for arrayIndex := uint(0); arrayIndex < _ValueFloatLen; arrayIndex++ {
				m.ValueFloat[arrayIndex], err = r.ReadFloat32();
				if err != nil {
					return err
				}
			}
			break
		case 3: // Read value_int32
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueInt32 = nil
				continue
			}
			_ValueInt32Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueInt32 = make([]int32, _ValueInt32Len);
			for arrayIndex := uint(0); arrayIndex < _ValueInt32Len; arrayIndex++ {
				m.ValueInt32[arrayIndex], err = r.ReadInt32();
				if err != nil {
					return err
				}
			}
			break
		case 4: // Read value_int64
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueInt64 = nil
				continue
			}
			_ValueInt64Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueInt64 = make([]int64, _ValueInt64Len);
			for arrayIndex := uint(0); arrayIndex < _ValueInt64Len; arrayIndex++ {
				m.ValueInt64[arrayIndex], err = r.ReadInt64();
				if err != nil {
					return err
				}
			}
			break
		case 5: // Read value_uint32
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueUint32 = nil
				continue
			}
			_ValueUint32Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueUint32 = make([]uint32, _ValueUint32Len);
			for arrayIndex := uint(0); arrayIndex < _ValueUint32Len; arrayIndex++ {
				m.ValueUint32[arrayIndex], err = r.ReadUint32();
				if err != nil {
					return err
				}
			}
			break
		case 6: // Read value_uint64
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueUint64 = nil
				continue
			}
			_ValueUint64Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueUint64 = make([]uint64, _ValueUint64Len);
			for arrayIndex := uint(0); arrayIndex < _ValueUint64Len; arrayIndex++ {
				m.ValueUint64[arrayIndex], err = r.ReadUint64();
				if err != nil {
					return err
				}
			}
			break
		case 7: // Read value_sint32
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueSint32 = nil
				continue
			}
			_ValueSint32Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueSint32 = make([]int32, _ValueSint32Len);
			for arrayIndex := uint(0); arrayIndex < _ValueSint32Len; arrayIndex++ {
				m.ValueSint32[arrayIndex], err = r.ReadInt32();
				if err != nil {
					return err
				}
			}
			break
		case 8: // Read value_sint64
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueSint64 = nil
				continue
			}
			_ValueSint64Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueSint64 = make([]int64, _ValueSint64Len);
			for arrayIndex := uint(0); arrayIndex < _ValueSint64Len; arrayIndex++ {
				m.ValueSint64[arrayIndex], err = r.ReadInt64();
				if err != nil {
					return err
				}
			}
			break
		case 9: // Read value_fixed32
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueFixed32 = nil
				continue
			}
			_ValueFixed32Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueFixed32 = make([]uint32, _ValueFixed32Len);
			for arrayIndex := uint(0); arrayIndex < _ValueFixed32Len; arrayIndex++ {
				m.ValueFixed32[arrayIndex], err = r.ReadUint32();
				if err != nil {
					return err
				}
			}
			break
		case 10: // Read value_fixed64
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueFixed64 = nil
				continue
			}
			_ValueFixed64Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueFixed64 = make([]uint64, _ValueFixed64Len);
			for arrayIndex := uint(0); arrayIndex < _ValueFixed64Len; arrayIndex++ {
				m.ValueFixed64[arrayIndex], err = r.ReadUint64();
				if err != nil {
					return err
				}
			}
			break
		case 11: // Read value_sfixed32
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueSfixed32 = nil
				continue
			}
			_ValueSfixed32Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueSfixed32 = make([]int32, _ValueSfixed32Len);
			for arrayIndex := uint(0); arrayIndex < _ValueSfixed32Len; arrayIndex++ {
				m.ValueSfixed32[arrayIndex], err = r.ReadInt32();
				if err != nil {
					return err
				}
			}
			break
		case 12: // Read value_sfixed64
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueSfixed64 = nil
				continue
			}
			_ValueSfixed64Len, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueSfixed64 = make([]int64, _ValueSfixed64Len);
			for arrayIndex := uint(0); arrayIndex < _ValueSfixed64Len; arrayIndex++ {
				m.ValueSfixed64[arrayIndex], err = r.ReadInt64();
				if err != nil {
					return err
				}
			}
			break
		case 13: // Read value_bool
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueBool = nil
				continue
			}
			_ValueBoolLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueBool = make([]bool, _ValueBoolLen);
			for arrayIndex := uint(0); arrayIndex < _ValueBoolLen; arrayIndex++ {
				m.ValueBool[arrayIndex], err = r.ReadBool();
				if err != nil {
					return err
				}
			}
			break
		case 14: // Read value_string
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueString = nil
				continue
			}
			_ValueStringLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueString = make([]string, _ValueStringLen);
			for arrayIndex := uint(0); arrayIndex < _ValueStringLen; arrayIndex++ {
				m.ValueString[arrayIndex], err = r.ReadString();
				if err != nil {
					return err
				}
			}
			break
		case 15: // Read value_bytes
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueBytes = nil
				continue
			}
			_ValueBytesLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueBytes = make([][]byte, _ValueBytesLen);
			for arrayIndex := uint(0); arrayIndex < _ValueBytesLen; arrayIndex++ {
				m.ValueBytes[arrayIndex], err = r.ReadBytes();
				if err != nil {
					return err
				}
			}
			break
		case 18: // Read ValueNestMessage
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueNestMessage = nil
				continue
			}
			_ValueNestMessageLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueNestMessage = make([]*DependTest, _ValueNestMessageLen);
			for arrayIndex := uint(0); arrayIndex < _ValueNestMessageLen; arrayIndex++ {
				isNil, err := r.NextFormatIsNull()
				if err != nil {
					return err
				}
				if isNil {
					m.ValueNestMessage[arrayIndex] = nil
					err = r.ReadNil()
				} else {
					m.ValueNestMessage[arrayIndex] = &DependTest{}
					err = m.ValueNestMessage[arrayIndex].Read(r)
				}
				if err != nil {
					return err
				}
			}
			break
		case 20: // Read ValueTestEnum
			isArrayNil, err := r.NextFormatIsNull() 
			if err != nil {
				return err
			}
			if isArrayNil {
				r.ReadNil()
				m.ValueTestEnum = nil
				continue
			}
			_ValueTestEnumLen, err := r.ReadArrayHeader();
			if err != nil {
				return err
			}
			m.ValueTestEnum = make([]TestEnum, _ValueTestEnumLen);
			for arrayIndex := uint(0); arrayIndex < _ValueTestEnumLen; arrayIndex++ {
				val, err := r.ReadInt32()
				m.ValueTestEnum[arrayIndex] = TestEnum(val)
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
