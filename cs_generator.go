package main

import (
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	ce "github.com/yazawa-ichio/protoc-gen-msgpack/code_emitter"
)

type csClassGenerator struct {
	e    *ce.CodeEmitter
	conv *CSConverter
}

func newCSGenerator() *csClassGenerator {
	return &csClassGenerator{e: new(ce.CodeEmitter), conv: &CSConverter{}}
}

func (g *csClassGenerator) genClass(message *messageData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, message.file)
	g.e.EmitLine("using Writer = ILib.ProtoPack.Writer;")
	g.e.EmitLine("using Reader = ILib.ProtoPack.Reader;")
	g.emitClass(message)
	fileName := g.conv.GetClassName(message.data.GetName()) + ".cs"
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *csClassGenerator) genEnum(enum *enumData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, enum.file)
	g.emitEnum(enum)
	fileName := g.conv.GetEnumName(enum.data.GetName()) + ".cs"
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *csClassGenerator) emitClass(message *messageData) {
	//Header
	if message.parent == nil {
		//TODO:Optional use interface
		g.e.NewLine()
		if emitNameSpace(g.e, g.conv.GetPackageName(message.data.GetPackage())) {
			defer g.e.EndBracket("")
		}
	}

	//emit Enum
	emitSummary(g.e, message.data.GetComments().GetLeading())
	g.e.EmitLine("public partial class " + g.conv.GetClassName(message.data.GetName()))
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	for i, f := range message.data.GetMessageFields() {
		if i > 0 {
			g.e.NewLine()
		}
		emitSummary(g.e, f.GetComments().GetLeading())
		g.e.EmitTab()
		g.e.EmitAppend("public " + g.conv.GetType(f))
		g.e.EmitAppend(" " + g.conv.GetFieldName(f.GetName()) + " { get; set; }")
		g.e.NewLine()
	}

	for _, e := range message.enums {
		g.e.NewLine()
		g.emitEnum(e)
	}

	for _, m := range message.children {
		g.e.NewLine()
		g.emitClass(m)
	}

	g.emitWriter(message)
	g.emitReader(message)
}

func (g *csClassGenerator) emitEnum(enum *enumData) {
	//Header
	if enum.parent == nil {
		if emitNameSpace(g.e, g.conv.GetPackageName(enum.data.GetPackage())) {
			defer g.e.EndBracket("")
		}
	}
	//emit Enum
	emitSummary(g.e, enum.data.GetComments().GetLeading())
	g.e.EmitLine("public enum %s", g.conv.GetEnumTypeName(enum.data.GetName()))
	g.e.StartBracket("")
	for _, f := range enum.data.GetValues() {
		g.e.EmitLine("%s = %s,", g.conv.GetEnumName(f.GetName()), strconv.Itoa(int(f.GetNumber())))
	}
	g.e.EndBracket("")
}

func isObject(f *protokit.FieldDescriptor) bool {
	return f.GetType().String() == "TYPE_MESSAGE" ||
		f.GetType().String() == "TYPE_BYTES" ||
		f.GetLabel().String() == "LABEL_REPEATED"
}

func (g *csClassGenerator) emitWriter(message *messageData) {
	g.e.NewLine()
	emitSummary(g.e, "Serialize Message")
	g.e.EmitLine("public void Write(Writer w)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Write Map Length")
	g.e.EmitLine("w.WriteMapHeader(%s);", strconv.Itoa(len(message.data.GetMessageFields())))

	for _, f := range message.data.GetMessageFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := g.conv.GetFieldName(f.GetName())
		if isObject(f) {
			g.e.EmitLine("var %s = this.%s;", filedName, filedName)
		}
		g.e.EmitLine("w.WriteTag(%s);", strconv.Itoa(int(f.GetNumber())))
		if f.GetLabel().String() == "LABEL_REPEATED" {
			g.e.EmitLine("if (%s == null)", filedName)
			g.e.Bracket("", func() {
				g.e.EmitLine("w.WriteNil();")
			})
			g.e.EmitLine("else")
			g.e.Bracket("", func() {
				g.e.EmitLine("var arrayLen = %s.Length;", filedName)
				g.e.EmitLine("w.WriteArrayHeader(arrayLen);")
				g.e.EmitLine("for(var arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)")
				g.e.Bracket("", func() {
					g.emitSerialize(f, "[arrayIndex]")
				})
			})
		} else {
			g.emitSerialize(f, "")
		}
	}

}

func (g *csClassGenerator) emitSerialize(f *protokit.FieldDescriptor, suffix string) {
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.EmitLine("if (%s == default(%s))", g.conv.GetFieldName(f.GetName())+suffix, msgType)
		g.e.Bracket("", func() {
			g.e.EmitLine("w.WriteNil();")
		})
		g.e.EmitLine("else")
		g.e.Bracket("", func() {
			g.e.EmitLine("%s.Write(w);", g.conv.GetFieldName(f.GetName())+suffix)
		})
		break
	case "TYPE_ENUM":
		g.e.EmitLine("w.Write((int)%s);", g.conv.GetFieldName(f.GetName())+suffix)
		break
	case "TYPE_BYTES":
		g.emitBytesSerialize(f, suffix)
		break
	default:
		g.e.EmitLine("w.Write(%s);", g.conv.GetFieldName(f.GetName())+suffix)
		break
	}
}

