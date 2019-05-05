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
module.exports = packer.proto.Forum.ForumData = class ForumData {
	constructor(init, pos) {
		this.data = null;
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
						continue;
					}
					this.data[arrayIndex] = new packer.proto.Forum.PostData();
					this.data[arrayIndex].read(r);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
