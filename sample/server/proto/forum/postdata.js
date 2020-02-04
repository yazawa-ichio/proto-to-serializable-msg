// sample/proto/Forum.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class PostData {
	constructor(init, pos) {
		this.id = 0;
		this.message = null;
		if(init == null || init == true){
			this.user = new _proto.Forum.User();
		} else {
			this.user = null;
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
		w.writeMapHeader(3);
		
		// Write Id
		w.writeTag(1);
		w.writeNumber(this.id);
		
		// Write Message
		w.writeTag(2);
		w.writeString(this.message);
		
		// Write User
		w.writeTag(3);
		if (!this.user) {
			w.writeNil();
		} else {
			this.user.write(w);
		}
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1:
				this.id = r.readNumber();
				break;
			case 2:
				this.message = r.readString();
				break;
			case 3:
				if(r.isNull()) {
					r.readNil();
					this.user = null;
				} else {
					this.user = new _proto.Forum.User(false);
					this.user.read(r);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = PostData
