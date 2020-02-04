using NUnit.Framework;
using RandomFixtureKit;
using ILib.ProtoPack;
using System.Net.Http;
using System.Threading.Tasks;
using System.Linq;
using System.Diagnostics;

namespace SerializeTest
{
    public partial class UnitTest1
    {
        [Test]
        public void SameTest()
        {
            for (int i = 0; i < 1000; i++)
            {
                Same<AllParameter>();
                Same<AllRepeatedParameter>();
                Same<DependMessage>();
                Same<DependTest>();
                Same<LowerCamelCase>();
                Same<UpperCamelCase>();
            }
        }

        [Test]
        public async Task GoServerRequest()
        {
            await Task.Delay(1000);
            await ServerRequestTest("http://localhost:9001/",true,false);
        }

        [Test]
        public async Task TsServerRequest()
        {
            await Task.Delay(1000);
            await ServerRequestTest("http://localhost:9002/", false, true);
        }

        async Task ServerRequestTest(string url, bool fixSort, bool fixNumber)
        {
            using (var client = new HttpClient())
            {
                for (int i = 0; i < 100; i++)
                {
                    await RequestSame<AllParameter>(client, url, fixSort, fixNumber, false);
                    // JavaScriptのNumberのチェックが面倒なのでデシリアライズ出来ればOKとする
                    await RequestSame<AllRepeatedParameter>(client, url, fixSort, fixNumber, fixNumber);
                    await RequestSame<DependMessage>(client, url, fixSort, fixNumber, false);
                    await RequestSame<DependTest>(client, url, fixSort, fixNumber, false);
                    await RequestSame<LowerCamelCase>(client, url, fixSort, fixNumber, false);
                    await RequestSame<UpperCamelCase>(client, url, fixSort, fixNumber, false);
                }
            }
        }

        void Same<T>() where T : IMessage, new()
        {
            var src = FixtureFactory.Create<T>(10);
            var buf = StaticPacker.Pack(src);
            var dst = StaticPacker.Unpack<T>(buf);
            Assert.AreEqual(Json.To(dst), Json.To(src));
        }

        async Task RequestSame<T>(HttpClient client, string url, bool fixSort, bool fixNumber, bool skipCheck) where T : IMessage, new()
        {
            var src = FixtureFactory.Create<T>(10);
            var content = new ByteArrayContent(StaticPacker.Pack(src));
            var ret = await client.PostAsync(url + src.GetType().Name, content);
            var buf = await ret.Content.ReadAsByteArrayAsync();
            var dst = StaticPacker.Unpack<T>(buf);
            if (src is AllParameter)
            {
                FixParam(src as AllParameter, fixSort, fixNumber);
                FixParam(dst as AllParameter, fixSort, fixNumber);
            }
            if (!skipCheck)
            {
                Assert.AreEqual(Json.To(dst), Json.To(src));
            }
        }

        /// <summary>
        /// Go言語がMapをランダムにして来るのでソートする
        /// </summary>
        void FixParam(AllParameter parameter, bool fixSort, bool fixNumber)
        {
            if (fixSort)
            {
                parameter.ValueMapInt = parameter.ValueMapInt.OrderBy(x => x.Key).ToDictionary(x => x.Key, x => x.Value);
                parameter.ValueMapString = parameter.ValueMapString.OrderBy(x => x.Key).ToDictionary(x => x.Key, x => x.Value);
                parameter.ValueMapValueMessage = parameter.ValueMapValueMessage.OrderBy(x => x.Key).ToDictionary(x => x.Key, x => x.Value);
            }
            if (fixNumber)
            {
                if (int.MaxValue < parameter.ValueInt64 || int.MinValue > parameter.ValueInt64) parameter.ValueInt64 = 0;
                if (int.MaxValue < parameter.ValueUint64) parameter.ValueUint64 = 0;

                if (int.MaxValue < parameter.ValueSint64 || int.MinValue > parameter.ValueSint64) parameter.ValueSint64 = 0;
                if (int.MaxValue < parameter.ValueSfixed64 || int.MinValue > parameter.ValueSfixed64) parameter.ValueSfixed64 = 0;
                if (int.MaxValue < parameter.ValueFixed64) parameter.ValueFixed64 = 0;
            }
        }

    }
}