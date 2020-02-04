// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
/**
 *  UpperCamelCase comment
 */
class UpperCamelCase {
	constructor(init, pos) {
		this.upperCamelCaseField = 0;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = UpperCamelCase
