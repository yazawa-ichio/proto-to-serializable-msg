package main

import plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"

type generator interface {
	genResponse() *plugin_go.CodeGeneratorResponse_File
}
