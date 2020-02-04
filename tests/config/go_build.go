package main

import (
	"log"

	proto_test "github.com/yazawa-ichio/protoc-gen-msgpack/tests/config/out/go"
	proto_root_test "github.com/yazawa-ichio/protoc-gen-msgpack/tests/config/out/go_root"
	proto_root_mypackage_test "github.com/yazawa-ichio/protoc-gen-msgpack/tests/config/out/go_root/mypackage"
)

func main() {
	// シリアライズのテストは別途行っているので
	// 異なるパラメータ出力でコンパイルエラーがないことが確認出来ればいい
	log.Print(proto_test.AllParameter{})
	log.Print(proto_test.MyPackage_MyMessage{})
	log.Print(proto_root_test.PackageMessage{})
	log.Print(proto_root_mypackage_test.MyMessage{})
}
