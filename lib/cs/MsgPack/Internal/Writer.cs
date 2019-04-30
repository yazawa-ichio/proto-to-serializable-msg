using System.IO;
using System.Text;

namespace ILib.MsgPack
{
	public class Writer
	{
		byte[] m_Buf;
		int m_Position;

		public int Length { get { return m_Position; } }

		public Writer(int initSize = 1024)
		{
			m_Buf = new byte[initSize];
		}

		public byte[] ToArray()
		{
			byte[] buf = new byte[m_Position];
			System.Buffer.BlockCopy(m_Buf, 0, buf, 0, m_Position);
			return buf;
		}

		public System.ArraySegment<byte> ReadSegment()
		{
			return new System.ArraySegment<byte>(m_Buf, 0, m_Position);
		}

		public void Read(byte[] dst, int offest = 0)
		{
			System.Buffer.BlockCopy(m_Buf, 0, dst, offest, m_Position);
		}

		public void Clear()
		{
			m_Position = 0;
		}

		void ExpendSizeIfNeed(int size)
		{
			if (m_Buf.Length < m_Position + size)
			{
				System.Array.Resize(ref m_Buf, m_Buf.Length + 1024);
			}
		}

		void WriteImpl(byte val)
		{
			ExpendSizeIfNeed(1);
			m_Buf[m_Position++] = val;
		}

		public void WriteNil()
		{
			WriteImpl(Format.Nil);
		}

		public void Write(bool value)
		{
			WriteImpl(value ? Format.True : Format.False);
		}

		public void Write(byte value)
		{
			if (value <= sbyte.MaxValue)
			{
				WritePositiveFixInt(value);
			}
			else
			{
				WriteImpl(Format.UInt8);
				WriteUInt8(value);
			}
		}

		public void Write(ushort value)
		{
			if (value <= byte.MaxValue)
			{
				Write((byte)value);
			}
			else
			{
				WriteImpl(Format.UInt16);
				WriteUInt16(value);
			}
		}

		public void Write(uint value)
		{
			if (value <= ushort.MaxValue)
			{
				Write((ushort)value);
			}
			else
			{
				WriteImpl(Format.UInt32);
				WriteUInt32(value);
			}
		}

		public void Write(ulong value)
		{
			if (value <= uint.MaxValue)
			{
				Write((uint)value);
			}
			else
			{
				WriteImpl(Format.UInt64);
				WriteUInt64(value);
			}
		}

		public void Write(sbyte value)
		{
			if (value >= 0)
			{
				Write((byte)value);
			}
			else if (value >= -32)
			{
				WriteNegativeFixInt(value);
			}
			else
			{
				WriteImpl(Format.Int8);
				WriteInt8(value);
			}
		}

		public void Write(short value)
		{
			if (value >= 0)
			{
				Write((ushort)value);
			}
			else if (value >= sbyte.MinValue)
			{
				Write((sbyte)value);
			}
			else
			{
				WriteImpl(Format.Int16);
				WriteInt16(value);
			}
		}

		public void Write(int value)
		{
			if (value >= 0)
			{
				Write((uint)value);
			}
			else if (value >= short.MinValue)
			{
				Write((short)value);
			}
			else
			{
				WriteImpl(Format.Int32);
				WriteInt32(value);
			}
		}

		public void Write(long value)
		{
			if (value >= 0)
			{
				Write((ulong)value);
			}
			else if (value >= int.MinValue)
			{
				Write((int)value);
			}
			else
			{
				WriteImpl(Format.Int64);
				WriteInt64(value);
			}
		}

		public void Write(float value)
		{
			WriteImpl(Format.Float32);
			ExpendSizeIfNeed(4);
			FloatUnion.Write(value, m_Buf, m_Position);
			m_Position += 4;
		}

		public void Write(double value)
		{
			WriteImpl(Format.Float64);
			ExpendSizeIfNeed(8);
			DoubleUnion.Write(value, m_Buf, m_Position);
			m_Position += 8;
		}

		public void Write(string value)
		{
			if (value == null)
			{
				WriteNil();
				return;
			}

			int length = Encoding.UTF8.GetByteCount(value);

			if (length <= 31)
			{
				WriteImpl((byte)(Format.FixStrMin | (byte)length));
			}
			else if (length <= byte.MaxValue)
			{
				WriteImpl(Format.Str8);
				WriteUInt8((byte)length);
			}
			else if (length <= ushort.MaxValue)
			{
				WriteImpl(Format.Str16);
				WriteUInt16((ushort)length);
			}
			else
			{
				WriteImpl(Format.Str32);
				WriteUInt32((uint)length);
			}

			ExpendSizeIfNeed(length);
			Encoding.UTF8.GetBytes(value, 0, value.Length, m_Buf, m_Position);
			m_Position += length;
		}

