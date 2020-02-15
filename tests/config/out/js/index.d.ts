/// <reference types="node" />
import * as packer from 'proto-msgpack'
export class AllParameter{
	valueDouble: number;
	valueFloat: number;
	valueInt32: number;
	valueInt64: number;
	valueUint32: number;
	valueUint64: number;
	valueSint32: number;
	valueSint64: number;
	valueFixed32: number;
	valueFixed64: number;
	valueSfixed32: number;
	valueSfixed64: number;
	valueBool: boolean;
	valueString: string;
	valueBytes: Uint8Array | null;
	valueMapString: Map<number, string>;
	valueMapInt: Map<string, number>;
	valueMessage: EmptyMessage | null;
	valueMapValueMessage: Map<number, DependTest | null>;
	valueTestEnum: TestEnum;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace AllParameter {
}
export class AllRepeatedParameter{
	valueDouble: Array<number>;
	valueFloat: Array<number>;
	valueInt32: Array<number>;
	valueInt64: Array<number>;
	valueUint32: Array<number>;
	valueUint64: Array<number>;
	valueSint32: Array<number>;
	valueSint64: Array<number>;
	valueFixed32: Array<number>;
	valueFixed64: Array<number>;
	valueSfixed32: Array<number>;
	valueSfixed64: Array<number>;
	valueBool: Array<boolean>;
	valueString: Array<string>;
	valueBytes: Array<Uint8Array | null>;
	valueNestMessage: Array<DependTest | null>;
	valueTestEnum: Array<TestEnum>;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace AllRepeatedParameter {
}
export class EmptyMessage{
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace EmptyMessage {
}
/**
 *  UpperCamelCase comment
 */
export class UpperCamelCase{
	upperCamelCaseField: number;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace UpperCamelCase {
}
/**
 *  lowerCamelCase comment
 */
export class LowerCamelCase{
	lowerCamelCaseField: number;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace LowerCamelCase {
}
export class SnakeCase{
	snakeCaseValue: number;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace SnakeCase {
	class NestSnakeCase{
		nestSnakeCaseValue: number;
		constructor(init?: boolean | Buffer, pos?: number) 
		pack(): Buffer;
		unpack(buf: Buffer, pos?: number): void;
		write(w: packer.ProtoWriter): void;
		read(r: packer.ProtoReader): void;
	}
	namespace NestSnakeCase {
	}
}
export class DependTest{
	msg: DependMessage | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace DependTest {
}
/**
 *  comment
 */
export enum TestEnum{
	/**
	 *  TEST_NONE 0 comment
	 */
	TEST_NONE = 0,
	/**
	 *  TestValue 0 comment
	 *  2Line
	 */
	TESTVALUE = 1,
	SNAKE_TEST_VALUE = 2
}
export class DependMessage{
	text: string;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace DependMessage {
}
export class PackageMessage{
	message: MyPackage.MyMessage | null;
	myEnum: MyPackage.MyEnum;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace PackageMessage {
}

export namespace DepDep {
	class DependTestMessage{
		message: PackageMessage | null;
		depDep: DependMessage | null;
		constructor(init?: boolean | Buffer, pos?: number) 
		pack(): Buffer;
		unpack(buf: Buffer, pos?: number): void;
		write(w: packer.ProtoWriter): void;
		read(r: packer.ProtoReader): void;
	}
	namespace DependTestMessage {
	}
}

export namespace MyPackage {
	class MyMessage{
		constructor(init?: boolean | Buffer, pos?: number) 
		pack(): Buffer;
		unpack(buf: Buffer, pos?: number): void;
		write(w: packer.ProtoWriter): void;
		read(r: packer.ProtoReader): void;
	}
	namespace MyMessage {
	}
	enum MyEnum{
		NONE = 0,
		VALUT = 2
	}
}
