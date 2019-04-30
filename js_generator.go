package main

import (
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	ce "github.com/yazawa-ichio/protoc-gen-msgpack/code_emitter"
)

type jsGenerator struct {
	e    *ce.CodeEmitter
	conv *JSConverter
}

func newJSGenerator() *jsGenerator {
	return &jsGenerator{e: new(ce.CodeEmitter), conv: &JSConverter{}}
}

func (g *jsGenerator) genResponseFile(data *protoData) []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}
	for _, msg := range data.messages {
		files = append(files, g.genClass(msg))
	}
	for _, enum := range data.enums {
		files = append(files, g.genEnum(enum))
	}
	return files
}

func (g *jsGenerator) genClass(message *messageData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, message.file)
	g.e.EmitLine("\"use strict\";")
	g.e.EmitLine("var packer = require('ilib_proto_pack');")
	g.emitClass(message)
	fileName := g.conv.GetFileName(message.data.GetFullName())
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func isUserDefine(f *protokit.FieldDescriptor) bool {
	return f.GetType().String() == "TYPE_MESSAGE" ||
		f.GetType().String() == "TYPE_ENUM"
}

func (g *jsGenerator) getRequireName(typeName string) string {
	name := g.conv.GetFileName(typeName)
	return strings.Replace(strings.Replace(name, ".", "_", -1), "-", "$", -1)
}

func (g *jsGenerator) emitDeps(m *protokit.Descriptor) {
	hits := make(map[string]string)
	for _, f := range m.GetMessageFields() {
		if !isUserDefine(f) {
			continue
		}
		name := f.GetTypeName()
		if _, hit := hits[f.GetTypeName()]; hit {
			continue
		}
		hits[name] = name
		g.e.EmitLine("var %s = require('./%s');", g.getRequireName(name), g.conv.GetFileName(name))
	}
}

func (g *jsGenerator) emitClass(message *messageData) {
	g.e.Bracket("class %s ", func() {
		g.e.Bracket("constructor() ", func() {
			for _, f := range message.data.GetMessageFields() {
				g.e.EmitLine("this.%s = %s;", f.GetName(), f.GetName())
			}
		})
		g.e.Bracket("write(w) ", func() {
			g.emitWriter(message)
		})
		g.e.Bracket("reader(r) ", func() {
			g.emitReader(message)
		})
	}, message.data.GetName())
	g.emitExports(message.data.GetName(), message.data.GetPackage(), message.parent)
}

func (g *jsGenerator) genEnum(enum *enumData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, enum.file)
	g.e.EmitLine("\"use strict\";")
	g.e.EmitLine("var packer = require('ilib_proto_pack');")
	g.emitEnum(enum)
	fileName := g.conv.GetFileName(enum.data.GetFullName())
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *jsGenerator) emitEnum(enum *enumData) {
	emitDoc(g.e, enum.data.GetComments().GetLeading())
	g.e.StartBracket("const %s = ", enum.data.GetName())
	vals := enum.data.GetValues()
	for i, v := range vals {
		g.emitComment(v.GetComments().GetLeading())
		if i < len(vals)-1 {
			g.e.EmitLine("%s: %d,", g.conv.GetEnumName(v.GetName()), v.GetNumber())
		} else {
			g.e.EmitLine("%s: %d", g.conv.GetEnumName(v.GetName()), v.GetNumber())
		}
	}
	g.e.EndBracket(";")
	g.emitExports(enum.data.GetName(), enum.data.GetPackage(), enum.parent)
}

func (g *jsGenerator) emitComment(comment string) bool {
	if comment == "" {
		return false
	}
	g.e.EmitLine("/*")
	for _, c := range strings.Split(comment, "\n") {
		g.e.EmitLine(c)
	}
	g.e.EmitLine("*/")
	return true
}

func emitDoc(e *ce.CodeEmitter, comment string) bool {
	if comment == "" {
		return false
	}
	e.EmitLine("/**")
	for _, c := range strings.Split(comment, "\n") {
		e.EmitLine(" *" + c)
	}
	e.EmitLine(" */")
	return true
}

func (g *jsGenerator) emitExports(name string, packageName string, parent *protokit.Descriptor) {

	g.e.EmitLine("//add exports")
	g.e.EmitLine("module.exports = %s;", name)

	g.e.EmitLine("//add proto")
	if parent != nil {
		parentName := parent.GetFullName()
		g.e.EmitLine("var parent = require('./%s');", g.conv.GetFileName(parentName))
		g.e.EmitLine("parent.%s = %s;", name, name)
	} else {
		g.e.StartBracket("if (!packer.proto) ")
		g.e.EmitLine("packer.proto = {};")
		g.e.EndBracket("")
		if packageName != "" {
			g.e.StartBracket("if (!packer.proto.%s) ", packageName)
			g.e.EmitLine("packer.proto.%s = {};", packageName)
			g.e.EndBracket("")
			g.e.EmitLine("packer.proto.%s.%s = %s;", packageName, name, name)
		} else {
			g.e.EmitLine("packer.proto.%s = %s;", name, name)
		}
	}
}

