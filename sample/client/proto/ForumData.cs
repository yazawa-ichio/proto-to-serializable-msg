//proto/Forum.proto
using ILib.ProtoPack;
using IWriter = ILib.ProtoPack.IWriter;
using IReader = ILib.ProtoPack.IReader;
using Provider = ILib.ProtoPack.InstanceProvider;

namespace Forum
{
	public partial class ForumData : IMessage 
	{
		public Forum.PostData[] Data { get; set; }

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
				if(this.Data != default(Forum.PostData[])) mapLen++;
				w.WriteMapHeader(mapLen);
			}
			
			// Write data
			if(!skipable || this.Data != default(Forum.PostData[])) {
				var Data = this.Data;
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
						if (Data[arrayIndex] == default(Forum.PostData))
						{
							w.WriteNil();
						}
						else
						{
							Data[arrayIndex].Write(w, skipable);
						}
					}
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
						r.ReadNil();
						this.Data = null;
						continue;
					}
					var Data = this.Data;
					var _DataLen = r.ReadArrayHeader();
					if(!overridable || Data == null)
					{
						Data = Provider.NewArray<Forum.PostData>(_DataLen);
					}
					else if(Data.Length != _DataLen)
					{
						System.Array.Resize(ref Data, _DataLen);
					}
					this.Data = Data;
					for(int arrayIndex = 0; arrayIndex < _DataLen; arrayIndex++)
					{
						if(r.IsNull())
						{
							Data[arrayIndex] = r.ReadNil<Forum.PostData>();
							continue;
						}
						if(!overridable || Data[arrayIndex] == default(Forum.PostData))
						{
							Data[arrayIndex] = Provider.New<Forum.PostData>();
						}
						Data[arrayIndex].Read(r, overridable);
					}
					break;
				default:
					r.Skip();
					break;
				}
			}
		}
	}
}
