// sample/proto/Forum.proto
using ILib.ProtoPack;
using System.Collections.Generic;

namespace Forum
{
	[System.Serializable]
	public partial class ForumInfo : IMessage
	{
		public Dictionary<int, Forum.User> Users;

		[System.Serializable]
		public partial class ForumNestInfo : IMessage
		{
			public Dictionary<int, Forum.PostData> UserPost;

			#region Serialization

			/// <summary>
			/// Serialize Message
			/// </summary>
			public void Write(IWriter w)
			{
				// Write Map Length
				w.WriteMapHeader(1);
				
				// Write user_post
				w.WriteTag(1);
				if (UserPost == null)
				{
					w.WriteNil();
				}
				else
				{
					var mapLen = UserPost.Count;
					w.WriteMapHeader(mapLen);
					foreach(var _UserPostEntry in UserPost){
						w.Write(_UserPostEntry.Key);
						w.Write(_UserPostEntry.Value);
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
					case 1: // Read user_post
						if(r.NextFormatIsNull())
						{
							r.ReadNil();
							UserPost = null;
							continue;
						}
						var _UserPostLen = r.ReadMapHeader();
						UserPost = new Dictionary<int, Forum.PostData>(_UserPostLen);
						for(int mapIndex = 0; mapIndex < _UserPostLen; mapIndex++)
						{
							var _UserPostKey = default(int);
							var _UserPostValue = default(Forum.PostData);
							_UserPostKey = r.ReadInt();
							_UserPostValue = r.ReadMessage<Forum.PostData>();
							UserPost[_UserPostKey] = _UserPostValue;
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

		#region Serialization

		/// <summary>
		/// Serialize Message
		/// </summary>
		public void Write(IWriter w)
		{
			// Write Map Length
			w.WriteMapHeader(1);
			
			// Write users
			w.WriteTag(1);
			if (Users == null)
			{
				w.WriteNil();
			}
			else
			{
				var mapLen = Users.Count;
				w.WriteMapHeader(mapLen);
				foreach(var _UsersEntry in Users){
					w.Write(_UsersEntry.Key);
					w.Write(_UsersEntry.Value);
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
				case 1: // Read users
					if(r.NextFormatIsNull())
					{
						r.ReadNil();
						Users = null;
						continue;
					}
					var _UsersLen = r.ReadMapHeader();
					Users = new Dictionary<int, Forum.User>(_UsersLen);
					for(int mapIndex = 0; mapIndex < _UsersLen; mapIndex++)
					{
						var _UsersKey = default(int);
						var _UsersValue = default(Forum.User);
						_UsersKey = r.ReadInt();
						_UsersValue = r.ReadMessage<Forum.User>();
						Users[_UsersKey] = _UsersValue;
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