func (g *jsGenerator) emitWriter(message *messageData) {
	g.e.EmitLine("// Write Map Length")
	g.e.EmitLine("w.writeMapHeader(%s);", strconv.Itoa(len(message.data.GetMessageFields())))
	for _, f := range message.data.GetMessageFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := g.conv.GetFieldName(f.GetName())
		g.e.EmitLine("w.writeTag(%s);", strconv.Itoa(int(f.GetNumber())))
		if f.GetLabel().String() == "LABEL_REPEATED" {
			g.e.Bracket("if (!this.%s) ", func() {
				g.e.EmitLine("w.writeNil();")
				g.e.EndAndStartBracket(" else ")
				g.e.EmitLine("const arrayLen = %s.Length;", filedName)
				g.e.EmitLine("w.writeArrayHeader(arrayLen);")
				g.e.EmitLine("for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)")
				g.e.Bracket("", func() {
					g.emitSerialize(f, "[arrayIndex]")
				})
			}, filedName)
		} else {
			g.emitSerialize(f, "")
		}
	}
}

func (g *jsGenerator) emitSerialize(f *protokit.FieldDescriptor, suffix string) {
	filedName := g.conv.GetFieldName(f.GetName()) + suffix
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		g.e.Bracket("if (!%s) ", func() {
			g.e.EmitLine("w.writeNil();")
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("this.%s.write(w);", filedName)
		}, filedName)
		break
	case "TYPE_BYTES":
		g.e.Bracket("if (!this.%s) ", func() {
			g.e.EmitLine("w.writeNil();")
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("w.writeBytes(this.%s);", filedName)
		}, filedName)
		break
	case "TYPE_BOOL":
		g.e.EmitLine("w.writeBool(this.%s);", filedName)
		break
	case "TYPE_STRING":
		g.e.EmitLine("w.writeString(this.%s);", filedName)
		break
	default:
		g.e.EmitLine("w.writeNumber(this.%s);", filedName)
		break
	}
}

func (g *jsGenerator) emitReader(message *messageData) {
	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("const mapLen = r.readMapHeader();")
	g.e.NewLine()
	g.e.Bracket("for(let i = 0; i < mapLen; i++) ", func() {
		g.e.EmitLine("const tag = r.readTag();")
		g.e.EmitLine("switch(tag) {")
		for _, f := range message.data.GetMessageFields() {
			g.e.Bracket("case %d: ", func() {
				if f.GetLabel().String() == "LABEL_REPEATED" {
					g.emitRepeatedDeserialize(f)
				} else {
					g.emitDeserialize(f, "")
				}
				g.e.EmitLine("break;")
			}, f.GetNumber())
		}
		g.e.EmitLine("default:")
		g.e.StartIndent()
		g.e.EmitLine("r.readSkip();")
		g.e.EmitLine("break;")
		g.e.EndIndent()
		g.e.EmitLine("}")
	})

}

func (g *jsGenerator) emitRepeatedDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	g.e.Bracket("if(r.isNull()) ", func() {
		g.e.EmitLine("%s = r.readNil();", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("let arrayLen = r.ReadArrayHeader();")
	g.e.EmitLine("this.%s = new Array(arrayLen);", filedName)
	g.e.Bracket("for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++) ", func() {
		g.emitDeserialize(f, "[arrayIndex]")
	})
}

func (g *jsGenerator) emitDeserialize(f *protokit.FieldDescriptor, suffix string) {
	filedName := g.conv.GetFieldName(f.GetName())
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("this.%s = r.readNil();", filedName+suffix)
			g.e.EmitLine("continue;")
		})
		g.e.StartBracket("if(!this.%s) ", filedName+suffix)
		g.e.EmitLine("this.%s = new packer.proto.%s();", filedName+suffix, msgType)
		g.e.EndBracket("")
		g.e.EmitLine("this.%s.read(r);", filedName+suffix)
		break
	case "TYPE_ENUM":
		g.e.EmitLine("this.%s = r.readInt();", filedName+suffix)
		break
	case "TYPE_BYTES":
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("this.%s = r.readNil();", filedName+suffix)
			g.e.EmitLine("continue;")
		})
		g.e.EmitLine("this.%s = r.readBytes();", filedName+suffix)
		break
	default:
		g.e.EmitLine("this.%s = r.read%s();", g.conv.GetFieldName(f.GetName())+suffix, strings.Title(g.conv.GetTypeImpl(f)))
		break
	}
}
