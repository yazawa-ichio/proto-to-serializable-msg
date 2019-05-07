/// <reference types="node" />
declare class Reader {
	constructor(buf?: Buffer, pos?: number | undefined);
	reset(buf: Buffer, pos?: number | undefined): void;
	nextFormat(): number;
	readFormat(): number;
	readBool(): boolean;
	readPositiveFixInt(): number;
	readUInt8(): number;
	readUInt16(): number;
	readUInt32(): number;
	readUInt64(): number;
	readNegativeFixInt(): number;
	readInt16(): number;
	readInt32(): number;
	readInt64(): number;
	readFloat(): number;
	readDouble(): number;
	readFixStr(): string | null;
	readStr8(): string | null;
	readStr16(): string | null;
	readStr32(): string | null;
	readBin8(): Buffer | null;
	readBin16(): Buffer | null;
	readBin32(): Buffer | null;
	readArrayLength(): number;
	readMapLength(): number;
	readExtLength(): number;
	readExtType(): number;
	readExt(length: number): Buffer | null;
	skip(): void;
}
export = Reader;