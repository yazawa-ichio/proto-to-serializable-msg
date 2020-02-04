// sample/proto/Forum.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace Forum
{
	[System.Serializable]
	public partial class ForumData : IMessage
	{
		public Forum.PostData[] Data;

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
			if (Data == null)
			{
				w.WriteNil();
			}
			else
			{
				var arrayLen = Data.Length;
				w.WriteArrayHeader(arrayLen);
				for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
				{
					w.Write(Data[arrayIndex]);
				}
			}
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
					if(r.NextFormatIsNull())
					{
						r.ReadNil();
						this.Data = null;
						continue;
					}
					var _DataLen = r.ReadArrayHeader();
					Data = new Forum.PostData[_DataLen];
					for(int arrayIndex = 0; arrayIndex < _DataLen; arrayIndex++)
					{
						Data[arrayIndex] = r.ReadMessage<Forum.PostData>();
					}
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
