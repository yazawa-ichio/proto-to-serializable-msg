using System.IO;
using System.Runtime.Serialization.Json;
using System;

namespace SerializeTest
{
	public class Json
    {
        public static string To(object obj)
        {
            using (var ms = new MemoryStream())
            using (var sr = new StreamReader(ms))
            {
                var serializer = new DataContractJsonSerializer(obj.GetType());
                serializer.WriteObject(ms, obj);
                ms.Position = 0;
                var json = sr.ReadToEnd();
                Console.WriteLine(obj + " to " + json);
                return json;
            }
        }
    }
}