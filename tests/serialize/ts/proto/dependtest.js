// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class DependTest {
	constructor(init, pos) {
		if(init == null || init == true){
			this.msg = new _proto.DependMessage();
		} else {
			this.msg = null;
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
		
		// Write msg
		w.writeTag(1000);
		if (!this.msg) {
			w.writeNil();
		} else {
			this.msg.write(w);
		}
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1000:
				if(r.isNull()) {
					r.readNil();
					this.msg = null;
				} else {
					this.msg = new _proto.DependMessage(false);
					this.msg.read(r);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = DependTest
