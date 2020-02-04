/// <reference types="node" />
import * as packer from 'proto-msgpack'
import * as proto from '../index'
import * as MyPackage from '../mypackage'

export class DependTestMessage{
	message: proto.PackageMessage | null;
	depDep: proto.DependMessage | null;
	constructor(init?: boolean | Buffer, pos?: number) 
}
export namespace DependTestMessage {
}
