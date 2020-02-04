// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class AllRepeatedParameter : IMessage
{
	public double[] ValueDouble { get; set; }

	public float[] ValueFloat { get; set; }

	public int[] ValueInt32 { get; set; }

	public long[] ValueInt64 { get; set; }

	public uint[] ValueUint32 { get; set; }

	public ulong[] ValueUint64 { get; set; }

	public int[] ValueSint32 { get; set; }

	public long[] ValueSint64 { get; set; }

	public uint[] ValueFixed32 { get; set; }

	public ulong[] ValueFixed64 { get; set; }

	public int[] ValueSfixed32 { get; set; }

	public long[] ValueSfixed64 { get; set; }

	public bool[] ValueBool { get; set; }

	public string[] ValueString { get; set; }

	public byte[][] ValueBytes { get; set; }

	public DependTest[] ValueNestMessage { get; set; }

	public TestEnum[] ValueTestEnum { get; set; }

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(17);
		
		// Write value_double
		w.WriteTag(1);
		if (ValueDouble == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueDouble.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueDouble[arrayIndex]);
			}
		}
		
		// Write value_float
		w.WriteTag(2);
		if (ValueFloat == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueFloat.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueFloat[arrayIndex]);
			}
		}
		
		// Write value_int32
		w.WriteTag(3);
		if (ValueInt32 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueInt32.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueInt32[arrayIndex]);
			}
		}
		
		// Write value_int64
		w.WriteTag(4);
		if (ValueInt64 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueInt64.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueInt64[arrayIndex]);
			}
		}
		
		// Write value_uint32
		w.WriteTag(5);
		if (ValueUint32 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueUint32.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueUint32[arrayIndex]);
			}
		}
		
		// Write value_uint64
		w.WriteTag(6);
		if (ValueUint64 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueUint64.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueUint64[arrayIndex]);
			}
		}
		
		// Write value_sint32
		w.WriteTag(7);
		if (ValueSint32 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueSint32.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueSint32[arrayIndex]);
			}
		}
		
		// Write value_sint64
		w.WriteTag(8);
		if (ValueSint64 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueSint64.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueSint64[arrayIndex]);
			}
		}
		
		// Write value_fixed32
		w.WriteTag(9);
		if (ValueFixed32 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueFixed32.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueFixed32[arrayIndex]);
			}
		}
		
		// Write value_fixed64
		w.WriteTag(10);
		if (ValueFixed64 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueFixed64.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueFixed64[arrayIndex]);
			}
		}
		
		// Write value_sfixed32
		w.WriteTag(11);
		if (ValueSfixed32 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueSfixed32.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueSfixed32[arrayIndex]);
			}
		}
		
		// Write value_sfixed64
		w.WriteTag(12);
		if (ValueSfixed64 == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueSfixed64.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueSfixed64[arrayIndex]);
			}
		}
		
		// Write value_bool
		w.WriteTag(13);
		if (ValueBool == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueBool.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueBool[arrayIndex]);
			}
		}
		
		// Write value_string
		w.WriteTag(14);
		if (ValueString == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueString.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueString[arrayIndex]);
			}
		}
		
		// Write value_bytes
		w.WriteTag(15);
		if (ValueBytes == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueBytes.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueBytes[arrayIndex]);
			}
		}
		
		// Write ValueNestMessage
		w.WriteTag(18);
		if (ValueNestMessage == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueNestMessage.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write(ValueNestMessage[arrayIndex]);
			}
		}
		
		// Write ValueTestEnum
		w.WriteTag(20);
		if (ValueTestEnum == null)
		{
			w.WriteNil();
		}
		else
		{
			var arrayLen = ValueTestEnum.Length;
			w.WriteArrayHeader(arrayLen);
			for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.Write((int)ValueTestEnum[arrayIndex]);
			}
		}
	}

	/// <summary>
	/// Deserialize Message
	/// </summary>
	public void Read(IReader r)
	{
		// Read Map Length
		var len = r.ReadMapHeader();

		for (var i = 0; i < len; i++)
		{
			var tag = r.ReadTag();
			switch(tag) {
			case 1: // Read value_double
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueDouble = null;
					continue;
				}
				var _ValueDoubleLen = r.ReadArrayHeader();
				ValueDouble = new double[_ValueDoubleLen];
				for(int arrayIndex = 0; arrayIndex < _ValueDoubleLen; arrayIndex++)
				{
					ValueDouble[arrayIndex] = r.ReadDouble();
				}
				break;
			case 2: // Read value_float
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueFloat = null;
					continue;
				}
				var _ValueFloatLen = r.ReadArrayHeader();
				ValueFloat = new float[_ValueFloatLen];
				for(int arrayIndex = 0; arrayIndex < _ValueFloatLen; arrayIndex++)
				{
					ValueFloat[arrayIndex] = r.ReadFloat();
				}
				break;
			case 3: // Read value_int32
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueInt32 = null;
					continue;
				}
				var _ValueInt32Len = r.ReadArrayHeader();
				ValueInt32 = new int[_ValueInt32Len];
				for(int arrayIndex = 0; arrayIndex < _ValueInt32Len; arrayIndex++)
				{
					ValueInt32[arrayIndex] = r.ReadInt();
				}
				break;
			case 4: // Read value_int64
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueInt64 = null;
					continue;
				}
				var _ValueInt64Len = r.ReadArrayHeader();
				ValueInt64 = new long[_ValueInt64Len];
				for(int arrayIndex = 0; arrayIndex < _ValueInt64Len; arrayIndex++)
				{
					ValueInt64[arrayIndex] = r.ReadLong();
				}
				break;
			case 5: // Read value_uint32
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueUint32 = null;
					continue;
				}
				var _ValueUint32Len = r.ReadArrayHeader();
				ValueUint32 = new uint[_ValueUint32Len];
				for(int arrayIndex = 0; arrayIndex < _ValueUint32Len; arrayIndex++)
				{
					ValueUint32[arrayIndex] = r.ReadUint();
				}
				break;
			case 6: // Read value_uint64
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueUint64 = null;
					continue;
				}
				var _ValueUint64Len = r.ReadArrayHeader();
				ValueUint64 = new ulong[_ValueUint64Len];
				for(int arrayIndex = 0; arrayIndex < _ValueUint64Len; arrayIndex++)
				{
					ValueUint64[arrayIndex] = r.ReadUlong();
				}
				break;
			case 7: // Read value_sint32
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueSint32 = null;
					continue;
				}
				var _ValueSint32Len = r.ReadArrayHeader();
				ValueSint32 = new int[_ValueSint32Len];
				for(int arrayIndex = 0; arrayIndex < _ValueSint32Len; arrayIndex++)
				{
					ValueSint32[arrayIndex] = r.ReadInt();
				}
				break;
			case 8: // Read value_sint64
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueSint64 = null;
					continue;
				}
				var _ValueSint64Len = r.ReadArrayHeader();
				ValueSint64 = new long[_ValueSint64Len];
				for(int arrayIndex = 0; arrayIndex < _ValueSint64Len; arrayIndex++)
				{
					ValueSint64[arrayIndex] = r.ReadLong();
				}
				break;
			case 9: // Read value_fixed32
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueFixed32 = null;
					continue;
				}
				var _ValueFixed32Len = r.ReadArrayHeader();
				ValueFixed32 = new uint[_ValueFixed32Len];
				for(int arrayIndex = 0; arrayIndex < _ValueFixed32Len; arrayIndex++)
				{
					ValueFixed32[arrayIndex] = r.ReadUint();
				}
				break;
			case 10: // Read value_fixed64
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueFixed64 = null;
					continue;
				}
				var _ValueFixed64Len = r.ReadArrayHeader();
				ValueFixed64 = new ulong[_ValueFixed64Len];
				for(int arrayIndex = 0; arrayIndex < _ValueFixed64Len; arrayIndex++)
				{
					ValueFixed64[arrayIndex] = r.ReadUlong();
				}
				break;
			case 11: // Read value_sfixed32
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueSfixed32 = null;
					continue;
				}
				var _ValueSfixed32Len = r.ReadArrayHeader();
				ValueSfixed32 = new int[_ValueSfixed32Len];
				for(int arrayIndex = 0; arrayIndex < _ValueSfixed32Len; arrayIndex++)
				{
					ValueSfixed32[arrayIndex] = r.ReadInt();
				}
				break;
			case 12: // Read value_sfixed64
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueSfixed64 = null;
					continue;
				}
				var _ValueSfixed64Len = r.ReadArrayHeader();
				ValueSfixed64 = new long[_ValueSfixed64Len];
				for(int arrayIndex = 0; arrayIndex < _ValueSfixed64Len; arrayIndex++)
				{
					ValueSfixed64[arrayIndex] = r.ReadLong();
				}
				break;
			case 13: // Read value_bool
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueBool = null;
					continue;
				}
				var _ValueBoolLen = r.ReadArrayHeader();
				ValueBool = new bool[_ValueBoolLen];
				for(int arrayIndex = 0; arrayIndex < _ValueBoolLen; arrayIndex++)
				{
					ValueBool[arrayIndex] = r.ReadBool();
				}
				break;
			case 14: // Read value_string
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueString = null;
					continue;
				}
				var _ValueStringLen = r.ReadArrayHeader();
				ValueString = new string[_ValueStringLen];
				for(int arrayIndex = 0; arrayIndex < _ValueStringLen; arrayIndex++)
				{
					ValueString[arrayIndex] = r.ReadString();
				}
				break;
			case 15: // Read value_bytes
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueBytes = null;
					continue;
				}
				var _ValueBytesLen = r.ReadArrayHeader();
				ValueBytes = new byte[_ValueBytesLen][];
				for(int arrayIndex = 0; arrayIndex < _ValueBytesLen; arrayIndex++)
				{
					ValueBytes[arrayIndex] = r.ReadBytes();
				}
				break;
			case 18: // Read ValueNestMessage
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueNestMessage = null;
					continue;
				}
				var _ValueNestMessageLen = r.ReadArrayHeader();
				ValueNestMessage = new DependTest[_ValueNestMessageLen];
				for(int arrayIndex = 0; arrayIndex < _ValueNestMessageLen; arrayIndex++)
				{
					ValueNestMessage[arrayIndex] = r.ReadMessage<DependTest>();
				}
				break;
			case 20: // Read ValueTestEnum
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					this.ValueTestEnum = null;
					continue;
				}
				var _ValueTestEnumLen = r.ReadArrayHeader();
				ValueTestEnum = new TestEnum[_ValueTestEnumLen];
				for(int arrayIndex = 0; arrayIndex < _ValueTestEnumLen; arrayIndex++)
				{
					ValueTestEnum[arrayIndex] = (TestEnum)r.ReadInt();
				}
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
