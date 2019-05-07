/// <reference types="node" />
declare class Writer {
	toBuffer(): Buffer;
	clear(): void;
	expendSizeIfNeed(size: number):void;
	writeNil() :void;
	writeBool(val: boolean): void;
	writeFloat(val: number): void;
	writeDouble(val: number): void;
	writeNumber(val: number): void;
	writeBytes(val: Buffer): void;
	writeArrayHeader(length: number): void;
	writeMapHeader(length: number): void;
	writeExt(extType: number, val: Buffer) : void;
}
export = Writer
