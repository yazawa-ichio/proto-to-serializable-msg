// sample/proto/Depend.proto
using ILib.ProtoPack;
using System.Collections.Generic;
[System.Serializable]
public partial class DependTest : IMessage
{
	public Forum.PostData Dep;

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write dep
		w.WriteTag(1);
		w.Write(Dep);
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
			case 1: // Read dep
				Dep = r.ReadMessage<Forum.PostData>();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
