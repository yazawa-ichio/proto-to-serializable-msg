package main

import plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"

type generator interface {
	genResponseFile(data *protoData) []*plugin_go.CodeGeneratorResponse_File
}
