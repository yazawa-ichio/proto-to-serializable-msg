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
}
module.exports = DependTest
