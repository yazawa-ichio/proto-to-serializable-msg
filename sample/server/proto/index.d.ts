/// <reference types="node" />
import * as packer from 'proto-msgpack'
import * as Forum from './forum'

export class DependTest{
	dep: Forum.PostData | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace DependTest {
}
