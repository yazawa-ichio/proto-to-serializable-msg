// sample/proto/Forum.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace Forum
{
	[System.Serializable]
	public partial class User : IMessage
	{
		public int Id;

		public string Name;

		public Forum.Roll Roll;

		#region Serialization

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w)
		{
			// Write Map Length
			w.WriteMapHeader(3);
			
			// Write Id
			w.WriteTag(1);
			w.Write(Id);
			
			// Write Name
			w.WriteTag(2);
			w.Write(Name);
			
			// Write Roll
			w.WriteTag(3);
			w.Write((int)Roll);
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
				case 1: // Read Id
					Id = r.ReadInt();
					break;
				case 2: // Read Name
					Name = r.ReadString();
					break;
				case 3: // Read Roll
					Roll = (Forum.Roll)r.ReadInt();
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
		#endregion

	}
}
