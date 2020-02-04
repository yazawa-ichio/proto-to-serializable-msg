// tests/proto/import.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class DependMessage {
	constructor(init, pos) {
		this.text = null;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = DependMessage
