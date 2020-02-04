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
}
module.exports = PackageMessage
