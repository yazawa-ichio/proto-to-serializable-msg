//proto/Forum.proto
"use strict";
var packer = require('proto-msgpack');
//add proto
if (!packer.proto) {
	packer.proto = {};
}
if (!packer.proto.Forum) {
	packer.proto.Forum = {};
}
module.exports = packer.proto.Forum.Roll = {
	GUEST: 0,
	MASTER: 1
};
