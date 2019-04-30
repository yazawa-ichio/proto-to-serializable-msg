namespace ILib.ProtoPack
{
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

}
