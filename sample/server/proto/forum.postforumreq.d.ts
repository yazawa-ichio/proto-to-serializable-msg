//proto/Forum.proto
/// <reference types="node" />
import * as packer from 'proto-msgpack'
import forum_postdata = require('./forum.postdata');
declare class PostForumReq{
	data: forum_postdata | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export = PostForumReq;
