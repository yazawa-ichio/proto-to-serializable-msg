import * as http from 'http';
import * as Forum from '../proto/forum';

const postForumData = new Forum.ForumData();
postForumData.data = new Array<Forum.PostData>();

http.createServer(function (req, res) {

	if (req.url == "/forum/postdata") {
		const chunks = new Array<Buffer>();
		req.on('data', chunk => chunks.push(chunk));
		req.on('end', () => {
			const reqData = new Forum.PostForumReq(Buffer.concat(chunks));
			console.log('ReqData: ', reqData);
			if (postForumData.data != null) {
				postForumData.data.push(reqData.data);
			}
			res.writeHead(200, { 'Content-Type': 'application/x-msgpack' });
			res.end(postForumData.pack());
		})
	} else {
		res.writeHead(200, { 'Content-Type': 'text/plain' });
		res.end('NotFound CMD\n' + req.url);
	}

}).listen(6438, '127.0.0.1');