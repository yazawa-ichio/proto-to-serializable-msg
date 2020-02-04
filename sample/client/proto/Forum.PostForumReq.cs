// sample/proto/Forum.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace Forum
{
	[System.Serializable]
	public partial class PostForumReq : IMessage
	{
		/// <summary>
		/// Data
		/// </summary>
		public Forum.PostData Data;

		#region Serialization

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w)
		{
			// Write Map Length
			w.WriteMapHeader(1);
			
			// Write data
			w.WriteTag(1);
			w.Write(Data);
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
				case 1: // Read data
					Data = r.ReadMessage<Forum.PostData>();
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
