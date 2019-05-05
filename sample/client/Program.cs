using System;
using Forum;
using System.Net;
using System.Net.Http;
using ILib.ProtoPack;

namespace ProtoSample
{
	class Program
	{

		static void Main(string[] args)
		{
			var rand = new System.Random();
			User user = new User();
			Console.WriteLine("Input Name:");
			user.Name = Console.ReadLine();
			user.Roll = Roll.Guest;
			user.Id = rand.Next();
			int id = 0;

			while (true)
			{
				Console.WriteLine("CMD: [postdata] [end]");
				switch (Console.ReadLine())
				{
					case "end":
						return;
					case "postdata":
					default:
						var postData = new PostData();
						postData.Id = id++;
						Console.WriteLine("Input Message:");
						postData.Message = Console.ReadLine();
						postData.User = user;
						var req = new PostForumReq() { Data = postData };
						var res = Reqest<ForumData>("postdata", req);
						Console.WriteLine(res.Data.Length);
						break;
				}
			}

		}
		static T Reqest<T>(string cmd, IMessage message) where T : IMessage, new()
		{
			using(var client = new WebClient()){
				var data = message.Pack();
				client.Headers[HttpRequestHeader.ContentType] = "application/x-msgpack";
				var res = client.UploadData("http://127.0.0.1:6438/forum/" + cmd, data);
				return res.Unpack<T>();
			}
		}
	}
}
