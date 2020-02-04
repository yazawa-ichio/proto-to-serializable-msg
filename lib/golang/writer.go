package protopack

import (
	"bytes"
	"encoding/binary"
	"math"
)

type msgWriter struct {
	buf bytes.Buffer
}

func newMsgWriter(buf bytes.Buffer) *msgWriter {
	return &msgWriter{
		buf: buf,
	}
}

func (w *msgWriter) Reset() {
	w.buf.Reset()
}

func (w *msgWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *msgWriter) writeImpl(val byte) error {
	return w.buf.WriteByte(val)
}

func (w *msgWriter) WriteNil() error {
	return w.writeImpl(formatNil)
}

func (w *msgWriter) WriteBool(value bool) error {
	if value {
		return w.writeImpl(formatTrue)
	} else {
		return w.writeImpl(formatFalse)
	}
}

func (w *msgWriter) WriteByte(value byte) error {
	if value <= math.MaxInt8 {
		return w.writePositiveFixInt(value)
	} else {
		if err := w.writeImpl(formatUInt8); err != nil {
			return err
		}
		return w.writeUInt8(value)
	}
}

func (w *msgWriter) WriteUint16(value uint16) error {
	if value <= math.MaxUint8 {
		return w.WriteByte(byte(value))
	} else {
		if err := w.writeImpl(formatUInt16); err != nil {
			return err
		}
		return w.writeUInt16(value)
	}
}

func (w *msgWriter) WriteUint32(value uint32) error {
	if value <= math.MaxUint16 {
		return w.WriteUint16(uint16(value))
	} else {
		if err := w.writeImpl(formatUInt32); err != nil {
			return err
		}
		return w.writeUInt32(value)
	}
}

func (w *msgWriter) WriteUint64(value uint64) error {
	if value <= math.MaxUint32 {
		return w.WriteUint32(uint32(value))
	} else {
		if err := w.writeImpl(formatUInt64); err != nil {
			return err
		}
		return w.writeUInt64(value)
	}
}

func (w *msgWriter) WriteInt8(value int8) error {
	if value >= 0 {
		return w.WriteByte(byte(value))
	} else if value >= -32 {
		return w.writeNegativeFixInt(value)
	} else {
		if err := w.writeImpl(formatInt8); err != nil {
			return err
		}
		return w.writeInt8(value)
	}
}

func (w *msgWriter) WriteInt16(value int16) error {
	if value >= 0 {
		return w.WriteUint16(uint16(value))
	} else if value >= math.MinInt8 {
		return w.WriteInt8(int8(value))
	} else {
		if err := w.writeImpl(formatInt16); err != nil {
			return err
		}
		return w.writeInt16(value)
	}
}

func (w *msgWriter) WriteInt32(value int32) error {
	if value >= 0 {
		return w.WriteUint32(uint32(value))
	} else if value >= math.MinInt16 {
		return w.WriteInt16(int16(value))
	} else {
		if err := w.writeImpl(formatInt32); err != nil {
			return err
		}
		return w.writeInt32(value)
	}
}

func (w *msgWriter) WriteInt64(value int64) error {
	if value >= 0 {
		return w.WriteUint64(uint64(value))
	} else if value >= math.MinInt32 {
		return w.WriteInt32(int32(value))
	} else {
		if err := w.writeImpl(formatInt64); err != nil {
			return err
		}
		return w.writeInt64(value)
	}
}

func (w *msgWriter) WriteFloat32(value float32) error {
	if err := w.writeImpl(formatFloat32); err != nil {
		return err
	}
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) WriteFloat64(value float64) error {
	if err := w.writeImpl(formatFloat64); err != nil {
		return err
	}
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) WriteString(value string) error {
	if value == "" {
		return w.WriteNil()
	}

	length := len(value)

	if length <= 31 {
		if err := w.writeImpl(byte((formatFixStrMin | byte(length)))); err != nil {
			return err
		}
	} else if length <= math.MaxUint8 {
		if err := w.writeImpl(formatStr8); err != nil {
			return err
		}
		if err := w.writeUInt8(byte(length)); err != nil {
			return err
		}
	} else if length <= math.MaxUint16 {
		if err := w.writeImpl(formatStr16); err != nil {
			return err
		}
		if err := w.writeUInt16(uint16(length)); err != nil {
			return err
		}
	} else {
		if err := w.writeImpl(formatStr32); err != nil {
			return err
		}
		if err := w.writeUInt32(uint32(length)); err != nil {
			return err
		}
	}
	_, err := w.buf.WriteString(value)
	return err
}

