// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class AllParameter : IMessage
{
	public double ValueDouble { get; set; }

	public float ValueFloat { get; set; }

	public int ValueInt32 { get; set; }

	public long ValueInt64 { get; set; }

	public uint ValueUint32 { get; set; }

	public ulong ValueUint64 { get; set; }

	public int ValueSint32 { get; set; }

	public long ValueSint64 { get; set; }

	public uint ValueFixed32 { get; set; }

	public ulong ValueFixed64 { get; set; }

	public int ValueSfixed32 { get; set; }

	public long ValueSfixed64 { get; set; }

	public bool ValueBool { get; set; }

	public string ValueString { get; set; }

	public byte[] ValueBytes { get; set; }

	public Dictionary<int, string> ValueMapString { get; set; }

	public Dictionary<string, int> ValueMapInt { get; set; }

	public EmptyMessage ValueMessage { get; set; }

	public Dictionary<int, DependTest> ValueMapValueMessage { get; set; }

	public TestEnum ValueTestEnum { get; set; }

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(20);
		
		// Write value_double
		w.WriteTag(1);
		w.Write(ValueDouble);
		
		// Write value_float
		w.WriteTag(2);
		w.Write(ValueFloat);
		
		// Write value_int32
		w.WriteTag(3);
		w.Write(ValueInt32);
		
		// Write value_int64
		w.WriteTag(4);
		w.Write(ValueInt64);
		
		// Write value_uint32
		w.WriteTag(5);
		w.Write(ValueUint32);
		
		// Write value_uint64
		w.WriteTag(6);
		w.Write(ValueUint64);
		
		// Write value_sint32
		w.WriteTag(7);
		w.Write(ValueSint32);
		
		// Write value_sint64
		w.WriteTag(8);
		w.Write(ValueSint64);
		
		// Write value_fixed32
		w.WriteTag(9);
		w.Write(ValueFixed32);
		
		// Write value_fixed64
		w.WriteTag(10);
		w.Write(ValueFixed64);
		
		// Write value_sfixed32
		w.WriteTag(11);
		w.Write(ValueSfixed32);
		
		// Write value_sfixed64
		w.WriteTag(12);
		w.Write(ValueSfixed64);
		
		// Write value_bool
		w.WriteTag(13);
		w.Write(ValueBool);
		
		// Write value_string
		w.WriteTag(14);
		w.Write(ValueString);
		
		// Write value_bytes
		w.WriteTag(15);
		w.Write(ValueBytes);
		
		// Write value_map_string
		w.WriteTag(16);
		if (ValueMapString == null)
		{
			w.WriteNil();
		}
		else
		{
			var mapLen = ValueMapString.Count;
			w.WriteMapHeader(mapLen);
			foreach(var _ValueMapStringEntry in ValueMapString){
				w.Write(_ValueMapStringEntry.Key);
				w.Write(_ValueMapStringEntry.Value);
			}
		}
		
		// Write value_map_int
		w.WriteTag(17);
		if (ValueMapInt == null)
		{
			w.WriteNil();
		}
		else
		{
			var mapLen = ValueMapInt.Count;
			w.WriteMapHeader(mapLen);
			foreach(var _ValueMapIntEntry in ValueMapInt){
				w.Write(_ValueMapIntEntry.Key);
				w.Write(_ValueMapIntEntry.Value);
			}
		}
		
		// Write value_message
		w.WriteTag(18);
		w.Write(ValueMessage);
		
		// Write value_map_value_message
		w.WriteTag(19);
		if (ValueMapValueMessage == null)
		{
			w.WriteNil();
		}
		else
		{
			var mapLen = ValueMapValueMessage.Count;
			w.WriteMapHeader(mapLen);
			foreach(var _ValueMapValueMessageEntry in ValueMapValueMessage){
				w.Write(_ValueMapValueMessageEntry.Key);
				w.Write(_ValueMapValueMessageEntry.Value);
			}
		}
		
		// Write value_testEnum
		w.WriteTag(20);
		w.Write((int)ValueTestEnum);
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
				ValueDouble = r.ReadDouble();
				break;
			case 2: // Read value_float
				ValueFloat = r.ReadFloat();
				break;
			case 3: // Read value_int32
				ValueInt32 = r.ReadInt();
				break;
			case 4: // Read value_int64
				ValueInt64 = r.ReadLong();
				break;
			case 5: // Read value_uint32
				ValueUint32 = r.ReadUint();
				break;
			case 6: // Read value_uint64
				ValueUint64 = r.ReadUlong();
				break;
			case 7: // Read value_sint32
				ValueSint32 = r.ReadInt();
				break;
			case 8: // Read value_sint64
				ValueSint64 = r.ReadLong();
				break;
			case 9: // Read value_fixed32
				ValueFixed32 = r.ReadUint();
				break;
			case 10: // Read value_fixed64
				ValueFixed64 = r.ReadUlong();
				break;
			case 11: // Read value_sfixed32
				ValueSfixed32 = r.ReadInt();
				break;
			case 12: // Read value_sfixed64
				ValueSfixed64 = r.ReadLong();
				break;
			case 13: // Read value_bool
				ValueBool = r.ReadBool();
				break;
			case 14: // Read value_string
				ValueString = r.ReadString();
				break;
			case 15: // Read value_bytes
				ValueBytes = r.ReadBytes();
				break;
			case 16: // Read value_map_string
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					ValueMapString = null;
					continue;
				}
				var _ValueMapStringLen = r.ReadMapHeader();
				ValueMapString = new Dictionary<int, string>(_ValueMapStringLen);
				for(int mapIndex = 0; mapIndex < _ValueMapStringLen; mapIndex++)
				{
					var _ValueMapStringKey = default(int);
					var _ValueMapStringValue = default(string);
					_ValueMapStringKey = r.ReadInt();
					_ValueMapStringValue = r.ReadString();
					ValueMapString[_ValueMapStringKey] = _ValueMapStringValue;
				}
				break;
			case 17: // Read value_map_int
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					ValueMapInt = null;
					continue;
				}
				var _ValueMapIntLen = r.ReadMapHeader();
				ValueMapInt = new Dictionary<string, int>(_ValueMapIntLen);
				for(int mapIndex = 0; mapIndex < _ValueMapIntLen; mapIndex++)
				{
					var _ValueMapIntKey = default(string);
					var _ValueMapIntValue = default(int);
					_ValueMapIntKey = r.ReadString();
					_ValueMapIntValue = r.ReadInt();
					ValueMapInt[_ValueMapIntKey] = _ValueMapIntValue;
				}
				break;
			case 18: // Read value_message
				ValueMessage = r.ReadMessage<EmptyMessage>();
				break;
			case 19: // Read value_map_value_message
				if(r.NextFormatIsNull())
				{
					r.ReadNil();
					ValueMapValueMessage = null;
					continue;
				}
				var _ValueMapValueMessageLen = r.ReadMapHeader();
				ValueMapValueMessage = new Dictionary<int, DependTest>(_ValueMapValueMessageLen);
				for(int mapIndex = 0; mapIndex < _ValueMapValueMessageLen; mapIndex++)
				{
					var _ValueMapValueMessageKey = default(int);
					var _ValueMapValueMessageValue = default(DependTest);
					_ValueMapValueMessageKey = r.ReadInt();
					_ValueMapValueMessageValue = r.ReadMessage<DependTest>();
					ValueMapValueMessage[_ValueMapValueMessageKey] = _ValueMapValueMessageValue;
				}
				break;
			case 20: // Read value_testEnum
				ValueTestEnum = (TestEnum)r.ReadInt();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
