// sample/proto/Forum.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace Forum
{
	[System.Serializable]
	public partial class PostData : IMessage
	{
		public int Id;

		public string Message;

		public Forum.User User;

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
			
			// Write Message
			w.WriteTag(2);
			w.Write(Message);
			
			// Write User
			w.WriteTag(3);
			w.Write(User);
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
				case 2: // Read Message
					Message = r.ReadString();
					break;
				case 3: // Read User
					User = r.ReadMessage<Forum.User>();
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
