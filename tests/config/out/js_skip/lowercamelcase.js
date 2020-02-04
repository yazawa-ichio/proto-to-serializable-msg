// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
/**
 *  lowerCamelCase comment
 */
class LowerCamelCase {
	constructor(init, pos) {
		this.lowerCamelCaseField = 0;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = LowerCamelCase
