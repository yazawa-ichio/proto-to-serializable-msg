using System;
using System.IO;
using System.Text;

namespace ILib.MsgPack
{
	public class Reader
	{
		byte[] m_Buf;
		int m_Pos;
		byte m_CurrentFormat;
		public byte CurrentFormat { get { return m_CurrentFormat; } }

		public Reader(byte[] buf, int pos = 0)
		{
			m_Buf = buf;
			m_Pos = pos;
		}

		public void Reset(byte[] buf, int pos = 0)
		{
			m_Buf = buf;
			m_Pos = pos;
		}

		public byte NextFormat()
		{
			return m_Buf[m_Pos];
		}

		public byte ReadFormat()
		{
			return m_CurrentFormat = m_Buf[m_Pos++];
		}

		public byte ReadPositiveFixInt()
		{
			return (byte)(m_CurrentFormat & 0x7f);
		}

		public byte ReadUInt8()
		{
			return m_Buf[m_Pos++];
		}

		public ushort ReadUInt16()
		{
			return (ushort)((m_Buf[m_Pos++] << 8) | m_Buf[m_Pos++]);
		}

		public uint ReadUInt32()
		{
			return ((uint)m_Buf[m_Pos++] << 24) | ((uint)m_Buf[m_Pos++] << 16) | ((uint)m_Buf[m_Pos++] << 8) | (uint)m_Buf[m_Pos++];
		}

		public ulong ReadUInt64()
		{
			return ((ulong)m_Buf[m_Pos++] << 56) | ((ulong)m_Buf[m_Pos++] << 48) | ((ulong)m_Buf[m_Pos++] << 40) | ((ulong)m_Buf[m_Pos++] << 32) | ((ulong)m_Buf[m_Pos++] << 24) | ((ulong)m_Buf[m_Pos++] << 16) | ((ulong)m_Buf[m_Pos++] << 8) | (ulong)m_Buf[m_Pos++];
		}

		public sbyte ReadNegativeFixInt()
		{
			return (sbyte)((m_CurrentFormat & 0x1f) - 0x20);
		}

		public sbyte ReadInt8()
		{
			return (sbyte)m_Buf[m_Pos++];
		}

		public short ReadInt16()
		{
			return (short)((m_Buf[m_Pos++] << 8) | m_Buf[m_Pos++]);
		}

		public int ReadInt32()
		{
			return (m_Buf[m_Pos++] << 24) | (m_Buf[m_Pos++] << 16) | (m_Buf[m_Pos++] << 8) | m_Buf[m_Pos++];
		}

		public long ReadInt64()
		{
			return ((long)m_Buf[m_Pos++] << 56) | ((long)m_Buf[m_Pos++] << 48) | ((long)m_Buf[m_Pos++] << 40) | ((long)m_Buf[m_Pos++] << 32) | ((long)m_Buf[m_Pos++] << 24) | ((long)m_Buf[m_Pos++] << 16) | ((long)m_Buf[m_Pos++] << 8) | (long)m_Buf[m_Pos++];
		}

		public float ReadFloat32()
		{
			m_Pos += 4;
			return FloatUnion.Read(m_Buf, m_Pos - 4);
		}

		public double ReadFloat64()
		{
			m_Pos += 8;
			return DoubleUnion.Read(m_Buf, m_Pos - 8);
		}

		public string ReadFixStr()
		{
			return ReadStringOfLength(m_CurrentFormat & 0x1f);
		}

		public string ReadStr8()
		{
			return ReadStringOfLength(ReadUInt8());
		}

		public string ReadStr16()
		{
			return ReadStringOfLength(ReadUInt16());
		}

		public string ReadStr32()
		{
			return ReadStringOfLength(Convert.ToInt32(ReadUInt32()));
		}

		public ArraySegment<byte> ReadBin8()
		{
			return ReadBytesOfLength(ReadUInt8());
		}

		public ArraySegment<byte> ReadBin16()
		{
			return ReadBytesOfLength(ReadUInt16());
		}

		public ArraySegment<byte> ReadBin32()
		{
			return ReadBytesOfLength(Convert.ToInt32(ReadUInt32()));
		}

		public int ReadArrayLength()
		{
			switch (m_CurrentFormat)
			{
				case Format.Nil: return 0;
				case Format.Array16: return ReadUInt16();
				case Format.Array32: return (int)ReadUInt32();
			}
			if (Format.IsFixArray(m_CurrentFormat))
			{
				return (byte)(m_CurrentFormat & 0xf);
			}
			throw new InvalidOperationException(string.Format("msgpack format invalid {0:X2}", m_CurrentFormat));
		}

