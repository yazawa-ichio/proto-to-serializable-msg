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
		w.writeMapHeader(17);
		
		// Write value_double
		w.writeTag(1);
		if (this.valueDouble == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueDouble.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeDouble(this.valueDouble[arrayIndex]);
			}
		}
		
		// Write value_float
		w.writeTag(2);
		if (this.valueFloat == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueFloat.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeFloat(this.valueFloat[arrayIndex]);
			}
		}
		
		// Write value_int32
		w.writeTag(3);
		if (this.valueInt32 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueInt32.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueInt32[arrayIndex]);
			}
		}
		
		// Write value_int64
		w.writeTag(4);
		if (this.valueInt64 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueInt64.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueInt64[arrayIndex]);
			}
		}
		
		// Write value_uint32
		w.writeTag(5);
		if (this.valueUint32 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueUint32.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueUint32[arrayIndex]);
			}
		}
		
		// Write value_uint64
		w.writeTag(6);
		if (this.valueUint64 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueUint64.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueUint64[arrayIndex]);
			}
		}
		
		// Write value_sint32
		w.writeTag(7);
		if (this.valueSint32 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueSint32.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueSint32[arrayIndex]);
			}
		}
		
		// Write value_sint64
		w.writeTag(8);
		if (this.valueSint64 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueSint64.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueSint64[arrayIndex]);
			}
		}
		
		// Write value_fixed32
		w.writeTag(9);
		if (this.valueFixed32 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueFixed32.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueFixed32[arrayIndex]);
			}
		}
		
		// Write value_fixed64
		w.writeTag(10);
		if (this.valueFixed64 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueFixed64.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueFixed64[arrayIndex]);
			}
		}
		
		// Write value_sfixed32
		w.writeTag(11);
		if (this.valueSfixed32 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueSfixed32.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueSfixed32[arrayIndex]);
			}
		}
		
		// Write value_sfixed64
		w.writeTag(12);
		if (this.valueSfixed64 == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueSfixed64.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueSfixed64[arrayIndex]);
			}
		}
		
		// Write value_bool
		w.writeTag(13);
		if (this.valueBool == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueBool.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeBool(this.valueBool[arrayIndex]);
			}
		}
		
		// Write value_string
		w.writeTag(14);
		if (this.valueString == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueString.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeString(this.valueString[arrayIndex]);
			}
		}
		
		// Write value_bytes
		w.writeTag(15);
		if (this.valueBytes == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueBytes.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				if (!this.valueBytes[arrayIndex]) {
					w.writeNil();
				} else {
					w.writeBytes(this.valueBytes[arrayIndex]);
				}
			}
		}
		
		// Write ValueNestMessage
		w.writeTag(18);
		if (this.valueNestMessage == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueNestMessage.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				if (!this.valueNestMessage[arrayIndex]) {
					w.writeNil();
				} else {
					this.valueNestMessage[arrayIndex].write(w);
				}
			}
		}
		
		// Write ValueTestEnum
		w.writeTag(20);
		if (this.valueTestEnum == null) {
			w.writeNil();
		} else {
			const arrayLen = this.valueTestEnum.length;
			w.writeArrayHeader(arrayLen);
			for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)
			{
				w.writeNumber(this.valueTestEnum[arrayIndex]);
			}
		}
	}
	read(r) {
		// Read Map Length
		const mapLen = r.readMapHeader();

		for(let i = 0; i < mapLen; i++) {
			const tag = r.readTag();
			switch(tag) {
			case 1:
				if(r.isNull()) {
					r.readNil();
					this.valueDouble = null;
					continue;
				}
				const _valueDoubleLen = r.readArrayHeader();
				this.valueDouble = new Array(_valueDoubleLen);
				for(let arrayIndex = 0; arrayIndex < _valueDoubleLen; arrayIndex++) {
					this.valueDouble[arrayIndex] = r.readDouble();
				}
				break;
			case 2:
				if(r.isNull()) {
					r.readNil();
					this.valueFloat = null;
					continue;
				}
				const _valueFloatLen = r.readArrayHeader();
				this.valueFloat = new Array(_valueFloatLen);
				for(let arrayIndex = 0; arrayIndex < _valueFloatLen; arrayIndex++) {
					this.valueFloat[arrayIndex] = r.readFloat();
				}
				break;
			case 3:
				if(r.isNull()) {
					r.readNil();
					this.valueInt32 = null;
					continue;
				}
				const _valueInt32Len = r.readArrayHeader();
				this.valueInt32 = new Array(_valueInt32Len);
				for(let arrayIndex = 0; arrayIndex < _valueInt32Len; arrayIndex++) {
					this.valueInt32[arrayIndex] = r.readNumber();
				}
				break;
			case 4:
				if(r.isNull()) {
					r.readNil();
					this.valueInt64 = null;
					continue;
				}
				const _valueInt64Len = r.readArrayHeader();
				this.valueInt64 = new Array(_valueInt64Len);
				for(let arrayIndex = 0; arrayIndex < _valueInt64Len; arrayIndex++) {
					this.valueInt64[arrayIndex] = r.readNumber();
				}
				break;
			case 5:
				if(r.isNull()) {
					r.readNil();
					this.valueUint32 = null;
					continue;
				}
				const _valueUint32Len = r.readArrayHeader();
				this.valueUint32 = new Array(_valueUint32Len);
				for(let arrayIndex = 0; arrayIndex < _valueUint32Len; arrayIndex++) {
					this.valueUint32[arrayIndex] = r.readNumber();
				}
				break;
			case 6:
				if(r.isNull()) {
					r.readNil();
					this.valueUint64 = null;
					continue;
				}
				const _valueUint64Len = r.readArrayHeader();
				this.valueUint64 = new Array(_valueUint64Len);
				for(let arrayIndex = 0; arrayIndex < _valueUint64Len; arrayIndex++) {
					this.valueUint64[arrayIndex] = r.readNumber();
				}
				break;
			case 7:
				if(r.isNull()) {
					r.readNil();
					this.valueSint32 = null;
					continue;
				}
				const _valueSint32Len = r.readArrayHeader();
				this.valueSint32 = new Array(_valueSint32Len);
				for(let arrayIndex = 0; arrayIndex < _valueSint32Len; arrayIndex++) {
					this.valueSint32[arrayIndex] = r.readNumber();
				}
				break;
			case 8:
				if(r.isNull()) {
					r.readNil();
					this.valueSint64 = null;
					continue;
				}
				const _valueSint64Len = r.readArrayHeader();
				this.valueSint64 = new Array(_valueSint64Len);
				for(let arrayIndex = 0; arrayIndex < _valueSint64Len; arrayIndex++) {
					this.valueSint64[arrayIndex] = r.readNumber();
				}
				break;
			case 9:
				if(r.isNull()) {
					r.readNil();
					this.valueFixed32 = null;
					continue;
				}
				const _valueFixed32Len = r.readArrayHeader();
				this.valueFixed32 = new Array(_valueFixed32Len);
				for(let arrayIndex = 0; arrayIndex < _valueFixed32Len; arrayIndex++) {
					this.valueFixed32[arrayIndex] = r.readNumber();
				}
				break;
			case 10:
				if(r.isNull()) {
					r.readNil();
					this.valueFixed64 = null;
					continue;
				}
				const _valueFixed64Len = r.readArrayHeader();
				this.valueFixed64 = new Array(_valueFixed64Len);
				for(let arrayIndex = 0; arrayIndex < _valueFixed64Len; arrayIndex++) {
					this.valueFixed64[arrayIndex] = r.readNumber();
				}
				break;
			case 11:
				if(r.isNull()) {
					r.readNil();
					this.valueSfixed32 = null;
					continue;
				}
				const _valueSfixed32Len = r.readArrayHeader();
				this.valueSfixed32 = new Array(_valueSfixed32Len);
				for(let arrayIndex = 0; arrayIndex < _valueSfixed32Len; arrayIndex++) {
					this.valueSfixed32[arrayIndex] = r.readNumber();
				}
				break;
			case 12:
				if(r.isNull()) {
					r.readNil();
					this.valueSfixed64 = null;
					continue;
				}
				const _valueSfixed64Len = r.readArrayHeader();
				this.valueSfixed64 = new Array(_valueSfixed64Len);
				for(let arrayIndex = 0; arrayIndex < _valueSfixed64Len; arrayIndex++) {
					this.valueSfixed64[arrayIndex] = r.readNumber();
				}
				break;
			case 13:
				if(r.isNull()) {
					r.readNil();
					this.valueBool = null;
					continue;
				}
				const _valueBoolLen = r.readArrayHeader();
				this.valueBool = new Array(_valueBoolLen);
				for(let arrayIndex = 0; arrayIndex < _valueBoolLen; arrayIndex++) {
					this.valueBool[arrayIndex] = r.readBool();
				}
				break;
			case 14:
				if(r.isNull()) {
					r.readNil();
					this.valueString = null;
					continue;
				}
				const _valueStringLen = r.readArrayHeader();
				this.valueString = new Array(_valueStringLen);
				for(let arrayIndex = 0; arrayIndex < _valueStringLen; arrayIndex++) {
					this.valueString[arrayIndex] = r.readString();
				}
				break;
			case 15:
				if(r.isNull()) {
					r.readNil();
					this.valueBytes = null;
					continue;
				}
				const _valueBytesLen = r.readArrayHeader();
				this.valueBytes = new Array(_valueBytesLen);
				for(let arrayIndex = 0; arrayIndex < _valueBytesLen; arrayIndex++) {
					if(r.isNull()) {
						r.readNil();
						this.valueBytes[arrayIndex] = null;
					} else {
						this.valueBytes[arrayIndex] = r.readBytes();
					}
				}
				break;
			case 18:
				if(r.isNull()) {
					r.readNil();
					this.valueNestMessage = null;
					continue;
				}
				const _valueNestMessageLen = r.readArrayHeader();
				this.valueNestMessage = new Array(_valueNestMessageLen);
				for(let arrayIndex = 0; arrayIndex < _valueNestMessageLen; arrayIndex++) {
					if(r.isNull()) {
						r.readNil();
						this.valueNestMessage[arrayIndex] = null;
					} else {
						this.valueNestMessage[arrayIndex] = new _proto.DependTest(false);
						this.valueNestMessage[arrayIndex].read(r);
					}
				}
				break;
			case 20:
				if(r.isNull()) {
					r.readNil();
					this.valueTestEnum = null;
					continue;
				}
				const _valueTestEnumLen = r.readArrayHeader();
				this.valueTestEnum = new Array(_valueTestEnumLen);
				for(let arrayIndex = 0; arrayIndex < _valueTestEnumLen; arrayIndex++) {
					this.valueTestEnum[arrayIndex] = r.readNumber();
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = AllRepeatedParameter
