"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const http = require("http");
const ForumReq = require("../proto/forum.postforumreq");
const ForumData = require("../proto/forum.forumdata");
const postForumData = new ForumData();
postForumData.data = new Array();
http.createServer(function (req, res) {
    if (req.url == "/forum/postdata") {
        const chunks = new Array();
        req.on('data', chunk => chunks.push(chunk));
        req.on('end', () => {
            const reqData = new ForumReq(Buffer.concat(chunks));
            console.log('ReqData: ', reqData);
            if (postForumData.data != null) {
                postForumData.data.push(reqData.data);
            }
            res.writeHead(200, { 'Content-Type': 'application/x-msgpack' });
            res.end(postForumData.pack());
        });
    }
    else {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('NotFound CMD\n' + req.url);
    }
}).listen(6438, '127.0.0.1');
