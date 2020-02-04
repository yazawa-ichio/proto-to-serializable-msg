// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class DependTest : IMessage
{
	public DependMessage Msg;

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write msg
		w.WriteTag(1000);
		w.Write(Msg);
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
			case 1000: // Read msg
				Msg = r.ReadMessage<DependMessage>();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
