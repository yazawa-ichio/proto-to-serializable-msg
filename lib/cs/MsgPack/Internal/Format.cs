namespace ILib.MsgPack
{
	public static class Format
	{
		public const byte PositiveFixIntMin = 0x00;
		public const byte PositiveFixIntMax = 0x7f;
		public const byte FixMapMin = 0x80;
		public const byte FixMapMax = 0x8f;
		public const byte FixArrayMin = 0x90;
		public const byte FixArrayMax = 0x9f;
		public const byte FixStrMin = 0xa0;
		public const byte FixStrMax = 0xbf;
		public const byte Nil = 0xc0;
		public const byte NeverUsed = 0xc1;
		public const byte False = 0xc2;
		public const byte True = 0xc3;
		public const byte Bin8 = 0xc4;
		public const byte Bin16 = 0xc5;
		public const byte Bin32 = 0xc6;
		public const byte Ext8 = 0xc7;
		public const byte Ext16 = 0xc8;
		public const byte Ext32 = 0xc9;
		public const byte Float32 = 0xca;
		public const byte Float64 = 0xcb;
		public const byte UInt8 = 0xcc;
		public const byte UInt16 = 0xcd;
		public const byte UInt32 = 0xce;
		public const byte UInt64 = 0xcf;
		public const byte Int8 = 0xd0;
		public const byte Int16 = 0xd1;
		public const byte Int32 = 0xd2;
		public const byte Int64 = 0xd3;
		public const byte FixExt1 = 0xd4;
		public const byte FixExt2 = 0xd5;
		public const byte FixExt4 = 0xd6;
		public const byte FixExt8 = 0xd7;
		public const byte FixExt16 = 0xd8;
		public const byte Str8 = 0xd9;
		public const byte Str16 = 0xda;
		public const byte Str32 = 0xdb;
		public const byte Array16 = 0xdc;
		public const byte Array32 = 0xdd;
		public const byte Map16 = 0xde;
		public const byte Map32 = 0xdf;
		public const byte NegativeFixIntMin = 0xe0;
		public const byte NegativeFixIntMax = 0xff;

		public static bool IsPositiveFixInt(byte format)
		{
			return Format.PositiveFixIntMax >= format && format >= Format.PositiveFixIntMin;
		}

		public static bool IsNegativeFixInt(byte format)
		{
			return Format.NegativeFixIntMax >= format && format >= Format.NegativeFixIntMin;
		}

		public static bool IsFixMap(byte format)
		{
			return Format.FixMapMax >= format && format >= Format.FixMapMin;
		}

		public static bool IsFixArray(byte format)
		{
			return Format.FixArrayMax >= format && format >= Format.FixArrayMin;
		}

		public static bool IsFixStr(byte format)
		{
			return Format.FixStrMax >= format && format >= Format.FixStrMin;
		}

		public static bool IsArrayFamily(byte format)
		{
			return IsFixArray(format) || Array16 == format || Array32 == format;
		}

		public static bool IsMapFamily(byte format)
		{
			return IsFixMap(format) || Map16 == format || Map32 == format;
		}

	}
}
