// tests/proto/depend/depend.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace DepDep
{
	public partial class DependTestMessage : IMessage
	{
		public PackageMessage Message;

		public DependMessage DepDep;

		#region Serialization

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w)
		{
			// Write Map Length
			w.WriteMapHeader(2);
			
			// Write message
			w.WriteTag(100);
			w.Write(Message);
			
			// Write dep_dep
			w.WriteTag(101);
			w.Write(DepDep);
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
				case 100: // Read message
					Message = r.ReadMessage<PackageMessage>();
					break;
				case 101: // Read dep_dep
					DepDep = r.ReadMessage<DependMessage>();
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