func (g *csClassGenerator) emitBytesSerialize(f *protokit.FieldDescriptor, suffix string) {
	filedName := g.conv.GetFieldName(f.GetName()) + suffix
	g.e.EmitLine("if (%s == null)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("w.WriteNil();")
	})
	g.e.EmitLine("else")
	g.e.Bracket("", func() {
		g.e.EmitLine("var bufLen = %s.Length;", filedName)
		g.e.EmitLine("w.WriteArrayHeader(bufLen);")
		g.e.EmitLine("for(var bufIndex = 0; bufIndex < bufLen; bufIndex++)")
		g.e.Bracket("", func() {
			g.e.EmitLine("w.Write(%s[bufIndex]);", filedName)
		})
	})
}

func emitFileInfo(e *ce.CodeEmitter, file *protokit.FileDescriptor) {
	e.EmitLine("//%s", file.GetName())
}

func emitNameSpace(e *ce.CodeEmitter, pkgName string) bool {
	if pkgName != "" {
		e.EmitLine("namespace %s", pkgName)
		e.StartBracket("")
		return true
	}
	return false
}

func emitSummary(e *ce.CodeEmitter, comment string) bool {
	if comment == "" {
		return false
	}
	e.EmitLine("/// <summary>")
	for _, c := range strings.Split(comment, "\n") {
		e.EmitLine("/// " + c)
	}
	e.EmitLine("/// </summary>")
	return true
}

func (g *csClassGenerator) emitReader(message *messageData) {
	g.e.NewLine()
	emitSummary(g.e, "Deserialize Message")
	g.e.EmitLine("public void Read(Reader r, bool overridable = false)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("var mapLen = r.ReadMapHeader();")
	g.e.NewLine()
	g.e.EmitLine("for(int i = 0; i < mapLen; i++)")
	g.e.Bracket("", func() {
		g.e.EmitLine("var tag = r.ReadTag();")
		g.e.EmitLine("switch(tag) {")
		for _, f := range message.data.GetMessageFields() {
			g.e.Bracket("case %d:", func() {
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
		g.e.EmitLine("r.ReadSkip();")
		g.e.EmitLine("break;")
		g.e.EndIndent()
		g.e.EmitLine("}")
	})

}

func (g *csClassGenerator) emitRepeatedDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	g.e.EmitLine("if(r.IsNull())")
	g.e.Bracket("", func() {
		g.e.EmitLine("%s = r.ReadNil();", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("var %s = this.%s;", filedName, filedName)
	g.e.EmitLine("var arrayLen = r.ReadArrayHeader();")
	g.e.EmitLine("if(!overridable || %s == null)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("%s = r.NewArray<%s>(arrayLen);", filedName, g.conv.GetTypeImpl(f))
	})
	g.e.EmitLine("else if(%s.Length != arrayLen)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("System.Array.Resize(ref %s, arrayLen);", filedName)
	})
	g.e.EmitLine("for(int arrayIndex = 0; arrayIndex < arrayLen; arrayIndex++)")
	g.e.Bracket("", func() {
		g.emitDeserialize(f, "[arrayIndex]")
	})
}

func (g *csClassGenerator) emitDeserialize(f *protokit.FieldDescriptor, suffix string) {
	filedName := g.conv.GetFieldName(f.GetName())
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.EmitLine("if(r.IsNull())")
		g.e.Bracket("", func() {
			g.e.EmitLine("%s = r.ReadNil();", filedName+suffix)
			g.e.EmitLine("continue;")
		})
		g.e.EmitLine("if(!overridable || %s == null)", filedName+suffix)
		g.e.Bracket("", func() {
			g.e.EmitLine("%s = r.New<%s>();", filedName+suffix, msgType)
		})
		g.e.EmitLine("%s.Read(r, overridable);", filedName+suffix)
		break
	case "TYPE_ENUM":
		g.e.EmitLine("%s = (%s)r.ReadInt();", filedName+suffix, g.conv.GetTypeImpl(f))
		break
	case "TYPE_BYTES":
		g.e.EmitLine("if(r.IsNull())")
		g.e.Bracket("", func() {
			g.e.EmitLine("%s = r.ReadNil();", filedName+suffix)
			g.e.EmitLine("continue;")
		})
		g.e.EmitLine("r.ReadBytes(ref %s, overridable);", filedName+suffix)
		break
	default:
		g.e.EmitLine("%s = r.Read%s();", g.conv.GetFieldName(f.GetName())+suffix, strings.Title(g.conv.GetTypeImpl(f)))
		break
	}
}