func (w *msgWriter) WriteBytes(bytes []byte) error {
	if bytes == nil {
		return w.WriteNil()
	}

	len := len(bytes)

	if len <= math.MaxUint8 {
		if err := w.writeImpl(formatBin8); err != nil {
			return err
		}
		if err := w.writeUInt8(byte(len)); err != nil {
			return err
		}
	} else if len <= math.MaxUint16 {
		if err := w.writeImpl(formatBin16); err != nil {
			return err
		}
		if err := w.writeUInt16(uint16(len)); err != nil {
			return err
		}
	} else {
		if err := w.writeImpl(formatBin32); err != nil {
			return err
		}
		if err := w.writeUInt32(uint32(len)); err != nil {
			return err
		}

	}

	_, err := w.buf.Write(bytes)
	return err
}

func (w *msgWriter) WriteArrayHeader(length int) error {
	if length <= 15 {
		return w.writeImpl(byte((byte(length) | formatFixArrayMin)))
	} else if length <= math.MaxUint16 {
		if err := w.writeImpl(formatArray16); err != nil {
			return err
		}
		return w.writeUInt16(uint16(length))
	} else {
		if err := w.writeImpl(formatArray32); err != nil {
			return err
		}
		return w.writeUInt32(uint32(length))
	}
}

func (w *msgWriter) WriteMapHeader(length int) error {
	if length <= 15 {
		return w.writeImpl(byte((byte(length) | formatFixMapMin)))
	} else if length <= math.MaxUint16 {
		if err := w.writeImpl(formatMap16); err != nil {
			return err
		}
		return w.writeUInt16(uint16(length))
	} else {
		if err := w.writeImpl(formatMap32); err != nil {
			return err
		}
		return w.writeUInt32(uint32(length))
	}
}

func (w *msgWriter) WriteExt(extType byte, bytes []byte) error {
	length := uint32(len(bytes))
	var err error
	if length == 1 {
		err = w.writeImpl(formatFixExt1)
	} else if length == 2 {
		err = w.writeImpl(formatFixExt2)
	} else if length == 4 {
		err = w.writeImpl(formatFixExt4)
	} else if length == 8 {
		err = w.writeImpl(formatFixExt8)
	} else if length == 16 {
		err = w.writeImpl(formatFixExt16)
	} else if length <= math.MaxUint8 {
		err = w.writeImpl(formatExt8)
		if err != nil {
			return err
		}
		err = w.writeUInt8(byte(length))
	} else if length <= math.MaxUint16 {
		err = w.writeImpl(formatExt16)
		if err != nil {
			return err
		}
		err = w.writeUInt16(uint16(length))
	} else if length <= math.MaxInt32 {
		err = w.writeImpl(formatExt32)
		if err != nil {
			return err
		}
		err = w.writeUInt32(length)
	}
	if err != nil {
		return err
	}
	err = w.writeImpl(extType)
	if err != nil {
		return err
	}
	_, err = w.buf.Write(bytes)
	return err
}

func (w *msgWriter) writePositiveFixInt(value byte) error {
	return w.writeImpl(byte((value | formatPositiveFixIntMin)))
}

func (w *msgWriter) writeUInt8(value byte) error {
	return w.writeImpl(value)
}

func (w *msgWriter) writeUInt16(value uint16) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) writeUInt32(value uint32) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) writeUInt64(value uint64) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) writeNegativeFixInt(value int8) error {
	return w.writeImpl(byte((byte(value) | formatNegativeFixIntMin)))
}

func (w *msgWriter) writeInt8(value int8) error {
	return w.writeImpl(byte(value))
}

func (w *msgWriter) writeInt16(value int16) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) writeInt32(value int32) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}

func (w *msgWriter) writeInt64(value int64) error {
	return binary.Write(&w.buf, binary.BigEndian, value)
}
