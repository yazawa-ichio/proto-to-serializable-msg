// tests/proto/import.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class DependMessage : IMessage
{
	public string Text { get; set; }

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write text
		w.WriteTag(500);
		w.Write(Text);
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
			case 500: // Read text
				Text = r.ReadString();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
