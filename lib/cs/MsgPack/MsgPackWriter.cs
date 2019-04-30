using System;
using ILib.MsgPack;

namespace ILib.ProtoPack
{
	public class MsgPackWriter : IWriter
	{
		Writer m_Writer;

		public int Length { get { return m_Writer.Length; } }

		public MsgPackWriter()
		{
			m_Writer = new Writer();
		}

		public byte[] ToArray()
		{
			return m_Writer.ToArray();
		}

		public void Read(byte[] dst, int offest = 0)
		{
			m_Writer.Read(dst, offest);
		}

		public ArraySegment<byte> ReadSegment()
		{
			return m_Writer.ReadSegment();
		}

		public void Clear()
		{
			m_Writer.Clear();
		}

		public void WriteMapHeader(int length)
		{
			m_Writer.WriteMapHeader(length);
		}

		public void WriteArrayHeader(int length)
		{
			m_Writer.WriteArrayHeader(length);
		}

		public void WriteTag(uint tag)
		{
			m_Writer.Write(tag);
		}

		public void WriteNil()
		{
			m_Writer.WriteNil();
		}

		public void Write(double val)
		{
			m_Writer.Write(val);
		}

		public void Write(float val)
		{
			m_Writer.Write(val);
		}

		public void Write(long val)
		{
			m_Writer.Write(val);
		}

		public void Write(ulong val)
		{
			m_Writer.Write(val);
		}

		public void Write(int val)
		{
			m_Writer.Write(val);
		}

		public void Write(uint val)
		{
			m_Writer.Write(val);
		}

		public void Write(bool val)
		{
			m_Writer.Write(val);
		}

		public void Write(string val)
		{
			m_Writer.Write(val);
		}

		public void WriteBytes(byte[] val)
		{
			m_Writer.Write(val);
		}
	}

}
