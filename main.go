package main

import (
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"

	"log"
)

func main() {
	// all the heavy lifting done for you!
	if err := protokit.RunPlugin(new(plugin)); err != nil {
		log.Fatal(err)
	}
}

// plugin is an implementation of protokit.Plugin
type plugin struct{}

func (p *plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	descriptors := protokit.ParseCodeGenRequest(r)

	resp := new(plugin_go.CodeGeneratorResponse)

	data := newProtoData(descriptors)

	//TODO:パラメーター
	gen := newCSGenerator()

	for _, msg := range data.messages {
		if msg.parent == nil {
			resp.File = append(resp.File, gen.genClass(msg))
		}
	}
	for _, enum := range data.enums {
		if enum.parent == nil {
			resp.File = append(resp.File, gen.genEnum(enum))
		}
	}

	return resp, nil
}
