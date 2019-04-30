using System;
using System.Runtime.InteropServices;

namespace ILib.MsgPack
{
	[StructLayout(LayoutKind.Explicit)]
	struct DoubleUnion
	{
		[FieldOffset(0)]
		double m_Value;
		[FieldOffset(0)]
		byte m_Byte0;
		[FieldOffset(1)]
		byte m_Byte1;
		[FieldOffset(2)]
		byte m_Byte2;
		[FieldOffset(3)]
		byte m_Byte3;
		[FieldOffset(4)]
		byte m_Byte4;
		[FieldOffset(5)]
		byte m_Byte5;
		[FieldOffset(6)]
		byte m_Byte6;
		[FieldOffset(7)]
		byte m_Byte7;

		DoubleUnion(double value)
		{
			this = default(DoubleUnion);
			this.m_Value = value;
		}

		internal static void Write(double value, byte[] buf, int offset)
		{
			var conv = new DoubleUnion(value);
			if (BitConverter.IsLittleEndian)
			{
				buf[offset + 0] = conv.m_Byte7;
				buf[offset + 1] = conv.m_Byte6;
				buf[offset + 2] = conv.m_Byte5;
				buf[offset + 3] = conv.m_Byte4;
				buf[offset + 4] = conv.m_Byte3;
				buf[offset + 5] = conv.m_Byte2;
				buf[offset + 6] = conv.m_Byte1;
				buf[offset + 7] = conv.m_Byte0;
			}
			else
			{
				buf[offset + 0] = conv.m_Byte0;
				buf[offset + 1] = conv.m_Byte1;
				buf[offset + 2] = conv.m_Byte2;
				buf[offset + 3] = conv.m_Byte3;
				buf[offset + 4] = conv.m_Byte4;
				buf[offset + 5] = conv.m_Byte5;
				buf[offset + 6] = conv.m_Byte6;
				buf[offset + 7] = conv.m_Byte7;
			}
		}

		internal static double Read(byte[] buf, int offset)
		{
			DoubleUnion conv = default(DoubleUnion);
			if (BitConverter.IsLittleEndian)
			{
				conv.m_Byte0 = buf[offset + 7];
				conv.m_Byte1 = buf[offset + 6];
				conv.m_Byte2 = buf[offset + 5];
				conv.m_Byte3 = buf[offset + 4];
				conv.m_Byte4 = buf[offset + 3];
				conv.m_Byte5 = buf[offset + 2];
				conv.m_Byte6 = buf[offset + 1];
				conv.m_Byte7 = buf[offset + 0];
			}
			else
			{
				conv.m_Byte0 = buf[offset + 0];
				conv.m_Byte1 = buf[offset + 1];
				conv.m_Byte2 = buf[offset + 2];
				conv.m_Byte3 = buf[offset + 3];
				conv.m_Byte4 = buf[offset + 4];
				conv.m_Byte5 = buf[offset + 5];
				conv.m_Byte6 = buf[offset + 6];
				conv.m_Byte7 = buf[offset + 7];
			}
			return conv.m_Value;
		}
	}
}