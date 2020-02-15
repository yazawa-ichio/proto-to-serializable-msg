const http = require('http');
const proto = require("./proto");


const postForumData = new proto.Forum.ForumData();
postForumData.data = []

http.createServer(function (req, res) {

	if (req.url == "/forum/postdata") {
		const chunks = [];
		req.on('data', chunk => chunks.push(chunk));
		req.on('end', () => {
			const reqData = new proto.Forum.PostForumReq(Buffer.concat(chunks));
			console.log('ReqData: ', reqData);
			postForumData.data.push(reqData.data);
			res.writeHead(200, { 'Content-Type': 'application/x-msgpack' });
			res.end(postForumData.pack());
		})
	} else {
		res.writeHead(200, { 'Content-Type': 'text/plain' });
		res.end('NotFound CMD\n' + req.url);
	}

}).listen(6438, '127.0.0.1');