		public void Write(byte[] bytes)
		{
			if (bytes == null)
			{
				WriteNil();
				return;
			}

			if (bytes.Length <= byte.MaxValue)
			{
				WriteImpl(Format.Bin8);
				WriteUInt8((byte)bytes.Length);
			}
			else if (bytes.Length <= ushort.MaxValue)
			{
				WriteImpl(Format.Bin16);
				WriteUInt16((ushort)bytes.Length);
			}
			else
			{
				WriteImpl(Format.Bin32);
				WriteUInt32((uint)bytes.Length);
			}

			ExpendSizeIfNeed(bytes.Length);
			System.Buffer.BlockCopy(bytes, 0, m_Buf, m_Position, bytes.Length);
			m_Position += bytes.Length;
		}

		public void WriteArrayHeader(int length)
		{
			if (length <= 15)
			{
				WriteImpl((byte)(length | Format.FixArrayMin));
			}
			else if (length <= ushort.MaxValue)
			{
				WriteImpl(Format.Array16);
				WriteUInt16((ushort)length);
			}
			else
			{
				WriteImpl(Format.Array32);
				WriteUInt32((uint)length);
			}
		}

		public void WriteMapHeader(int length)
		{
			if (length <= 15)
			{
				WriteImpl((byte)(length | Format.FixMapMin));
			}
			else if (length <= ushort.MaxValue)
			{
				WriteImpl(Format.Map16);
				WriteUInt16((ushort)length);
			}
			else
			{
				WriteImpl(Format.Map32);
				WriteUInt32((uint)length);
			}
		}

		public void WriteExt(byte extType, byte[] bytes)
		{
			uint length = bytes.Length;
			if (length == 1)
			{
				WriteImpl(Format.FixExt1);
			}
			else if (length == 2)
			{
				WriteImpl(Format.FixExt2);
			}
			else if (length == 4)
			{
				WriteImpl(Format.FixExt4);
			}
			else if (length == 8)
			{
				WriteImpl(Format.FixExt8);
			}
			else if (length == 16)
			{
				WriteImpl(Format.FixExt16);
			}
			else if (length <= byte.MaxValue)
			{
				WriteImpl(Format.Ext8);
				WriteUInt8((byte)length);
			}
			else if (length <= ushort.MaxValue)
			{
				WriteImpl(Format.Ext16);
				WriteUInt16((ushort)length);
			}
			else if (length <= uint.MaxValue)
			{
				WriteImpl(Format.Ext32);
				WriteUInt32(length);
			}
			WriteImpl((byte)extType);
			ExpendSizeIfNeed(bytes.Length);
			System.Buffer.BlockCopy(bytes, 0, m_Buf, m_Position, bytes.Length);
			m_Position += bytes.Length;
		}

		void WritePositiveFixInt(byte value)
		{
			WriteImpl((byte)(value | Format.PositiveFixIntMin));
		}

		void WriteUInt8(byte value)
		{
			WriteImpl(value);
		}

		void WriteUInt16(ushort value)
		{
			ExpendSizeIfNeed(2);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

		void WriteUInt32(uint value)
		{
			ExpendSizeIfNeed(4);
			m_Buf[m_Position++] = (byte)(value >> 24);
			m_Buf[m_Position++] = (byte)(value >> 16);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

		void WriteUInt64(ulong value)
		{
			ExpendSizeIfNeed(8);
			m_Buf[m_Position++] = (byte)(value >> 56);
			m_Buf[m_Position++] = (byte)(value >> 48);
			m_Buf[m_Position++] = (byte)(value >> 40);
			m_Buf[m_Position++] = (byte)(value >> 32);
			m_Buf[m_Position++] = (byte)(value >> 24);
			m_Buf[m_Position++] = (byte)(value >> 16);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

		void WriteNegativeFixInt(sbyte value)
		{
			WriteImpl((byte)((byte)value | Format.NegativeFixIntMin));
		}

		void WriteInt8(sbyte value)
		{
			WriteImpl((byte)value);
		}

		void WriteInt16(short value)
		{
			ExpendSizeIfNeed(2);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

		void WriteInt32(int value)
		{
			ExpendSizeIfNeed(4);
			m_Buf[m_Position++] = (byte)(value >> 24);
			m_Buf[m_Position++] = (byte)(value >> 16);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

		void WriteInt64(long value)
		{
			ExpendSizeIfNeed(8);
			m_Buf[m_Position++] = (byte)(value >> 56);
			m_Buf[m_Position++] = (byte)(value >> 48);
			m_Buf[m_Position++] = (byte)(value >> 40);
			m_Buf[m_Position++] = (byte)(value >> 32);
			m_Buf[m_Position++] = (byte)(value >> 24);
			m_Buf[m_Position++] = (byte)(value >> 16);
			m_Buf[m_Position++] = (byte)(value >> 8);
			m_Buf[m_Position++] = (byte)(value);
		}

	}
}
