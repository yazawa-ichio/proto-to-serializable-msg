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
	data *protoData
}

func newJSGenerator(data *protoData) *jsGenerator {
	return &jsGenerator{
		e:    new(ce.CodeEmitter),
		conv: newJSConverter(data),
		data: data,
	}
}

func (g *jsGenerator) genResponseFile(data *protoData) []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}
	for _, msg := range data.messages {
		if !msg.mapEntry {
			files = append(files, g.genClass(msg))
		}
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
	g.e.EmitLine("var packer = require('proto-msgpack');")
	g.emitDeps(message.data)
	g.emitClass(message)
	fileName := g.conv.GetFileName(message.data.GetFullName())
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *jsGenerator) getRequireName(typeName string) string {
	name := g.conv.GetFileName(typeName)
	return strings.Replace(strings.Replace(name, ".", "_", -1), "-", "$", -1)
}

func (g *jsGenerator) emitDeps(m *protokit.Descriptor) {
	hits := make(map[string]string)
	for _, f := range m.GetMessageFields() {
		if g.data.isMapEntry(f) {
			key, val := g.data.getMapKeyValue(f)
			g.emitRequire(hits, key)
			g.emitRequire(hits, val)
		} else {
			g.emitRequire(hits, f)
		}
	}
}

func (g *jsGenerator) emitRequire(hits map[string]string, f *protokit.FieldDescriptor) {
	if !g.data.isUserDefine(f) {
		return
	}
	name := f.GetTypeName()
	if _, hit := hits[f.GetTypeName()]; hit {
		return
	}
	hits[name] = name
	g.e.EmitLine("require('./%s');", g.conv.GetFileName(name))
}

func (g *jsGenerator) emitClass(message *messageData) {
	g.emitExports(message.data.GetName(), message.data.GetPackage(), message.parent, "class "+message.data.GetName()+" ")
	g.e.StartIndent()
	g.e.Bracket("constructor(init, pos) ", func() {
		for _, f := range message.data.GetMessageFields() {
			if f.GetLabel().String() == "LABEL_REPEATED" {
				g.e.EmitLine("this.%s = null;", f.GetName())
				continue
			}
			switch f.GetType().String() {
			case "TYPE_ENUM", "TYPE_FLOAT", "TYPE_DOUBLE":
				g.e.EmitLine("this.%s = 0;", f.GetName())
				break
			case "TYPE_MESSAGE":
				g.e.Bracket("if(init == null || init == true)", func() {
					g.e.EmitLine("this.%s = new packer.proto.%s();", f.GetName(), g.conv.GetType(f))
					g.e.EndAndStartBracket(" else ")
					g.e.EmitLine("this.%s = null;", f.GetName())
				})
				break
			default:
				if g.conv.GetTypeImpl(f) == "number" {
					g.e.EmitLine("this.%s = 0;", f.GetName())
				} else {
					g.e.EmitLine("this.%s = null;", f.GetName())
				}
				break
			}
		}
		g.e.Bracket("if(Buffer.isBuffer(init))", func() {
			g.e.EmitLine("this.read(new packer.ProtoReader(init, pos));")
		})
	})
	g.e.Bracket("pack() ", func() {
		g.e.EmitLine("const w = packer.defaultWriter;")
		g.e.EmitLine("w.clear();")
		g.e.EmitLine("this.write(w);")
		g.e.EmitLine("return w.toBuffer();")
	})
	g.e.Bracket("unpack(buf, pos) ", func() {
		g.e.Bracket("if(!Buffer.isBuffer(buf))", func() {
			g.e.EmitLine("this.read(buf);")
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("this.read(new packer.ProtoReader(buf, pos));")
		})
	})
	g.e.Bracket("write(w) ", func() {
		g.emitWriter(message)
	})
	g.e.Bracket("read(r) ", func() {
		g.emitReader(message)
	})
	g.e.EndBracket("")
}

