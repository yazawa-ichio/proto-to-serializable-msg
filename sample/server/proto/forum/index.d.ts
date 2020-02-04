/// <reference types="node" />
import * as packer from 'proto-msgpack'
import * as proto from '../index'

export class PostData{
	id: number;
	message: string;
	user: User | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace PostData {
}
export class User{
	id: number;
	name: string;
	roll: Roll;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace User {
}
export class PostForumReq{
	/**
	 * Data
	 */
	data: PostData | null;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace PostForumReq {
}
export class ForumData{
	data: Array<PostData | null>;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace ForumData {
}
export class ForumInfo{
	users: Map<number, User | null>;
	constructor(init?: boolean | Buffer, pos?: number) 
	pack(): Buffer;
	unpack(buf: Buffer, pos?: number): void;
	write(w: packer.ProtoWriter): void;
	read(r: packer.ProtoReader): void;
}
export namespace ForumInfo {
	class ForumNestInfo{
		userPost: Map<number, PostData | null>;
		constructor(init?: boolean | Buffer, pos?: number) 
		pack(): Buffer;
		unpack(buf: Buffer, pos?: number): void;
		write(w: packer.ProtoWriter): void;
		read(r: packer.ProtoReader): void;
	}
	namespace ForumNestInfo {
	}
}
/**
 * ロールです
 */
export enum Roll{
	/**
	 * ゲスト
	 */
	GUEST = 0,
	/**
	 * マスター
	 */
	MASTER = 1
}
