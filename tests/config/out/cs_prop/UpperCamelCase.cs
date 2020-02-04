// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
/// <summary>
///  UpperCamelCase comment
/// </summary>
[System.Serializable]
public partial class UpperCamelCase : IMessage
{
	public int UpperCamelCaseField { get; set; }

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write UpperCamelCaseField
		w.WriteTag(1);
		w.Write(UpperCamelCaseField);
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
			case 1: // Read UpperCamelCaseField
				UpperCamelCaseField = r.ReadInt();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
