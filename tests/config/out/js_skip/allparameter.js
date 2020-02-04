// tests/proto/test.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./index.js');
class AllParameter {
	constructor(init, pos) {
		this.valueDouble = 0;
		this.valueFloat = 0;
		this.valueInt32 = 0;
		this.valueInt64 = 0;
		this.valueUint32 = 0;
		this.valueUint64 = 0;
		this.valueSint32 = 0;
		this.valueSint64 = 0;
		this.valueFixed32 = 0;
		this.valueFixed64 = 0;
		this.valueSfixed32 = 0;
		this.valueSfixed64 = 0;
		this.valueBool = null;
		this.valueString = null;
		this.valueBytes = null;
		this.valueMapString = null;
		this.valueMapInt = null;
		if(init == null || init == true){
			this.valueMessage = new _proto.EmptyMessage();
		} else {
			this.valueMessage = null;
		}
		this.valueMapValueMessage = null;
		this.valueTestEnum = 0;
		if(Buffer.isBuffer(init)){
			this.read(new _packer.ProtoReader(init, pos));
		}
	}
}
module.exports = AllParameter
