//proto/Forum.proto
/// <reference types="node" />
import * as packer from 'proto-msgpack'
import forum_roll = require('./forum.roll');
declare class User{
	Id: number;
	Name: string;
	Roll: forum_roll;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export = User;
