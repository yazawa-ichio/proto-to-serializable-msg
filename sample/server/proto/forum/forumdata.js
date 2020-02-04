// sample/proto/Forum.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class ForumData {
	constructor(init, pos) {
		this.data = null;
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
		if (this.data == null) {
			w.writeNil();
		} else {
			const arrayLen = this.data.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				if (!this.data[arrayIndex]) {
					w.writeNil();
				} else {
					this.data[arrayIndex].write(w);
				}
			}
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
				const _dataLen = r.readArrayHeader();
				this.data = new Array(_dataLen);
				for(let arrayIndex = 0; arrayIndex < _dataLen; arrayIndex++) {
					if(r.isNull()) {
						r.readNil();
						this.data[arrayIndex] = null;
					} else {
						this.data[arrayIndex] = new _proto.Forum.PostData(false);
						this.data[arrayIndex].read(r);
					}
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = ForumData
