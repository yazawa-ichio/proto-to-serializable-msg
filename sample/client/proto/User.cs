//proto/Forum.proto
using ILib.ProtoPack;
using IWriter = ILib.ProtoPack.IWriter;
using IReader = ILib.ProtoPack.IReader;
using Provider = ILib.ProtoPack.InstanceProvider;

namespace Forum
{
	public partial class User : IMessage 
	{
		public int Id { get; set; }

		public string Name { get; set; }

		public Forum.Roll Roll { get; set; }

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
				if(this.Name != default(string)) mapLen++;
				if(this.Roll != default(Forum.Roll)) mapLen++;
				w.WriteMapHeader(mapLen);
			}
			
			// Write Id
			if(!skipable || this.Id != default(int)) {
				w.WriteTag(1);
				w.Write(Id);
			}
			
			// Write Name
			if(!skipable || this.Name != default(string)) {
				w.WriteTag(2);
				w.Write(Name);
			}
			
			// Write Roll
			if(!skipable || this.Roll != default(Forum.Roll)) {
				w.WriteTag(3);
				w.Write((int)Roll);
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
					Name = r.ReadString();
					break;
				case 3:
					Roll = (Forum.Roll)r.ReadInt();
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
	}
}
