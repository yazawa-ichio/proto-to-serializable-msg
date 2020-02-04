// tests/proto/import.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class PackageMessage {
	constructor(init, pos) {
		if(init == null || init == true){
			this.message = new _proto.MyPackage.MyMessage();
		} else {
			this.message = null;
		}
		this.myEnum = 0;
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
		w.writeMapHeader(2);
		
		// Write message
		w.writeTag(100);
		if (!this.message) {
			w.writeNil();
		} else {
			this.message.write(w);
		}
		
		// Write myEnum
		w.writeTag(101);
		w.writeNumber(this.myEnum);
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 100:
				if(r.isNull()) {
					r.readNil();
					this.message = null;
				} else {
					this.message = new _proto.MyPackage.MyMessage(false);
					this.message.read(r);
				}
				break;
			case 101:
				this.myEnum = r.readNumber();
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = PackageMessage
