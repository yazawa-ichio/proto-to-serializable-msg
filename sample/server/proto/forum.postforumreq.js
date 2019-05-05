//proto/Forum.proto
"use strict";
var packer = require('proto-msgpack');
require('./forum.postdata.js');
//add proto
if (!packer.proto) {
	packer.proto = {};
}
if (!packer.proto.Forum) {
	packer.proto.Forum = {};
}
module.exports = packer.proto.Forum.PostForumReq = class PostForumReq {
	constructor(init, pos) {
		if(init == null || init == true){
			this.data = new packer.proto.Forum.PostData();
		} else {
			this.data = null;
		}
		if(Buffer.isBuffer(init)){
			this.read(new packer.ProtoReader(init, pos));
		}
	}
	pack() {
		const w = packer.defaultWriter;
		w.clear();
		this.write(w);
		return w.toBuffer();
	}
	unpack(buf, pos) {
		if(!Buffer.isBuffer(buf)){
			this.read(buf);
		} else {
			this.read(new packer.ProtoReader(buf, pos));
		}
	}
	write(w) {
		// Write Map Length
		w.writeMapHeader(1);
		
		// Write data
		w.writeTag(1);
		if (!this.data) {
			w.writeNil();
		} else {
			this.data.write(w);
		}
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1:
				if(r.isNull()) {
					r.readNil();
					this.data = null;
					continue;
				}
				this.data = new packer.proto.Forum.PostData();
				this.data.read(r);
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