		public int ReadMapLength()
		{
			switch (m_CurrentFormat)
			{
				case Format.Map16: return ReadUInt16();
				case Format.Map32: return (int)ReadUInt32();
			}
			if (Format.IsFixMap(m_CurrentFormat))
			{
				return (byte)(m_CurrentFormat & 0xf);
			}
			throw new InvalidOperationException(string.Format("msgpack format invalid {0:X2}", m_CurrentFormat));
		}

		public uint ReadExtLength()
		{
			switch (m_CurrentFormat)
			{
				case Format.FixExt1: return 1;
				case Format.FixExt2: return 2;
				case Format.FixExt4: return 4;
				case Format.FixExt8: return 8;
				case Format.FixExt16: return 16;
				case Format.Ext8: return ReadUInt8();
				case Format.Ext16: return ReadUInt16();
				case Format.Ext32: return ReadUInt32();
			}
			throw new InvalidOperationException(string.Format("msgpack format invalid {0:X2}", m_CurrentFormat));
		}

		public sbyte ReadExtType()
		{
			if (m_CurrentFormat == Format.UInt8)
			{
				return (sbyte)ReadUInt8();
			}
			else if (m_CurrentFormat == Format.Int8)
			{
				return (sbyte)ReadUInt8();
			}
			else if (Format.IsPositiveFixInt(m_CurrentFormat))
			{
				return (sbyte)ReadPositiveFixInt();
			}
			else if (Format.IsNegativeFixInt(m_CurrentFormat))
			{
				return ReadNegativeFixInt();
			}
			throw new InvalidOperationException(string.Format("msgpack format invalid {0:X2}", m_CurrentFormat));
		}

		public void Skip()
		{
			switch (ReadFormat())
			{
				case Format.Nil:
				case Format.False:
				case Format.True:
					return;
				case Format.UInt8:
				case Format.Int8:
					m_Pos += 1;
					return;
				case Format.UInt16:
				case Format.Int16:
					m_Pos += 2;
					return;
				case Format.UInt32:
				case Format.Int32:
				case Format.Float32:
					m_Pos += 4;
					return;
				case Format.UInt64:
				case Format.Int64:
				case Format.Float64:
					m_Pos += 8;
					return;
				case Format.Str8:
				case Format.Bin8:
					m_Pos += ReadUInt8();
					return;
				case Format.Str16:
				case Format.Bin16:
					m_Pos += ReadUInt16();
					return;
				case Format.Str32:
				case Format.Bin32:
					m_Pos += (int)ReadUInt32();
					return;
				case Format.FixExt1:
					m_Pos += 2;
					return;
				case Format.FixExt2:
					m_Pos += 3;
					return;
				case Format.FixExt4:
					m_Pos += 5;
					return;
				case Format.FixExt8:
					m_Pos += 9;
					return;
				case Format.FixExt16:
					m_Pos += 17;
					return;
				case Format.Ext8:
					m_Pos += 1 + ReadUInt8();
					return;
				case Format.Ext16:
					m_Pos += 1 + ReadUInt16();
					return;
				case Format.Ext32:
					m_Pos += 1 + (int)ReadUInt32();
					return;
			}
			if (Format.IsPositiveFixInt(m_CurrentFormat) || Format.IsNegativeFixInt(m_CurrentFormat))
			{
				return;
			}
			else if (Format.IsFixStr(m_CurrentFormat))
			{
				m_Pos += m_CurrentFormat & 0x1f;
			}
			else if (Format.IsArrayFamily(m_CurrentFormat))
			{
				for (int length = ReadArrayLength(); length > 0; length--)
				{
					Skip();
				}
			}
			else if (Format.IsMapFamily(m_CurrentFormat))
			{
				for (int length = ReadMapLength(); length > 0; length--)
				{
					Skip();
					Skip();
				}
			}
		}

		string ReadStringOfLength(int length)
		{
			m_Pos += length;
			return Encoding.UTF8.GetString(m_Buf, m_Pos - length, length);
		}

		internal ArraySegment<byte> ReadBytesOfLength(int length)
		{
			m_Pos += length;
			return new ArraySegment<byte>(m_Buf, m_Pos - length, length);
		}
	}
}
