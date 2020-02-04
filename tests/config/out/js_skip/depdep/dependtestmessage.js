// tests/proto/depend/depend.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
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
}
module.exports = DependTestMessage
