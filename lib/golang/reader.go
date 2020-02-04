package protopack

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type msgReader struct {
	r             *bufio.Reader
	currentFormat byte
}

func newMsgReader(reader io.Reader) *msgReader {
	return &msgReader{
		r: bufio.NewReader(reader),
	}
}

func (r *msgReader) Reset(reader io.Reader) {
	r.r = bufio.NewReader(reader)
}

func (r *msgReader) NextFormat() (byte, error) {
	buf, err := r.r.Peek(1)
	if err != nil {
		return 0, err
	}
	return buf[0], err
}

func (r *msgReader) ReadFormat() (byte, error) {
	format, err := r.r.ReadByte()
	r.currentFormat = format
	return format, err
}

func (r *msgReader) ReadPositiveFixInt() byte {
	return byte((r.currentFormat & 0x7f))
}

func (r *msgReader) ReadUInt8() (byte, error) {
	return r.r.ReadByte()
}

func (r *msgReader) ReadUInt16() (uint16, error) {
	buf := make([]byte, 2)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buf), err
}

func (r *msgReader) ReadUInt32() (uint32, error) {
	buf := make([]byte, 4)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(buf), err
}

func (r *msgReader) ReadUInt64() (uint64, error) {
	buf := make([]byte, 8)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buf), err
}

func (r *msgReader) ReadNegativeFixInt() int8 {
	return int8(((r.currentFormat & 0x1f) - 0x20))
}

func (r *msgReader) ReadInt8() (int8, error) {
	ret, err := r.r.ReadByte()
	if err != io.EOF && err != nil {
		return 0, err
	}
	return int8(ret), err
}

func (r *msgReader) ReadInt16() (int16, error) {
	buf := make([]byte, 2)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(buf)), err
}

func (r *msgReader) ReadInt32() (int32, error) {
	buf := make([]byte, 4)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(buf)), err
}

func (r *msgReader) ReadInt64() (int64, error) {
	buf := make([]byte, 8)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(buf)), err
}

func (r *msgReader) ReadFloat32() (float32, error) {
	buf := make([]byte, 4)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(buf)), err
}

func (r *msgReader) ReadFloat64() (float64, error) {
	buf := make([]byte, 8)
	_, err := r.r.Read(buf)
	if err != io.EOF && err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(buf)), err
}

func (r *msgReader) ReadFixStr() (string, error) {
	len := byte(r.currentFormat & 0x1f)
	ret := make([]byte, len)
	_, err := r.r.Read(ret)
	if err != io.EOF && err != nil {
		return "", err
	}
	return string(ret), err
}

func (r *msgReader) ReadStr8() (string, error) {
	len, err := r.ReadUInt8()
	if err != nil {
		return "", err
	}
	ret := make([]byte, len)
	_, err = r.r.Read(ret)
	if err != io.EOF && err != nil {
		return "", err
	}
	return string(ret), err
}

func (r *msgReader) ReadStr16() (string, error) {
	len, err := r.ReadUInt16()
	if err != nil {
		return "", err
	}
	buf := make([]byte, len)
	_, err = r.r.Read(buf)
	if err != io.EOF && err != nil {
		return "", err
	}
	return string(buf), err
}

func (r *msgReader) ReadStr32() (string, error) {
	len, err := r.ReadUInt32()
	if err != nil {
		return "", err
	}
	buf := make([]byte, len)
	_, err = r.r.Read(buf)
	if err != io.EOF && err != nil {
		return "", err
	}
	return string(buf), err
}

func (r *msgReader) ReadBin8() ([]byte, error) {
	len, err := r.ReadUInt8()
	if err != nil {
		return nil, err
	}
	ret := make([]byte, len)
	_, err = r.r.Read(ret)
	if err != io.EOF && err != nil {
		return nil, err
	}
	return ret, err
}

func (r *msgReader) ReadBin16() ([]byte, error) {
	len, err := r.ReadUInt16()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, len)
	_, err = r.r.Read(buf)
	if err != io.EOF && err != nil {
		return nil, err
	}
	return buf, err
}

func (r *msgReader) ReadBin32() ([]byte, error) {
	len, err := r.ReadUInt32()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, len)
	_, err = r.r.Read(buf)
	if err != io.EOF && err != nil {
		return nil, err
	}
	return buf, err
}

