/// <reference types="node" />
import * as packer from 'proto-msgpack'
import * as proto from '../index'
import * as DepDep from '../depdep'

export class MyMessage{
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace MyMessage {
}
export enum MyEnum{
	NONE = 0,
	VALUT = 2
}
