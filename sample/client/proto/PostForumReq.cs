//proto/Forum.proto
using ILib.ProtoPack;
using IWriter = ILib.ProtoPack.IWriter;
using IReader = ILib.ProtoPack.IReader;
using Provider = ILib.ProtoPack.InstanceProvider;

namespace Forum
{
	public partial class PostForumReq : IMessage 
	{
		public Forum.PostData Data { get; set; }

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w, bool skipable = true)
		{
			// Write Map Length
			if (!skipable) {
				w.WriteMapHeader(1);
			} else {
				int mapLen = 0;
				if(this.Data != default(Forum.PostData)) mapLen++;
				w.WriteMapHeader(mapLen);
			}
			
			// Write data
			if(!skipable || this.Data != default(Forum.PostData)) {
				var Data = this.Data;
				w.WriteTag(1);
				if (Data == default(Forum.PostData))
				{
					w.WriteNil();
				}
				else
				{
					Data.Write(w, skipable);
				}
			}
		}

		/// <summary>
		/// Deserialize Message
		/// </summary>
		public void Read(IReader r, bool overridable = false)
		{
			// Read Map Length
			var mapLen = r.ReadMapHeader();
			uint tag = 0;
			int index = 0;

			while ((tag = r.ReadTag(index++, mapLen)) != 0)
			{
				switch(tag) {
				case 1:
					if(r.IsNull())
					{
						Data = r.ReadNil<Forum.PostData>();
						continue;
					}
					if(!overridable || Data == default(Forum.PostData))
					{
						Data = Provider.New<Forum.PostData>();
					}
					Data.Read(r, overridable);
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
	}
}
