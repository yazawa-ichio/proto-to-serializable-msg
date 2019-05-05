using System;
using ILib.MsgPack;

namespace ILib.ProtoPack.MsgPack
{
	public class MsgPackReader : IReader
	{
		Reader m_Reader;

		public MsgPackReader(byte[] buf, int offest = 0)
		{
			m_Reader = new Reader(buf, offest);
		}

		public void Reset(byte[] buf, int offest = 0)
		{
			m_Reader.Reset(buf, offest);
		}

		string MissMatchErrorText()
		{
			return string.Format("MsgPackReader fail read format miss match 0x{0:X2}", m_Reader.CurrentFormat);
		}

		public void Skip()
		{
			m_Reader.Skip();
		}

		public uint ReadTag(int index, int length)
		{
			if (index < length)
			{
				return ReadUint();
			}
			return 0;
		}

		public bool IsNull()
		{
			return m_Reader.NextFormat() == Format.Nil;
		}

		public void ReadNil()
		{
			if (m_Reader.ReadFormat() != Format.Nil)
			{
				throw new InvalidOperationException(MissMatchErrorText());
			}
		}

		public T ReadNil<T>()
		{
			if (m_Reader.ReadFormat() != Format.Nil)
			{
				throw new InvalidOperationException(MissMatchErrorText());
			}
			return default(T);
		}

		public T New<T>() where T : new()
		{
			return new T();
		}

		public T[] NewArray<T>(int length)
		{
			return new T[length];
		}

		public void ReadBytes(ref byte[] buf, bool overridable)
		{
			ArraySegment<byte> source;
			switch (m_Reader.ReadFormat())
			{
				case Format.Bin8:
					source = m_Reader.ReadBin8();
					break;
				case Format.Bin16:
					source = m_Reader.ReadBin16();
					break;
				case Format.Bin32:
					source = m_Reader.ReadBin32();
					break;
				default:
					throw new InvalidOperationException(MissMatchErrorText());
			}
			if (buf == null || !overridable)
			{
				buf = new byte[source.Count];
			}
			else if (buf.Length != source.Count)
			{
				Array.Resize(ref buf, source.Count);
			}
			Buffer.BlockCopy(source.Array, source.Offset, buf, 0, source.Count);
		}

		public bool ReadBool()
		{
			switch (m_Reader.ReadFormat())
			{
				case Format.True:
					return true;
				case Format.False:
					return false;
			}
			throw new InvalidOperationException(MissMatchErrorText());
		}

		public int ReadInt()
		{
			return (int)ReadLong();
		}

		public uint ReadUint()
		{
			return (uint)ReadUlong();
		}

		public long ReadLong()
		{
			var format = m_Reader.ReadFormat();
			switch (format)
			{
				case Format.Float64:
					return (long)m_Reader.ReadFloat64();
				case Format.Float32:
					return (long)m_Reader.ReadFloat32();
				case Format.UInt8:
					return m_Reader.ReadUInt8();
				case Format.UInt16:
					return m_Reader.ReadUInt16();
				case Format.UInt32:
					return m_Reader.ReadUInt32();
				case Format.Int8:
					return m_Reader.ReadInt8();
				case Format.Int16:
					return m_Reader.ReadInt16();
				case Format.Int32:
					return m_Reader.ReadInt32();
				case Format.Int64:
					return m_Reader.ReadInt64();
			}
			if (Format.IsPositiveFixInt(format))
			{
				return m_Reader.ReadPositiveFixInt();
			}
			else if (Format.IsNegativeFixInt(format))
			{
				return m_Reader.ReadNegativeFixInt();
			}
			throw new InvalidOperationException(MissMatchErrorText());
		}

		public ulong ReadUlong()
		{
			var format = m_Reader.ReadFormat();
			switch (format)
			{
				case Format.UInt8:
					return m_Reader.ReadUInt8();
				case Format.UInt16:
					return m_Reader.ReadUInt16();
				case Format.UInt32:
					return m_Reader.ReadUInt32();
				case Format.UInt64:
					return m_Reader.ReadUInt64();
				case Format.Int8:
					return (ulong)m_Reader.ReadInt8();
				case Format.Int16:
					return (ulong)m_Reader.ReadInt16();
				case Format.Int32:
					return (ulong)m_Reader.ReadInt32();
				case Format.Int64:
					return (ulong)m_Reader.ReadInt64();
			}
			if (Format.IsPositiveFixInt(format))
			{
				return m_Reader.ReadPositiveFixInt();
			}
			throw new InvalidOperationException(MissMatchErrorText());
		}

		public double ReadDouble()
		{
			if (m_Reader.NextFormat() == Format.Float64)
			{
				m_Reader.ReadFormat();
				return m_Reader.ReadFloat64();
			}
			else if (m_Reader.NextFormat() == Format.Float32)
			{
				m_Reader.ReadFormat();
				return m_Reader.ReadFloat32();
			}
			return ReadLong();
		}

		public float ReadFloat()
		{
			if(m_Reader.NextFormat() == Format.Float32) {
				m_Reader.ReadFormat();
				return m_Reader.ReadFloat32();
			}
			return ReadLong();
		}

		public string ReadString()
		{
			var format = m_Reader.ReadFormat();
			switch (format)
			{
				case Format.Str8:
					return m_Reader.ReadStr8();
				case Format.Str16:
					return m_Reader.ReadStr16();
				case Format.Str32:
					return m_Reader.ReadStr32();
				case Format.Nil:
					return null;
			}
			if(Format.IsFixStr(format))
			{
				return m_Reader.ReadFixStr();
			}
			throw new InvalidOperationException(MissMatchErrorText());
		}

		public int ReadMapHeader()
		{
			m_Reader.ReadFormat();
			return m_Reader.ReadMapLength();
		}

		public int ReadArrayHeader()
		{
			m_Reader.ReadFormat();
			return m_Reader.ReadArrayLength();
		}
	}
}
