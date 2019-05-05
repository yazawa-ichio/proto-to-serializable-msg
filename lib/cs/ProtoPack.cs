namespace ILib.ProtoPack
{
	public interface IMessage
	{
		void Write(IWriter w, bool skipable = true);
		void Read(IReader r, bool overridable = false);
	}

	public interface IWriter
	{
		int Length { get; }
		byte[] ToArray();
		void Read(byte[] dst, int offest = 0);
		System.ArraySegment<byte> ReadSegment();
		void Clear();
		void WriteMapHeader(int length);
		void WriteArrayHeader(int length);
		void WriteTag(uint tag);
		void WriteNil();
		void Write(double val);
		void Write(float val);
		void Write(long val);
		void Write(ulong val);
		void Write(int val);
		void Write(uint val);
		void Write(bool val);
		void Write(string val);
		void WriteBytes(byte[] val);
	}

	public interface IReader
	{
		int Offset { get; }
		void Reset(byte[] buf, int offest = 0);
		void Skip();
		bool IsNull();
		uint ReadTag(int index, int length);
		void ReadNil();
		T ReadNil<T>();
		void ReadBytes(ref byte[] buf, bool overridable);
		bool ReadBool();
		int ReadInt();
		uint ReadUint();
		long ReadLong();
		ulong ReadUlong();
		double ReadDouble();
		float ReadFloat();
		string ReadString();
		int ReadMapHeader();
		int ReadArrayHeader();
	}

	public interface IInstanceProvider
	{
		T Provide<T>() where T : new();
		T[] ProvideArray<T>(int length);
	}

	public class InstanceProvider : IInstanceProvider
	{
		static IInstanceProvider s_Provider = new InstanceProvider();

		public static void SetProvider(InstanceProvider provider)
		{
			s_Provider = provider;
		}

		public static T New<T>() where T : new()
		{
			return s_Provider.Provide<T>();
		}

		public static T[] NewArray<T>(int length) where T : new()
		{
			return s_Provider.ProvideArray<T>(length);
		}

		private InstanceProvider()
		{
		}


		T IInstanceProvider.Provide<T>()
		{
			return new T();
		}

		T[] IInstanceProvider.ProvideArray<T>(int length)
		{
			return new T[length];
		}

	}


	public interface IPacker
	{
		byte[] Pack(IMessage message, bool skipable = true);
		T Unpack<T>(byte[] buf) where T : IMessage, new();
		T Unpack<T>(byte[] buf, ref int offest) where T : IMessage, new();
		void Unpack(IMessage message, byte[] buf);
		void Unpack(IMessage message, byte[] buf, ref int offest);
	}

	public class Packer : Packer<MsgPackWriter, MsgPackReader>
	{
		public Packer() : base(new MsgPackWriter(), new MsgPackReader()) { }
	}

	public class Packer<TWriter, UReader> : IPacker where TWriter : IWriter where UReader : IReader
	{
		TWriter m_Writer;
		UReader m_Reader;
		public TWriter Writer { get { return m_Writer; } }
		public UReader Reader { get { return m_Reader; } }

		public Packer(TWriter w, UReader r)
		{
			m_Writer = w;
			m_Reader = r;
		}

		public byte[] Pack(IMessage message, bool skipable = true)
		{
			m_Writer.Clear();
			message.Write(m_Writer, skipable);
			return m_Writer.ToArray();
		}

		public T Unpack<T>(byte[] buf) where T : IMessage, new()
		{
			int offest = 0;
			return Unpack<T>(buf, ref offest);
		}

		public T Unpack<T>(byte[] buf, ref int offest) where T : IMessage, new()
		{
			m_Reader.Reset(buf, offest);
			T message = new T();
			message.Read(m_Reader, false);
			offest = m_Reader.Offset;
			m_Reader.Reset(null, 0);
			return message;
		}

		public void Unpack(IMessage message, byte[] buf)
		{
			int offest = 0;
			Unpack(message, buf, ref offest);
		}

		public void Unpack(IMessage message, byte[] buf, ref int offest)
		{
			m_Reader.Reset(buf, offest);
			message.Read(m_Reader, true);
			offest = m_Reader.Offset;
			m_Reader.Reset(null, 0);
		}

	}

	public static class StaticPacker
	{
		[System.ThreadStatic]
		static IPacker s_Packer;
		static IPacker GetPacker()
		{
			if (s_Packer == null)
			{
				if (PackerProvider == null)
				{
					s_Packer = new Packer();
				}
				else
				{
					s_Packer = PackerProvider();
				}
			}
			return s_Packer;
		}

		public static System.Func<IPacker> PackerProvider;

		public static byte[] Pack(this IMessage message, bool skipable = true)
		{
			return GetPacker().Pack(message, skipable);
		}
		public static T Unpack<T>(this byte[] buf) where T : IMessage, new()
		{
			return GetPacker().Unpack<T>(buf);
		}

		public static T Unpack<T>(this byte[] buf, ref int offest) where T : IMessage, new()
		{
			return GetPacker().Unpack<T>(buf, ref offest);
		}

		public static void Unpack(this IMessage message, byte[] buf)
		{
			GetPacker().Unpack(message, buf);
		}

		public static void Unpack(this IMessage message, byte[] buf, ref int offest)
		{
			GetPacker().Unpack(message, buf, ref offest);
		}

	}

}
