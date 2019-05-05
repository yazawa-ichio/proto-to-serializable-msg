//proto/Forum.proto
"use strict";
var packer = require('proto-msgpack');
require('./forum.roll.js');
//add proto
if (!packer.proto) {
	packer.proto = {};
}
if (!packer.proto.Forum) {
	packer.proto.Forum = {};
}
module.exports = packer.proto.Forum.User = class User {
	constructor(init, pos) {
		this.Id = 0;
		this.Name = null;
		this.Roll = 0;
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
		
		// Write Name
		w.writeTag(2);
		w.writeString(this.Name);
		
		// Write Roll
		w.writeTag(3);
		w.writeNumber(this.Roll);
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
				this.Name = r.readString();
				break;
			case 3:
				this.Roll = r.readNumber();
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
