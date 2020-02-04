package protopack

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type protoReader struct {
	r *msgReader
}

func (r *protoReader) Reset(reader io.Reader) {
	r.r.Reset(reader)
}

func (r *protoReader) Skip() error {
	return r.r.Skip()
}

func (r *protoReader) missMatchError() error {
	return errors.WithStack(fmt.Errorf("protoReader fail read format miss match %v", r.r.currentFormat))
}

func (r *protoReader) NextFormatIsNull() (bool, error) {
	format, err := r.r.NextFormat()
	return format == formatNil, err
}

func (r *protoReader) ReadTag() (uint32, error) {
	return r.ReadUint32()
}

func (r *protoReader) ReadNil() error {
	format, err := r.r.ReadFormat()
	if err != nil {
		return err
	}
	if format != formatNil {
		return r.missMatchError()
	}
	return nil
}

func (r *protoReader) ReadBytes() ([]byte, error) {
	format, err := r.r.ReadFormat()
	if err != nil {
		return nil, err
	}
	switch format {
	case formatNil:
		return nil, err
	case formatBin8:
		return r.r.ReadBin8()
	case formatBin16:
		return r.r.ReadBin16()
	case formatBin32:
		return r.r.ReadBin32()
	}
	return nil, r.missMatchError()
}

func (r *protoReader) ReadBool() (bool, error) {
	format, err := r.r.ReadFormat()
	if err != nil {
		return false, err
	}
	switch format {
	case formatTrue:
		return true, nil
	case formatFalse:
		return false, nil
	}
	return false, r.missMatchError()
}

func (r *protoReader) ReadInt32() (int32, error) {
	ret, err := r.ReadInt64()
	return int32(ret), err
}

func (r *protoReader) ReadUint32() (uint32, error) {
	ret, err := r.ReadUint64()
	return uint32(ret), err
}

func (r *protoReader) ReadInt64() (int64, error) {
	format, err := r.r.ReadFormat()
	if err != nil {
		return 0, err
	}
	switch format {
	case formatFloat64:
		ret, err := r.r.ReadFloat64()
		return int64(ret), err
	case formatFloat32:
		ret, err := r.r.ReadFloat32()
		return int64(ret), err
	case formatUInt8:
		ret, err := r.r.ReadUInt8()
		return int64(ret), err
	case formatUInt16:
		ret, err := r.r.ReadUInt16()
		return int64(ret), err
	case formatUInt32:
		ret, err := r.r.ReadUInt32()
		return int64(ret), err
	case formatUInt64:
		ret, err := r.r.ReadUInt64()
		return int64(ret), err
	case formatInt8:
		ret, err := r.r.ReadInt8()
		return int64(ret), err
	case formatInt16:
		ret, err := r.r.ReadInt16()
		return int64(ret), err
	case formatInt32:
		ret, err := r.r.ReadInt32()
		return int64(ret), err
	case formatInt64:
		ret, err := r.r.ReadInt64()
		return int64(ret), err
	}
	if isPositiveFixInt(format) {
		return int64(r.r.ReadPositiveFixInt()), nil
	} else if isNegativeFixInt(format) {
		return int64(r.r.ReadNegativeFixInt()), nil
	}
	return 0, r.missMatchError()
}

func (r *protoReader) ReadUint64() (uint64, error) {
	format, err := r.r.ReadFormat()
	if err != nil {
		return 0, err
	}
	switch format {
	case formatUInt8:
		ret, err := r.r.ReadUInt8()
		return uint64(ret), err
	case formatUInt16:
		ret, err := r.r.ReadUInt16()
		return uint64(ret), err
	case formatUInt32:
		ret, err := r.r.ReadUInt32()
		return uint64(ret), err
	case formatUInt64:
		ret, err := r.r.ReadUInt64()
		return uint64(ret), err
	case formatInt8:
		ret, err := r.r.ReadInt8()
		return uint64(ret), err
	case formatInt16:
		ret, err := r.r.ReadInt16()
		return uint64(ret), err
	case formatInt32:
		ret, err := r.r.ReadInt32()
		return uint64(ret), err
	case formatInt64:
		ret, err := r.r.ReadInt64()
		return uint64(ret), err
	}
	if isPositiveFixInt(format) {
		return uint64(r.r.ReadPositiveFixInt()), nil
	}
	return 0, r.missMatchError()
}

func (r *protoReader) ReadFloat64() (float64, error) {
	format, err := r.r.NextFormat()
	if err != nil {
		return 0, err
	}
	if format == formatFloat64 {
		if _, err = r.r.ReadFormat(); err != nil {
			return 0, err
		}
		return r.r.ReadFloat64()
	} else if format == formatFloat32 {
		if _, err = r.r.ReadFormat(); err != nil {
			return 0, err
		}
		ret, err := r.r.ReadFloat32()
		return float64(ret), err
	}
	ret, err := r.ReadInt64()
	return float64(ret), err
}

func (r *protoReader) ReadFloat32() (float32, error) {
	format, err := r.r.NextFormat()
	if err != nil {
		return 0, err
	}
	if format == formatFloat32 {
		if _, err = r.r.ReadFormat(); err != nil {
			return 0, err
		}
		ret, err := r.r.ReadFloat32()
		return float32(ret), err
	}
	ret, err := r.ReadInt64()
	return float32(ret), err
}

func (r *protoReader) ReadString() (string, error) {
	format, err := r.r.ReadFormat()
	if err != nil {
		return "", err
	}

	switch format {
	case formatStr8:
		return r.r.ReadStr8()
	case formatStr16:
		return r.r.ReadStr16()
	case formatStr32:
		return r.r.ReadStr32()
	case formatNil:
		return "", err
	}
	if isFixStr(format) {
		return r.r.ReadFixStr()
	}
	return "", r.missMatchError()
}

func (r *protoReader) ReadMapHeader() (uint, error) {
	_, err := r.r.ReadFormat()
	if err != nil {
		return 0, err
	}
	return r.r.ReadMapLength()
}

func (r *protoReader) ReadArrayHeader() (uint, error) {
	_, err := r.r.ReadFormat()
	if err != nil {
		return 0, err
	}
	return r.r.ReadArrayLength()
}

func (r *protoReader) ReadMessage(msg Message) error {
	return msg.Read(r)
}
