//proto/Forum.proto
using ILib.ProtoPack;
using IWriter = ILib.ProtoPack.IWriter;
using IReader = ILib.ProtoPack.IReader;
using Provider = ILib.ProtoPack.InstanceProvider;

namespace Forum
{
	public partial class PostData : IMessage 
	{
		public int Id { get; set; }

		public string Message { get; set; }

		public Forum.User User { get; set; }

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w, bool skipable = true)
		{
			// Write Map Length
			if (!skipable) {
				w.WriteMapHeader(3);
			} else {
				int mapLen = 0;
				if(this.Id != default(int)) mapLen++;
				if(this.Message != default(string)) mapLen++;
				if(this.User != default(Forum.User)) mapLen++;
				w.WriteMapHeader(mapLen);
			}
			
			// Write Id
			if(!skipable || this.Id != default(int)) {
				w.WriteTag(1);
				w.Write(Id);
			}
			
			// Write Message
			if(!skipable || this.Message != default(string)) {
				w.WriteTag(2);
				w.Write(Message);
			}
			
			// Write User
			if(!skipable || this.User != default(Forum.User)) {
				var User = this.User;
				w.WriteTag(3);
				if (User == default(Forum.User))
				{
					w.WriteNil();
				}
				else
				{
					User.Write(w, skipable);
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
					Id = r.ReadInt();
					break;
				case 2:
					Message = r.ReadString();
					break;
				case 3:
					if(r.IsNull())
					{
						User = r.ReadNil<Forum.User>();
						continue;
					}
					if(!overridable || User == default(Forum.User))
					{
						User = Provider.New<Forum.User>();
					}
					User.Read(r, overridable);
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
	}
}
