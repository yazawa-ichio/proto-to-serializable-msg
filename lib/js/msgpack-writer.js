"use strict";

module.exports = class Writer {

	constructor() {
		this.buf = Buffer.alloc(256);
		this.pos = 0;
	}

	toBuffer() {
		return this.buf.slice(0, this.pos);
	}

	clear() {
		this.pos = 0;
	}

	expendSizeIfNeed(size) {
		if (this.buf.length < this.pos + size) {
			this.buf = Buffer.concat([this.buf, Buffer.alloc(256)])
		}
	}

	writeNil() {
		this.expendSizeIfNeed(1);
		this.buf[this.pos++] = 0xc0;
	}

	writeBool(val) {
		this.expendSizeIfNeed(1);
		this.buf[this.pos++] = val ? 0xc3 : 0xc2;
	}

	writeFloat(val) {
		this.expendSizeIfNeed(5);
		this.buf[this.pos++] = 0xca;
		this.buf.writeFloatBE(val, this.pos);
		this.pos += 4;
	}

	writeDouble(val) {
		this.expendSizeIfNeed(9);
		this.buf[this.pos++] = 0xcb;
		this.buf.writeDoubleBE(val, this.pos);
		this.pos += 8;
	}

	writeNumber(val) {
		if (val >= 0) {
			if (val < 0x80) {
				// positive fixnum
				this.expendSizeIfNeed(1);
				this.buf[this.pos++] = val;
			} else if (val < 0x100) {
				// uint 8
				this.expendSizeIfNeed(2);
				this.buf[this.pos++] = 0xcc;
				this.buf[this.pos++] = val;
			} else if (val < 0x10000) {
				// uint 16
				this.expendSizeIfNeed(3);
				this.buf[this.pos++] = 0xcd;
				this.buf.writeUInt16BE(val, this.pos);
				this.pos += 2;
			} else if (val < 0x100000000) {
				// uint 32
				this.expendSizeIfNeed(5);
				this.buf[this.pos++] = 0xce;
				this.buf.writeUInt32BE(val, this.pos);
				this.pos += 4;
			} else {
				// uint 64
				this.expendSizeIfNeed(9);
				this.buf[this.pos++] = 0xcf;
				let tmpVal = val;
				for (let i = 7; i >= 0; i--) {
					this.buf[this.pos + i] = tmpVal & 255;
					tmpVal /= 256;
				}
				this.pos += 8;
			}
		} else {
			if (val >= -0x20) {
				// negative fixnum
				this.expendSizeIfNeed(1);
				this.buf[this.pos++] = val | 0x20;
			} else if (val >= -0x80) {
				// int 8
				this.expendSizeIfNeed(2);
				this.buf[this.pos++] = 0xd0;
				this.buf[this.pos++] = val;
			} else if (val >= -0x8000) {
				// int 16
				this.expendSizeIfNeed(3);
				this.buf[this.pos++] = 0xd1;
				this.buf.writeInt16BE(val, this.pos);
				this.pos += 2;
			} else if (val >= -0x80000000) {
				// int 32
				this.expendSizeIfNeed(5);
				this.buf[this.pos++] = 0xd2;
				this.buf.writeInt32BE(val, this.pos);
				this.pos += 4;
			} else {
				// int 64
				this.expendSizeIfNeed(9);
				this.buf[this.pos++] = 0xd3;
				let tmpVal = val + 1;
				for (let i = 7; i >= 0; i--) {
					this.buf[this.pos + i] = ((-tmpVal) & 255) ^ 255;
					tmpVal /= 256;
				}
				this.pos += 8;
			}
		}
	}

	writeString(val) {
		if (!val) {
			this.writeNil();
			return;
		}
		const srtBuf = Buffer.from(val, 'utf-8');
		const length = srtBuf.length;
		if (length < 0x20) {
			// fixstr
			this.expendSizeIfNeed(length + 1);
			this.buf[this.pos++] = length | 0xa0;
		} else if (length < 0x100) {
			// str 8
			this.expendSizeIfNeed(length + 2);
			this.buf[this.pos++] = 0xd9;
			this.buf.writeUInt8(length, this.pos);
			this.pos += 1;
		} else if (length < 0x10000) {
			// str 16
			this.expendSizeIfNeed(length + 3);
			this.buf[this.pos++] = 0xda;
			this.buf.writeUInt16BE(length, this.pos);
			this.pos += 2;
		} else {
			// str 32
			this.expendSizeIfNeed(length + 4);
			this.buf[this.pos++] = 0xdb;
			this.buf.writeUInt32BE(length, this.pos);
			this.pos += 3;
		}
		srtBuf.copy(this.buf, this.pos, 0, length);
		this.pos += length;
	}

	writeBytes(val) {
		if (!val) {
			writeNil();
			return;
		}
		const length = val.length;
		if (length < 0x100) {
			// bin 8
			this.expendSizeIfNeed(length + 2);
			this.buf[this.pos++] = 0xc4;
			this.buf.writeUInt8(length, this.pos);
			this.pos += 1;
		} else if (length < 0x10000) {
			// bin 16
			this.expendSizeIfNeed(length + 3);
			this.buf[this.pos++] = 0xc5;
			this.buf.writeUInt16BE(length, this.pos);
			this.pos += 2;
		} else {
			// bin 32
			this.expendSizeIfNeed(length + 3);
			this.buf[this.pos++] = 0xc6;
			this.buf.writeUInt32BE(length, this.pos);
			this.pos += 3;
		}
		val.copy(this.buf, this.pos, 0, length);
		this.pos += length;
	}

	writeArrayHeader(length) {
		if (length < 0x10) {
			// fixarray
			this.expendSizeIfNeed(1);
			this.buf[this.pos++] = length | 0x90;
		} else if (length < 0x10000) {
			// array 16
			this.expendSizeIfNeed(3);
			this.buf[this.pos++] = 0xdc;
			this.buf.writeUInt16BE(length, this.pos);
			this.pos += 2;
		} else {
			// array 32
			this.expendSizeIfNeed(5);
			this.buf[this.pos++] = 0xdd;
			this.buf.writeUInt32BE(length, this.pos);
			this.pos += 4;
		}
	}

	writeMapHeader(length) {
		if (length < 0x10) {
			// fixmap
			this.expendSizeIfNeed(1);
			this.buf[this.pos++] = length | 0x80;
		} else if (length < 0x10000) {
			// map 16
			this.expendSizeIfNeed(3);
			this.buf[this.pos++] = 0xde;
			this.buf.writeUInt16BE(length, this.pos);
			this.pos += 2;
		} else {
			// map 32
			this.expendSizeIfNeed(5);
			this.buf[this.pos++] = 0xdf;
			this.buf.writeUInt32BE(length, this.pos);
			this.pos += 4;
		}
	}

	writeExt(extType, val) {
		const length = val.length;
		if (length == 1) {
			this.expendSizeIfNeed(2 + length)
			this.buf[this.pos++] = 0xd4;
		} else if (length == 2) {
			this.expendSizeIfNeed(2 + length)
			this.buf[this.pos++] = 0xd5;
		} else if (length == 4) {
			this.expendSizeIfNeed(2 + length)
			this.buf[this.pos++] = 0xd6;
		} else if (length == 8) {
			this.expendSizeIfNeed(2 + length)
			this.buf[this.pos++] = 0xd7;
		} else if (length == 16) {
			this.expendSizeIfNeed(2 + length)
			this.buf[this.pos++] = 0xd8;
		} else if (length < 0x100) {
			this.expendSizeIfNeed(3 + length)
			this.buf[this.pos++] = 0xc7;
			this.buf[this.pos++] = length;
		} else if (length < 0x10000) {
			this.expendSizeIfNeed(4 + length)
			this.buf[this.pos++] = 0xc8;
			this.buf.writeUInt16BE(length, this.pos);
			this.pos += 2;
		} else {
			this.expendSizeIfNeed(5 + length)
			this.buf[this.pos++] = 0xc9;
			this.buf.writeUInt32BE(length, this.pos);
			this.pos += 4;
		}
		this.buf[this.pos++] = extType;
		val.copy(this.buf, this.pos, 0, length);
		this.pos += length;
	}

}

