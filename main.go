package main

import (
	"errors"

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

	gen, err := getGenerator(r.GetParameter(), data)
	if err != nil {
		return nil, err
	}

	resp.File = append(resp.File, gen.genResponseFile(data)...)

	return resp, nil
}

func getGenerator(prm string, data *protoData) (generator, error) {
	switch prm {
	case "js":
		return newJSGenerator(data), nil
	case "cs":
		return newCSGenerator(data), nil
	case "ts":
		return newTSGenerator(data), nil
	}
	return nil, errors.New("not found generator")
}
