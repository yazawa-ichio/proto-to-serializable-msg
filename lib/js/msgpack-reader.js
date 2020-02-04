"use strict";

module.exports = class Reader {

	constructor(buf, pos) {
		this.buf = buf;
		this.pos = pos | 0;
		this.format = 0;
	}

	reset(buf, pos) {
		this.buf = buf;
		this.pos = pos | 0;
		this.format = 0;
	}

	nextFormat() {
		return this.buf[this.pos];
	}

	readFormat() {
		return this.format = this.buf[this.pos++];
	}

	readBool() {
		return this.format == 0xc3;
	}

	readPositiveFixInt() {
		return this.format & 0x7f;
	}

	readUInt8() {
		return this.buf[this.pos++];
	}

	readUInt16() {
		this.pos += 2;
		return this.buf.readUInt16BE(this.pos - 2);
	}

	readUInt32() {
		this.pos += 4;
		return this.buf.readUInt32BE(this.pos - 4);
	}

	readUInt64() {
		this.pos += 8;
		return this.buf.readUInt32BE(this.pos - 8) * 4294967296 + this.buf.readUInt32BE(this.pos - 4);
	}

	readNegativeFixInt() {
		return (this.format & 0x1f) - 0x20;
	}

	readInt16() {
		this.pos += 2;
		return this.buf.readInt16BE(this.pos - 2);
	}

	readInt32() {
		this.pos += 4;
		return this.buf.readInt32BE(this.pos - 4);
	}

	readInt64() {
		this.pos += 8;
		return this.buf.readInt32BE(this.pos - 8) * 4294967296 + this.buf.readUInt32BE(this.pos - 4);
	}

	readFloat() {
		this.pos += 4;
		return this.buf.readFloatBE(this.pos - 4);
	}

	readDouble() {
		this.pos += 8;
		return this.buf.readDoubleBE(this.pos - 8);
	}

	readFixStr() {
		const length = this.format & 0x1f;
		this.pos += length;
		return this.buf.toString('utf-8', this.pos - length, this.pos);
	}

	readStr8() {
		const length = this.readUInt8();
		this.pos += length;
		return this.buf.toString('utf-8', this.pos - length, this.pos);
	}

	readStr16() {
		const length = this.readUInt16();
		this.pos += length;
		return this.buf.toString('utf-8', this.pos - length, this.pos);
	}

	readStr32() {
		const length = this.readUInt32();
		this.pos += length;
		return this.buf.toString('utf-8', this.pos - length, this.pos);
	}

	readBin8() {
		const length = this.readUInt8();
		this.pos += length;
		return this.buf.slice(this.pos - length, this.pos);
	}

	readBin16() {
		const length = this.readUInt16();
		this.pos += length;
		return this.buf.slice(this.pos - length, this.pos);
	}

	readBin32() {
		const length = this.readUInt32();
		this.pos += length;
		return this.buf.slice(this.pos - length, this.pos);
	}

	readArrayLength() {
		if (this.format == 0xc0) {
			return 0;
		} else if (this.format == 0xdc) {
			return this.readUInt16();
		} else if (this.format == 0xdd) {
			return this.readUInt32();
		} if ((this.format & 0xf0) === 0x90) {
			return this.format & 0xf;
		}
		throw new Error();
	}

	readMapLength() {
		if (this.format == 0xc0) {
			return 0;
		} else if (this.format == 0xde) {
			return this.readUInt16();
		} else if (this.format == 0xdf) {
			return this.readUInt32();
		} if ((this.format & 0xf0) === 0x80) {
			return this.format & 0xf;
		}
		throw new Error();
	}

	readExtLength() {
		switch (this.format) {
			case 0xd4: return 1;
			case 0xd5: return 2;
			case 0xd6: return 4;
			case 0xd7: return 8;
			case 0xd8: return 16;
			case 0xc7: return this.readUInt8();
			case 0xc8: return this.readUInt16();
			case 0xc9: return this.readUInt32();
		}
		throw new Error();
	}

	readExtType() {
		if (this.format == 0xcc) {
			return this.readUInt8();
		} else if (this.format == 0xd0) {
			return this.readUInt8();
		} else if ((this.format & 0x80) === 0x00) {
			return this.readPositiveFixInt();
		} else if ((this.format & 0xe0) === 0xe0) {
			return this.readNegativeFixInt();
		}
		throw new Error();
	}

	readExt(length) {
		this.pos += length;
		return this.buf.slice(this.pos - length, this.pos);
	}

	skip() {
		switch (this.readFormat()) {
			case 0xc0:
			case 0xc2:
			case 0xc3:
				return;
			case 0xcc:
			case 0xd0:
				this.pos += 1;
				return;
			case 0xcd:
			case 0xd1:
				this.pos += 2;
				return;
			case 0xce:
			case 0xd2:
			case 0xca:
				this.pos += 4;
				return;
			case 0xcf:
			case 0xd3:
			case 0xcb:
				this.pos += 8;
				return;
			case 0xd9:
			case 0xc4:
				this.pos += this.readUInt8();
				return;
			case 0xda:
			case 0xc5:
				this.pos += this.readUInt16();
				return;
			case 0xdb:
			case 0xc6:
				this.pos += this.readUInt32();
				return;
			case 0xd4:
				this.pos += 2;
				return;
			case 0xd5:
				this.pos += 3;
				return;
			case 0xd6:
				this.pos += 5;
				return;
			case 0xd7:
				this.pos += 9;
				return;
			case 0xd8:
				this.pos += 17;
				return;
			case 0xc7:
				this.pos += 1 + this.readUInt8();
				return;
			case 0xc8:
				this.pos += 1 + this.readUInt16();
				return;
			case 0xc9:
				this.pos += 1 + this.readUInt32();
				return;
		}
		if ((this.format & 0x80) === 0x00 || (this.format & 0xe0) === 0xe0) {
			return;
		}
		else if ((this.format & 0xe0) === 0xa0) {
			this.pos += this.format & 0x1f;
		}
		else if ((this.format & 0xf0) === 0x90) {
			for (let length = this.readArrayLength(); length > 0; length--) {
				this.skip();
			}
		}
		else if ((this.format & 0xf0) === 0x80) {
			for (let length = this.readMapLength(); length > 0; length--) {
				this.skip();
				this.skip();
			}
		}
	}

}
