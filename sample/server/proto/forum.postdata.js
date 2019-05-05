//proto/Forum.proto
"use strict";
var packer = require('proto-msgpack');
require('./forum.user.js');
//add proto
if (!packer.proto) {
	packer.proto = {};
}
if (!packer.proto.Forum) {
	packer.proto.Forum = {};
}
module.exports = packer.proto.Forum.PostData = class PostData {
	constructor(init, pos) {
		this.Id = 0;
		this.Message = null;
		if(init == null || init == true){
			this.User = new packer.proto.Forum.User();
		} else {
			this.User = null;
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
		w.writeMapHeader(3);
		
		// Write Id
		w.writeTag(1);
		w.writeNumber(this.Id);
		
		// Write Message
		w.writeTag(2);
		w.writeString(this.Message);
		
		// Write User
		w.writeTag(3);
		if (!this.User) {
			w.writeNil();
		} else {
			this.User.write(w);
		}
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1:
				this.Id = r.readNumber();
				break;
			case 2:
				this.Message = r.readString();
				break;
			case 3:
				if(r.isNull()) {
					r.readNil();
					this.User = null;
					continue;
				}
				this.User = new packer.proto.Forum.User();
				this.User.read(r);
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
