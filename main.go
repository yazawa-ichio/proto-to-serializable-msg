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
	gen := newJSGenerator()
	resp.File = append(resp.File, gen.genResponseFile(data)...)
	resp.File = append(resp.File, newCSGenerator().genResponseFile(data)...)

	return resp, nil
}
