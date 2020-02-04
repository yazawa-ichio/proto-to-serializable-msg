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
	pack() {
		const w = _packer.defaultWriter;
		w.clear();
		this.write(w);
		return w.toBuffer();
	}
	unpack(buf, pos) {
		if(!Buffer.isBuffer(buf)){
			this.read(buf);
		} else {
			this.read(new _packer.ProtoReader(buf, pos));
		}
	}
	write(w) {
		// Write Map Length
		w.writeMapHeader(20);
		
		// Write value_double
		w.writeTag(1);
		w.writeDouble(this.valueDouble);
		
		// Write value_float
		w.writeTag(2);
		w.writeFloat(this.valueFloat);
		
		// Write value_int32
		w.writeTag(3);
		w.writeNumber(this.valueInt32);
		
		// Write value_int64
		w.writeTag(4);
		w.writeNumber(this.valueInt64);
		
		// Write value_uint32
		w.writeTag(5);
		w.writeNumber(this.valueUint32);
		
		// Write value_uint64
		w.writeTag(6);
		w.writeNumber(this.valueUint64);
		
		// Write value_sint32
		w.writeTag(7);
		w.writeNumber(this.valueSint32);
		
		// Write value_sint64
		w.writeTag(8);
		w.writeNumber(this.valueSint64);
		
		// Write value_fixed32
		w.writeTag(9);
		w.writeNumber(this.valueFixed32);
		
		// Write value_fixed64
		w.writeTag(10);
		w.writeNumber(this.valueFixed64);
		
		// Write value_sfixed32
		w.writeTag(11);
		w.writeNumber(this.valueSfixed32);
		
		// Write value_sfixed64
		w.writeTag(12);
		w.writeNumber(this.valueSfixed64);
		
		// Write value_bool
		w.writeTag(13);
		w.writeBool(this.valueBool);
		
		// Write value_string
		w.writeTag(14);
		w.writeString(this.valueString);
		
		// Write value_bytes
		w.writeTag(15);
		if (!this.valueBytes) {
			w.writeNil();
		} else {
			w.writeBytes(this.valueBytes);
		}
		
		// Write value_map_string
		w.writeTag(16);
		if (this.valueMapString == null) {
			w.writeNil();
		} else {
			const mapLen = this.valueMapString.size;
			w.writeMapHeader(mapLen);
			this.valueMapString.forEach(function(value, key){
				w.writeNumber(key);
				w.writeString(value);
			});
		}
		
		// Write value_map_int
		w.writeTag(17);
		if (this.valueMapInt == null) {
			w.writeNil();
		} else {
			const mapLen = this.valueMapInt.size;
			w.writeMapHeader(mapLen);
			this.valueMapInt.forEach(function(value, key){
				w.writeString(key);
				w.writeNumber(value);
			});
		}
		
		// Write value_message
		w.writeTag(18);
		if (!this.valueMessage) {
			w.writeNil();
		} else {
			this.valueMessage.write(w);
		}
		
		// Write value_map_value_message
		w.writeTag(19);
		if (this.valueMapValueMessage == null) {
			w.writeNil();
		} else {
			const mapLen = this.valueMapValueMessage.size;
			w.writeMapHeader(mapLen);
			this.valueMapValueMessage.forEach(function(value, key){
				w.writeNumber(key);
				if (!value) {
					w.writeNil();
				} else {
					value.write(w);
				}
			});
		}
		
		// Write value_testEnum
		w.writeTag(20);
		w.writeNumber(this.valueTestEnum);
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1:
				this.valueDouble = r.readDouble();
				break;
			case 2:
				this.valueFloat = r.readFloat();
				break;
			case 3:
				this.valueInt32 = r.readNumber();
				break;
			case 4:
				this.valueInt64 = r.readNumber();
				break;
			case 5:
				this.valueUint32 = r.readNumber();
				break;
			case 6:
				this.valueUint64 = r.readNumber();
				break;
			case 7:
				this.valueSint32 = r.readNumber();
				break;
			case 8:
				this.valueSint64 = r.readNumber();
				break;
			case 9:
				this.valueFixed32 = r.readNumber();
				break;
			case 10:
				this.valueFixed64 = r.readNumber();
				break;
			case 11:
				this.valueSfixed32 = r.readNumber();
				break;
			case 12:
				this.valueSfixed64 = r.readNumber();
				break;
			case 13:
				this.valueBool = r.readBool();
				break;
			case 14:
				this.valueString = r.readString();
				break;
			case 15:
				if(r.isNull()) {
					r.readNil();
					this.valueBytes = null;
				} else {
					this.valueBytes = r.readBytes();
				}
				break;
			case 16:
				if(r.isNull()) {
					r.readNil();
					this.valueMapString = null;
					continue;
				}
				const _valueMapStringLen = r.readMapHeader();
				this.valueMapString = new Map();
				for(let mapIndex = 0; mapIndex < _valueMapStringLen; mapIndex++) {
					let key;
					let value;
					key = r.readNumber();
					value = r.readString();
					this.valueMapString.set(key, value);
				}
				break;
			case 17:
				if(r.isNull()) {
					r.readNil();
					this.valueMapInt = null;
					continue;
				}
				const _valueMapIntLen = r.readMapHeader();
				this.valueMapInt = new Map();
				for(let mapIndex = 0; mapIndex < _valueMapIntLen; mapIndex++) {
					let key;
					let value;
					key = r.readString();
					value = r.readNumber();
					this.valueMapInt.set(key, value);
				}
				break;
			case 18:
				if(r.isNull()) {
					r.readNil();
					this.valueMessage = null;
				} else {
					this.valueMessage = new _proto.EmptyMessage(false);
					this.valueMessage.read(r);
				}
				break;
			case 19:
				if(r.isNull()) {
					r.readNil();
					this.valueMapValueMessage = null;
					continue;
				}
				const _valueMapValueMessageLen = r.readMapHeader();
				this.valueMapValueMessage = new Map();
				for(let mapIndex = 0; mapIndex < _valueMapValueMessageLen; mapIndex++) {
					let key;
					let value;
					key = r.readNumber();
					if(r.isNull()) {
						r.readNil();
						value = null;
					} else {
						value = new _proto.DependTest(false);
						value.read(r);
					}
					this.valueMapValueMessage.set(key, value);
				}
				break;
			case 20:
				this.valueTestEnum = r.readNumber();
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = AllParameter
