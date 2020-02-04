// sample/proto/Forum.proto
"use strict";
const _packer = require('proto-msgpack');
const _proto = require('./../index.js');
class ForumInfo {
	constructor(init, pos) {
		this.users = null;
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
		w.writeMapHeader(1);
		
		// Write users
		w.writeTag(1);
		if (this.users == null) {
			w.writeNil();
		} else {
			const mapLen = this.users.size;
			w.writeMapHeader(mapLen);
			this.users.forEach(function(value, key){
				w.writeNumber(key);
				if (!value) {
					w.writeNil();
				} else {
					value.write(w);
				}
			});
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
					this.users = null;
					continue;
				}
				const _usersLen = r.readMapHeader();
				this.users = new Map();
				for(let mapIndex = 0; mapIndex < _usersLen; mapIndex++) {
					let key;
					let value;
					key = r.readNumber();
					if(r.isNull()) {
						r.readNil();
						value = null;
					} else {
						value = new _proto.Forum.User(false);
						value.read(r);
					}
					this.users.set(key, value);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
module.exports = ForumInfo
ForumInfo.ForumNestInfo = class ForumNestInfo {
	constructor(init, pos) {
		this.userPost = null;
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
		w.writeMapHeader(1);
		
		// Write user_post
		w.writeTag(1);
		if (this.userPost == null) {
			w.writeNil();
		} else {
			const mapLen = this.userPost.size;
			w.writeMapHeader(mapLen);
			this.userPost.forEach(function(value, key){
				w.writeNumber(key);
				if (!value) {
					w.writeNil();
				} else {
					value.write(w);
				}
			});
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
					this.userPost = null;
					continue;
				}
				const _userPostLen = r.readMapHeader();
				this.userPost = new Map();
				for(let mapIndex = 0; mapIndex < _userPostLen; mapIndex++) {
					let key;
					let value;
					key = r.readNumber();
					if(r.isNull()) {
						r.readNil();
						value = null;
					} else {
						value = new _proto.Forum.PostData(false);
						value.read(r);
					}
					this.userPost.set(key, value);
				}
				break;
			default:
				r.skip();
				break;
			}
		}
	}
}
