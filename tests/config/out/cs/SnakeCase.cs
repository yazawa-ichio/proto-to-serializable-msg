// tests/proto/test.proto
using ILib.ProtoPack;
using System.Collections.Generic;
public partial class SnakeCase : IMessage
{
	public int SnakeCaseValue;

	public partial class NestSnakeCase : IMessage
	{
		public int NestSnakeCaseValue;

		#region Serialization

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w)
		{
			// Write Map Length
			w.WriteMapHeader(1);
			
			// Write nest_snake_case_value
			w.WriteTag(1);
			w.Write(NestSnakeCaseValue);
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
				case 1: // Read nest_snake_case_value
					NestSnakeCaseValue = r.ReadInt();
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
		#endregion

	}

	#region Serialization

	/// <summary>
	/// Serialize Message
	/// </summary>
	public void Write(IWriter w)
	{
		// Write Map Length
		w.WriteMapHeader(1);
		
		// Write snake_case_value
		w.WriteTag(1);
		w.Write(SnakeCaseValue);
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
			case 1: // Read snake_case_value
				SnakeCaseValue = r.ReadInt();
				break;
			default:
				r.Skip();
				break;
			}
		}
	}
	#endregion

}
