"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
var http = __importStar(require("http"));
var proto = __importStar(require("../proto"));
console.log("start");
http.createServer(function (req, res) {
    if (req.url == null) {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('NotFound Model\n' + req.url);
        return;
    }
    var model = req.url.substring(1);
    var chunks = new Array();
    req.on('data', function (chunk) { return chunks.push(chunk); });
    req.on('end', function () {
        var buf = Buffer.concat(chunks);
        switch (model) {
            case "AllParameter":
                {
                    var p = new proto.AllParameter(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
            case "AllRepeatedParameter":
                {
                    var p = new proto.AllRepeatedParameter(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
            case "DependMessage":
                {
                    var p = new proto.DependMessage(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
            case "DependTest":
                {
                    var p = new proto.DependTest(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
            case "LowerCamelCase":
                {
                    var p = new proto.LowerCamelCase(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
            case "UpperCamelCase":
                {
                    var p = new proto.UpperCamelCase(buf);
                    //console.log(p);
                    res.writeHead(200, { 'Content-Type': 'application/octet-stream' });
                    res.end(p.pack());
                }
                break;
        }
    });
}).listen(9002, '0.0.0.0');
