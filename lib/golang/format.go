package protopack

const (
	formatPositiveFixIntMin byte = 0x00
	formatPositiveFixIntMax byte = 0x7f
	formatFixMapMin         byte = 0x80
	formatFixMapMax         byte = 0x8f
	formatFixArrayMin       byte = 0x90
	formatFixArrayMax       byte = 0x9f
	formatFixStrMin         byte = 0xa0
	formatFixStrMax         byte = 0xbf
	formatNil               byte = 0xc0
	formatNeverUsed         byte = 0xc1
	formatFalse             byte = 0xc2
	formatTrue              byte = 0xc3
	formatBin8              byte = 0xc4
	formatBin16             byte = 0xc5
	formatBin32             byte = 0xc6
	formatExt8              byte = 0xc7
	formatExt16             byte = 0xc8
	formatExt32             byte = 0xc9
	formatFloat32           byte = 0xca
	formatFloat64           byte = 0xcb
	formatUInt8             byte = 0xcc
	formatUInt16            byte = 0xcd
	formatUInt32            byte = 0xce
	formatUInt64            byte = 0xcf
	formatInt8              byte = 0xd0
	formatInt16             byte = 0xd1
	formatInt32             byte = 0xd2
	formatInt64             byte = 0xd3
	formatFixExt1           byte = 0xd4
	formatFixExt2           byte = 0xd5
	formatFixExt4           byte = 0xd6
	formatFixExt8           byte = 0xd7
	formatFixExt16          byte = 0xd8
	formatStr8              byte = 0xd9
	formatStr16             byte = 0xda
	formatStr32             byte = 0xdb
	formatArray16           byte = 0xdc
	formatArray32           byte = 0xdd
	formatMap16             byte = 0xde
	formatMap32             byte = 0xdf
	formatNegativeFixIntMin byte = 0xe0
	formatNegativeFixIntMax byte = 0xff
)

func isPositiveFixInt(format byte) bool {
	return formatPositiveFixIntMax >= format && format >= formatPositiveFixIntMin
}

func isNegativeFixInt(format byte) bool {
	return formatNegativeFixIntMax >= format && format >= formatNegativeFixIntMin
}

func isFixMap(format byte) bool {
	return formatFixMapMax >= format && format >= formatFixMapMin
}

func isFixArray(format byte) bool {
	return formatFixArrayMax >= format && format >= formatFixArrayMin
}

func isFixStr(format byte) bool {
	return formatFixStrMax >= format && format >= formatFixStrMin
}

func isArrayFamily(format byte) bool {
	return isFixArray(format) || formatArray16 == format || formatArray32 == format
}

func isMapFamily(format byte) bool {
	return isFixMap(format) || formatMap16 == format || formatMap32 == format
}
