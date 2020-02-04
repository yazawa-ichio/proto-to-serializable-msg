import * as http from 'http';
import * as proto from '../proto';
import * as mypackage from '../proto/mypackage';

console.log("start")
http.createServer(function (req, res) {
	if (req.url == null)
	{
		res.writeHead(200, { 'Content-Type': 'text/plain' });
		res.end('NotFound Model\n' + req.url);
		return;
	}
	const model = req.url.substring(1);
	const chunks = new Array<Buffer>();
	req.on('data', chunk => chunks.push(chunk));
	req.on('end', () => {
		const buf = Buffer.concat(chunks);
		switch (model) {
			case "AllParameter":
				{
					const p = new proto.AllParameter(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
			case "AllRepeatedParameter":
				{
					const p = new proto.AllRepeatedParameter(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
			case "DependMessage":
				{
					const p = new proto.DependMessage(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
			case "DependTest":
				{
					const p = new proto.DependTest(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
			case "LowerCamelCase":
				{
					const p = new proto.LowerCamelCase(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
			case "UpperCamelCase":
				{
					const p = new proto.UpperCamelCase(buf);
					//console.log(p);
					res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
					res.end(p.pack());
				}
				break
		}
	});
}).listen(9002, '0.0.0.0');