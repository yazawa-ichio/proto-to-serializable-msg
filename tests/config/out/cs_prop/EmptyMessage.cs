// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class EmptyMessage : IMessage
{

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(0);
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
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
