using System.Threading;
namespace ILib.ProtoPack
{
	public interface IMessage
	{
		void Write(IWriter w);
		void Read(IReader r);
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
		void Write(byte[] val);
		void Write(IMessage val);
	}

	public interface IReader
	{
		int Offset { get; }
		void Reset(byte[] buf, int offest = 0);
		void Skip();
		bool NextFormatIsNull();
		uint ReadTag();
		void ReadNil();
		T ReadMessage<T>() where T : IMessage, new();
		byte[] ReadBytes();
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

	public static class StaticPacker
	{
		static ThreadLocal<MsgPackWriter> s_Writer = new ThreadLocal<MsgPackWriter>(() => new MsgPackWriter());
		static ThreadLocal<MsgPackReader> s_Reader = new ThreadLocal<MsgPackReader>(() => new MsgPackReader());

		public static byte[] Pack(this IMessage message)
		{
			var w = s_Writer.Value;
			w.Clear();
			message.Write(w);
			return w.ToArray();
		}

		public static T Unpack<T>(this byte[] buf) where T : IMessage, new()
		{
			int offest = 0;
			return Unpack<T>(buf, ref offest);
		}

		public static T Unpack<T>(this byte[] buf, ref int offest) where T : IMessage, new()
		{
			var r = s_Reader.Value;
			r.Reset(buf, offest);
			T message = new T();
			message.Read(r);
			offest = r.Offset;
			r.Reset(null, 0);
			return message;
		}

	}

}