func (g *jsGenerator) genEnum(enum *enumData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, enum.file)
	g.e.EmitLine("\"use strict\";")
	g.e.EmitLine("var packer = require('proto-msgpack');")
	g.emitEnum(enum)
	fileName := g.conv.GetFileName(enum.data.GetFullName())
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *jsGenerator) emitEnum(enum *enumData) {
	g.emitComment(enum.data.GetComments().GetLeading())
	g.emitExports(enum.data.GetName(), enum.data.GetPackage(), enum.parent, "")
	g.e.StartIndent()
	//	g.e.StartBracket("const %s = ", enum.data.GetName())
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

func (g *jsGenerator) emitExports(name string, packageName string, parent *protokit.Descriptor, suffix string) {

	g.e.EmitLine("//add proto")
	if parent != nil {
		parentName := parent.GetFullName()
		g.e.EmitLine("var parent = require('./%s');", g.conv.GetFileName(parentName))
		g.e.EmitLine("parent.%s = %s{", name, suffix)
	} else {
		g.e.StartBracket("if (!packer.proto) ")
		g.e.EmitLine("packer.proto = {};")
		g.e.EndBracket("")
		if packageName != "" {
			g.e.StartBracket("if (!packer.proto.%s) ", packageName)
			g.e.EmitLine("packer.proto.%s = {};", packageName)
			g.e.EndBracket("")
			g.e.EmitLine("module.exports = packer.proto.%s.%s = %s{", packageName, name, suffix)
		} else {
			g.e.EmitLine("module.exports = packer.proto.%s = %s{", name, suffix)
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

		if g.data.isMapEntry(f) {
			g.e.Bracket("if (this.%s == null) ", func() {
				g.e.EmitLine("w.writeNil();")
				g.e.EndAndStartBracket(" else ")
				g.e.EmitLine("const mapLen = this.%s.size;", filedName)
				g.e.EmitLine("w.WriteMapHeader(mapLen);")
				g.e.StartBracket("this.%s.forEach(function(value, key)", filedName)
				mapKey, mapVal := g.data.getMapKeyValue(f)
				g.emitSerialize(mapKey, "", false)
				g.emitSerialize(mapVal, "", false)
				g.e.EndBracket(");")
			}, filedName)
		} else if f.GetLabel().String() == "LABEL_REPEATED" {
			g.e.Bracket("if (this.%s == null) ", func() {
				g.e.EmitLine("w.writeNil();")
				g.e.EndAndStartBracket(" else ")
				g.e.EmitLine("const arrayLen = this.%s.length;", filedName)
				g.e.EmitLine("w.writeArrayHeader(arrayLen);")
				g.e.EmitLine("for(let arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)")
				g.e.Bracket("", func() {
					g.emitSerialize(f, "[arrayIndex]", true)
				})
			}, filedName)
		} else {
			g.emitSerialize(f, "", true)
		}
	}
}

func (g *jsGenerator) emitSerialize(f *protokit.FieldDescriptor, suffix string, field bool) {
	prefix := ""
	if field {
		prefix = "this."
	}
	filedName := prefix + g.conv.GetFieldName(f.GetName()) + suffix
	typeName := strings.Title(g.conv.GetTypeImpl(f))
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		g.e.Bracket("if (!%s) ", func() {
			g.e.EmitLine("w.writeNil();")
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("%s.write(w);", filedName)
		}, filedName)
		break
	case "TYPE_BYTES":
		g.e.Bracket("if (!%s) ", func() {
			g.e.EmitLine("w.writeNil();")
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("w.writeBytes(%s);", filedName)
		}, filedName)
		break
	case "TYPE_ENUM":
		g.e.EmitLine("w.writeNumber(%s);", filedName)
		break
	default:
		g.e.EmitLine("w.write%s(%s);", typeName, filedName)
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
			g.e.EmitLine("case %d:", f.GetNumber())
			g.e.StartIndent()
			if g.data.isMapEntry(f) {
				g.emitMapDeserialize(f)
			} else if f.GetLabel().String() == "LABEL_REPEATED" {
				g.emitRepeatedDeserialize(f)
			} else {
				g.emitDeserialize(f, "", true)
			}
			g.e.EmitLine("break;")
			g.e.EndIndent()
		}
		g.e.EmitLine("default:")
		g.e.StartIndent()
		g.e.EmitLine("r.skip();")
		g.e.EmitLine("break;")
		g.e.EndIndent()
		g.e.EmitLine("}")
	})

}

func (g *jsGenerator) emitMapDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	key, value := g.data.getMapKeyValue(f)
	g.e.Bracket("if(r.isNull()) ", func() {
		g.e.EmitLine("r.readNil();")
		g.e.EmitLine("this.%s = null;", filedName)
		g.e.EmitLine("continue;")
	})

	g.e.EmitLine("const _%sLen = r.readMapHeader();", filedName)
	g.e.EmitLine("this.%s = new Map();", filedName)
	g.e.Bracket("for(let mapIndex = 0; mapIndex < _%sLen; mapIndex++) ", func() {
		g.e.EmitLine("let key;")
		g.e.EmitLine("let value;")
		g.emitDeserialize(key, "", false)
		g.emitDeserialize(value, "", false)
		g.e.EmitLine("this.%s.set(key, value);", filedName)

	}, filedName)
}

func (g *jsGenerator) emitRepeatedDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	g.e.Bracket("if(r.isNull()) ", func() {
		g.e.EmitLine("r.readNil();")
		g.e.EmitLine("this.%s = null;", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("const _%sLen = r.readArrayHeader();", filedName)
	g.e.EmitLine("this.%s = new Array(_%sLen);", filedName, filedName)
	g.e.Bracket("for(let arrayIndex = 0; arrayIndex < _%sLen; arrayIndex++) ", func() {
		g.emitDeserialize(f, "[arrayIndex]", true)
	}, filedName)
}

func (g *jsGenerator) emitDeserialize(f *protokit.FieldDescriptor, suffix string, field bool) {
	prefix := ""
	if field {
		prefix = "this."
	}
	filedName := prefix + g.conv.GetFieldName(f.GetName()) + suffix
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("r.readNil();")
			g.e.EmitLine("%s = null;", filedName+suffix)
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("%s = new packer.proto.%s();", filedName+suffix, msgType)
			g.e.EmitLine("%s.read(r);", filedName+suffix)
		})
		break
	case "TYPE_ENUM":
		g.e.EmitLine("%s = r.readNumber();", filedName+suffix)
		break
	case "TYPE_BYTES":
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("r.readNil();")
			g.e.EmitLine("%s = null;", filedName+suffix)
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("%s = r.readBytes();", filedName+suffix)
		})
		break
	default:
		g.e.EmitLine("%s = r.read%s();", filedName, strings.Title(g.conv.GetTypeImpl(f)))
		break
	}
}
