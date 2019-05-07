//proto/Forum.proto
/// <reference types="node" />
import * as packer from 'proto-msgpack'
import forum_user = require('./forum.user');
declare class PostData{
	Id: number;
	Message: string;
	User: forum_user | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export = PostData;
