/// <reference types="node" />
import Reader = require('./msgpack-reader')
import Writer = require('./msgpack-writer')

export class ProtoWriter extends Writer {
	writeTag(val: number): void;
}

export var defaultWriter: ProtoWriter;

export class ProtoReader {
	constructor(buf: Buffer, pos?: number | undefined);
	isNull(): boolean;
	readNil(): null;
	readBool(): boolean;
	readTag(): number;
	readDouble(): number;
	readFloat(): number;
	readNumber(): number;
	readBytes(): Buffer;
	readString(): string;
	readArrayHeader(): number;
	readMapHeader(): number;
	readSkip(): void;
}


