// sample/proto/Forum.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class User {
	constructor(init, pos) {
		this.id = 0;
		this.name = null;
		this.roll = 0;
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
		
		// Write Name
		w.writeTag(2);
		w.writeString(this.name);
		
		// Write Roll
		w.writeTag(3);
		w.writeNumber(this.roll);
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
				this.name = r.readString();
				break;
			case 3:
				this.roll = r.readNumber();
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = User
