// tests/proto/import.proto
using ILib.ProtoPack;
using System.Collections.Generic;
public partial class PackageMessage : IMessage
{
	public MyPackage.MyMessage Message;

	public MyPackage.MyEnum MyEnum;

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(2);
		
		// Write message
		w.WriteTag(100);
		w.Write(Message);
		
		// Write myEnum
		w.WriteTag(101);
		w.Write((int)MyEnum);
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
			case 100: // Read message
				Message = r.ReadMessage<MyPackage.MyMessage>();
				break;
			case 101: // Read myEnum
				MyEnum = (MyPackage.MyEnum)r.ReadInt();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
