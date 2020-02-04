// sample/proto/Forum.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class PostForumReq {
	constructor(init, pos) {
		if(init == null || init == true){
			this.data = new _proto.Forum.PostData();
		} else {
			this.data = null;
		}
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
	pack() {
		const w = _packer.defaultWriter;
		w.clear();
		this.write(w);
		return w.toBuffer();
	}
	unpack(buf, pos) {
		if(!Buffer.isBuffer(buf)){
			this.read(buf);
		} else {
			this.read(new _packer.ProtoReader(buf, pos));
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
				} else {
					this.data = new _proto.Forum.PostData(false);
					this.data.read(r);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = PostForumReq