func (r *msgReader) ReadArrayLength() (uint, error) {
	switch r.currentFormat {
	case formatNil:
		return 0, nil
	case formatArray16:
		ret, err := r.ReadUInt16()
		return uint(ret), err
	case formatArray32:
		ret, err := r.ReadUInt32()
		return uint(ret), err
	}
	if isFixArray(r.currentFormat) {
		return uint(r.currentFormat & 0xf), nil
	}
	return 0, fmt.Errorf("msgpack format invalid %v", r.currentFormat)
}

func (r *msgReader) ReadMapLength() (uint, error) {
	switch r.currentFormat {
	case formatMap16:
		ret, err := r.ReadUInt16()
		return uint(ret), err
	case formatMap32:
		ret, err := r.ReadUInt32()
		return uint(ret), err
	}
	if isFixMap(r.currentFormat) {
		return uint(r.currentFormat & 0xf), nil
	}
	return 0, fmt.Errorf("msgpack format invalid %v", r.currentFormat)
}

func (r *msgReader) ReadExtLength() (uint, error) {
	switch r.currentFormat {
	case formatFixExt1:
		return 1, nil
	case formatFixExt2:
		return 2, nil
	case formatFixExt4:
		return 4, nil
	case formatFixExt8:
		return 8, nil
	case formatFixExt16:
		return 16, nil
	case formatExt8:
		ret, err := r.ReadUInt8()
		return uint(ret), err
	case formatExt16:
		ret, err := r.ReadUInt16()
		return uint(ret), err
	case formatExt32:
		ret, err := r.ReadUInt32()
		return uint(ret), err
	}
	return 0, fmt.Errorf("msgpack format invalid %v", r.currentFormat)
}

func (r *msgReader) ReadExtType() (int8, error) {
	if r.currentFormat == formatUInt8 {
		ret, err := r.ReadUInt8()
		return int8(ret), err
	} else if r.currentFormat == formatInt8 {
		ret, err := r.ReadUInt8()
		return int8(ret), err
	} else if isPositiveFixInt(r.currentFormat) {
		return int8(r.ReadPositiveFixInt()), nil
	} else if isNegativeFixInt(r.currentFormat) {
		return r.ReadNegativeFixInt(), nil
	}
	return 0, fmt.Errorf("msgpack format invalid %v", r.currentFormat)
}

func (r *msgReader) seek(n int) error {
	_, err := r.r.Discard(n)
	return err
}

func (r *msgReader) Skip() error {
	format, err := r.ReadFormat()
	if err != nil {
		return err
	}
	switch format {
	case formatNil:
	case formatFalse:
	case formatTrue:
		return nil
	case formatUInt8:
	case formatInt8:
		return r.seek(1)
	case formatUInt16:
	case formatInt16:
		return r.seek(2)
	case formatUInt32:
	case formatInt32:
	case formatFloat32:
		return r.seek(4)
	case formatUInt64:
	case formatInt64:
	case formatFloat64:
		return r.seek(8)
	case formatStr8:
	case formatBin8:
		len, err := r.ReadUInt8()
		if err != nil {
			return err
		}
		return r.seek(int(len))
	case formatStr16:
	case formatBin16:
		len, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		return r.seek(int(len))
	case formatStr32:
	case formatBin32:
		len, err := r.ReadUInt32()
		if err != nil {
			return err
		}
		return r.seek(int(len))
	case formatFixExt1:
		return r.seek(2)
	case formatFixExt2:
		return r.seek(3)
	case formatFixExt4:
		return r.seek(5)
	case formatFixExt8:
		return r.seek(9)
	case formatFixExt16:
		return r.seek(17)
	case formatExt8:
		len, err := r.ReadUInt8()
		if err != nil {
			return err
		}
		return r.seek(int(len) + 1)
	case formatExt16:
		len, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		return r.seek(int(len) + 1)
	case formatExt32:
		len, err := r.ReadUInt32()
		if err != nil {
			return err
		}
		return r.seek(int(len) + 1)
	}
	if isPositiveFixInt(r.currentFormat) || isNegativeFixInt(r.currentFormat) {
		return nil
	} else if isFixStr(r.currentFormat) {
		return r.seek(int(r.currentFormat & 0x1f))
	} else if isArrayFamily(r.currentFormat) {
		len, err := r.ReadArrayLength()
		if err != nil {
			return err
		}
		for i := 0; i < int(len); i++ {
			err = r.Skip()
			if err != nil {
				return err
			}
		}
	} else if isMapFamily(r.currentFormat) {
		len, err := r.ReadMapLength()
		if err != nil {
			return err
		}
		len *= 2
		for i := 0; i < int(len); i++ {
			err = r.Skip()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
