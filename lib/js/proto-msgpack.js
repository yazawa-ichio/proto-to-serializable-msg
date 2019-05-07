"use strict";

const Reader = require('./msgpack-reader');
const Writer = require('./msgpack-writer');

exports.ProtoWriter = class ProtoWriter extends Writer {
	writeTag(val) {
		super.writeNumber(val)
	}
}

exports.defaultWriter = new exports.ProtoWriter();

exports.ProtoReader = class ProtoReader {

	constructor(buf, pos){
		this._r = new Reader(buf, pos);
	}

	isNull() {
		return this._r.nextFormat() == 0xc0;
	}

	readNil(){
		if (this._r.readFormat() != 0xc0){
			throw new Error();
		}
		return null;
	}

	readBool(){
		return this._r.readFormat() == 0xc3;
	}

	readTag() {
		return this.readNumber();
	}

	readDouble() {
		return this.readNumber();
	}

	readFloat() {
		return this.readNumber();
	}

	readNumber(){
		switch (this._r.readFormat()){
			case 0xca: return this._r.readFloat();
			case 0xcb: return this._r.readDouble();
			case 0xcc: return this._r.readUInt8();
			case 0xcd: return this._r.readUInt16();
			case 0xce: return this._r.readUInt32();
			case 0xcf: return this._r.readUInt64();
			case 0xd0: return this._r.readInt8();
			case 0xd1: return this._r.readInt16();
			case 0xd2: return this._r.readInt32();
			case 0xd3: return this._r.readInt64();
			default:
				if ((this._r.format & 0x80) === 0x00) return this._r.readPositiveFixInt();
				if ((this._r.format & 0xe0) === 0xe0) return this._r.readNegativeFixInt();
		}
		throw new Error();
	}

	readBytes() {
		switch (this._r.readFormat()) {
			case 0xc4: return this._r.readBin8();
			case 0xc5: return this._r.readBin16();
			case 0xc6: return this._r.readBin32();
		}
		throw new Error();
	}

	readString() {
		switch (this._r.readFormat()) {
			case 0xc0: return "";
			case 0xd9: return this._r.readStr8();
			case 0xda: return this._r.readBin16();
			case 0xdb: return this._r.readBin32();
			default:
				if (this._r.format >= 0xa0 && this._r.format <= 0xbf) {
					return this._r.readFixStr();
				}
				break;
		}
		throw new Error();
	}

	readArrayHeader(){
		this._r.readFormat();
		return this._r.readArrayLength();
	}

	readMapHeader(){
		this._r.readFormat();
		return this._r.readMapLength();
	}
	readSkip(){
		this._r.skip();
	}
}

