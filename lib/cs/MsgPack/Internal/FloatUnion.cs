using System;
using System.Runtime.InteropServices;

namespace ILib.MsgPack
{
	[StructLayout(LayoutKind.Explicit)]
	struct FloatUnion
	{
		[FieldOffset(0)]
		float m_Value;
		[FieldOffset(0)]
		byte m_Byte0;
		[FieldOffset(1)]
		byte m_Byte1;
		[FieldOffset(2)]
		byte m_Byte2;
		[FieldOffset(3)]
		byte m_Byte3;

		FloatUnion(float value)
		{
			this = default(FloatUnion);
			this.m_Value = value;
		}

		internal static void Write(float value, byte[] buf, int offset)
		{
			FloatUnion conv = new FloatUnion(value);
			if (BitConverter.IsLittleEndian)
			{
				buf[offset + 0] = conv.m_Byte3;
				buf[offset + 1] = conv.m_Byte2;
				buf[offset + 2] = conv.m_Byte1;
				buf[offset + 3] = conv.m_Byte0;
			}
			else
			{
				buf[offset + 0] = conv.m_Byte0;
				buf[offset + 1] = conv.m_Byte1;
				buf[offset + 2] = conv.m_Byte2;
				buf[offset + 3] = conv.m_Byte3;
			}
		}

		internal static float Read(byte[] buf, int offset)
		{
			FloatUnion conv = default(FloatUnion);
			if (BitConverter.IsLittleEndian)
			{
				conv.m_Byte0 = buf[offset + 3];
				conv.m_Byte1 = buf[offset + 2];
				conv.m_Byte2 = buf[offset + 1];
				conv.m_Byte3 = buf[offset + 0];
			}
			else
			{
				conv.m_Byte0 = buf[offset + 0];
				conv.m_Byte1 = buf[offset + 1];
				conv.m_Byte2 = buf[offset + 2];
				conv.m_Byte3 = buf[offset + 3];
			}
			return conv.m_Value;
		}
	}
}
