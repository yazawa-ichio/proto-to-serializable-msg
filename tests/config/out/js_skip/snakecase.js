// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class SnakeCase {
	constructor(init, pos) {
		this.snakeCaseValue = 0;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = SnakeCase
SnakeCase.NestSnakeCase = class NestSnakeCase {
	constructor(init, pos) {
		this.nestSnakeCaseValue = 0;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
