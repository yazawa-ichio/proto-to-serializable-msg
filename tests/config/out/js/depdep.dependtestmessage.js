// tests/proto/depend/depend.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class DependTestMessage {
	constructor(init, pos) {
		if(init == null || init == true){
			this.message = new _proto.PackageMessage();
		} else {
			this.message = null;
		}
		if(init == null || init == true){
			this.depDep = new _proto.DependMessage();
		} else {
			this.depDep = null;
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
		w.writeMapHeader(2);
		
		// Write message
		w.writeTag(100);
		if (!this.message) {
			w.writeNil();
		} else {
			this.message.write(w);
		}
		
		// Write dep_dep
		w.writeTag(101);
		if (!this.depDep) {
			w.writeNil();
		} else {
			this.depDep.write(w);
		}
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
					this.message = new _proto.PackageMessage(false);
					this.message.read(r);
				}
				break;
			case 101:
				if(r.isNull()) {
					r.readNil();
					this.depDep = null;
				} else {
					this.depDep = new _proto.DependMessage(false);
					this.depDep.read(r);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = DependTestMessage
