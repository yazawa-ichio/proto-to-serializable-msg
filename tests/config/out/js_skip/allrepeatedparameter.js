// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class AllRepeatedParameter {
	constructor(init, pos) {
		this.valueDouble = null;
		this.valueFloat = null;
		this.valueInt32 = null;
		this.valueInt64 = null;
		this.valueUint32 = null;
		this.valueUint64 = null;
		this.valueSint32 = null;
		this.valueSint64 = null;
		this.valueFixed32 = null;
		this.valueFixed64 = null;
		this.valueSfixed32 = null;
		this.valueSfixed64 = null;
		this.valueBool = null;
		this.valueString = null;
		this.valueBytes = null;
		this.valueNestMessage = null;
		this.valueTestEnum = null;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = AllRepeatedParameter
