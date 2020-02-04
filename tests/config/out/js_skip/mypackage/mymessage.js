// tests/proto/packagetest/package.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class MyMessage {
	constructor(init, pos) {
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = MyMessage
