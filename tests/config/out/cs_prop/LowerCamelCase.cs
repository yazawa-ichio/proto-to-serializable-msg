// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
/// <summary>
///  lowerCamelCase comment
/// </summary>
[System.Serializable]
public partial class LowerCamelCase : IMessage
{
	public int LowerCamelCaseField { get; set; }

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write lowerCamelCaseField
		w.WriteTag(1);
		w.Write(LowerCamelCaseField);
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
			case 1: // Read lowerCamelCaseField
				LowerCamelCaseField = r.ReadInt();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
