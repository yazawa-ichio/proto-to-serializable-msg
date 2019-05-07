//proto/Forum.proto
/// <reference types="node" />
import * as packer from 'proto-msgpack'
import forum_postdata = require('./forum.postdata');
declare class ForumData{
	data: Array<forum_postdata | null> | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export = ForumData;
