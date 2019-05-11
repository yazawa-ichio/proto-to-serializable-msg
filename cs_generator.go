package main

import (
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	ce "github.com/yazawa-ichio/protoc-gen-msgpack/code_emitter"
)

type csGenerator struct {
	e    *ce.CodeEmitter
	conv *CSConverter
	data *protoData
}

func newCSGenerator(data *protoData) *csGenerator {
	return &csGenerator{
		e:    new(ce.CodeEmitter),
		conv: newCSConverter(data),
		data: data,
	}
}

func (g *csGenerator) genResponseFile(data *protoData) []*plugin_go.CodeGeneratorResponse_File {
	files := []*plugin_go.CodeGeneratorResponse_File{}
	for _, msg := range data.messages {
		if msg.parent == nil {
			files = append(files, g.genClass(msg))
		}
	}
	for _, enum := range data.enums {
		if enum.parent == nil {
			files = append(files, g.genEnum(enum))
		}
	}
	return files
}

func (g *csGenerator) genClass(message *messageData) *plugin_go.CodeGeneratorResponse_File {
	g.e.Reset()
	emitFileInfo(g.e, message.file)
	g.e.EmitLine("using ILib.ProtoPack;")
	g.emitClass(message)
	fileName := g.conv.GetClassName(message.data.GetName()) + ".cs"
	content := g.e.String()
	return &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
	}
}

func (g *csGenerator) genEnum(enum *enumData) *plugin_go.CodeGeneratorResponse_File {
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

func (g *csGenerator) emitClass(message *messageData) {
	//Header
	if message.parent == nil {
		//TODO:Optional use interface
		g.e.NewLine()
		if emitNameSpace(g.e, g.conv.GetPackageName(message.data.GetPackage())) {
			defer g.e.EndBracket("")
		}
	}

	//emit Class
	emitSummary(g.e, message.data.GetComments().GetLeading())
	g.e.EmitLine("public partial class " + g.conv.GetClassName(message.data.GetName()) + " : IMessage ")
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
		if !m.mapEntry {
			g.e.NewLine()
			g.emitClass(m)
		}
	}

	g.e.NewLine()
	g.e.NewLine()
	g.e.EmitAppend("#region Serialization")
	g.e.NewLine()

	g.emitWriter(message)
	g.emitReader(message)

	g.e.EmitAppend("#endregion")
	g.e.NewLine()

}

func (g *csGenerator) emitEnum(enum *enumData) {
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

func (g *csGenerator) emitWriter(message *messageData) {
	g.e.NewLine()
	emitSummary(g.e, "Serialize Message")
	g.e.EmitLine("public void Write(IWriter w, bool skipable = true)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Write Map Length")
	g.e.Bracket("if (!skipable) ", func() {
		g.e.EmitLine("w.WriteMapHeader(%s);", strconv.Itoa(len(message.data.GetMessageFields())))
		g.e.EndAndStartBracket(" else ")
		g.e.EmitLine("int mapLen = 0;")
		for _, f := range message.data.GetMessageFields() {
			filedName := g.conv.GetFieldName(f.GetName())
			msgType := g.conv.GetType(f)
			g.e.EmitLine("if(this.%s != default(%s)) mapLen++;", filedName, msgType)
		}
		g.e.EmitLine("w.WriteMapHeader(mapLen);")
	})

	for _, f := range message.data.GetMessageFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := g.conv.GetFieldName(f.GetName())
		msgType := g.conv.GetType(f)
		g.e.StartBracket("if(!skipable || this.%s != default(%s)) ", filedName, msgType)

		if isObject(f) {
			g.e.EmitLine("var %s = this.%s;", filedName, filedName)
		}

		g.e.EmitLine("w.WriteTag(%s);", strconv.Itoa(int(f.GetNumber())))

		if g.data.isMapEntry(f) {
			g.e.EmitLine("if (%s == null)", filedName)
			g.e.Bracket("", func() {
				g.e.EmitLine("w.WriteNil();")
			})
			g.e.EmitLine("else")
			g.e.Bracket("", func() {
				g.e.EmitLine("var mapLen = %s.Count;", filedName)
				g.e.EmitLine("w.WriteMapHeader(mapLen);")
				g.e.Bracket("foreach(var _%sEntry in %s)", func() {
					mapKey, mapVal := g.data.getMapKeyValue(f)
					g.emitSerialize(mapKey, "_"+filedName+"Entry.", "")
					g.emitSerialize(mapVal, "_"+filedName+"Entry.", "")
				}, filedName, filedName)
			})

		} else if f.GetLabel().String() == "LABEL_REPEATED" {
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
					g.emitSerialize(f, "", "[arrayIndex]")
				})
			})
		} else {
			g.emitSerialize(f, "", "")
		}

		g.e.EndBracket("")
	}

}

func (g *csGenerator) emitSerialize(f *protokit.FieldDescriptor, prefix string, suffix string) {
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.EmitLine("if (%s == default(%s))", prefix+g.conv.GetFieldName(f.GetName())+suffix, msgType)
		g.e.Bracket("", func() {
			g.e.EmitLine("w.WriteNil();")
		})
		g.e.EmitLine("else")
		g.e.Bracket("", func() {
			g.e.EmitLine("%s.Write(w, skipable);", prefix+g.conv.GetFieldName(f.GetName())+suffix)
		})
		break
	case "TYPE_ENUM":
		g.e.EmitLine("w.Write((int)%s);", prefix+g.conv.GetFieldName(f.GetName())+suffix)
		break
	case "TYPE_BYTES":
		g.emitBytesSerialize(f, prefix, suffix)
		break
	default:
		g.e.EmitLine("w.Write(%s);", prefix+g.conv.GetFieldName(f.GetName())+suffix)
		break
	}
}

func (g *csGenerator) emitBytesSerialize(f *protokit.FieldDescriptor, prefix string, suffix string) {
	filedName := prefix + g.conv.GetFieldName(f.GetName()) + suffix
	g.e.EmitLine("if (%s == null)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("w.WriteNil();")
	})
	g.e.EmitLine("else")
	g.e.Bracket("", func() {
		g.e.EmitLine("w.WriteBytes(%s);", filedName)
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

func (g *csGenerator) emitReader(message *messageData) {
	g.e.NewLine()
	emitSummary(g.e, "Deserialize Message")
	g.e.EmitLine("public void Read(IReader r, bool overridable = false)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("var mapLen = r.ReadMapHeader();")
	g.e.EmitLine("uint tag = 0;")
	g.e.EmitLine("int index = 0;")
	g.e.NewLine()
	g.e.EmitLine("while ((tag = r.ReadTag(index++, mapLen)) != 0)")
	g.e.Bracket("", func() {
		g.e.EmitLine("switch(tag) {")
		for _, f := range message.data.GetMessageFields() {
			g.e.EmitLine("case %d:", f.GetNumber())
			g.e.StartIndent()
			if g.data.isMapEntry(f) {
				g.emitMapDeserialize(f)
			} else if f.GetLabel().String() == "LABEL_REPEATED" {
				g.emitRepeatedDeserialize(f)
			} else {
				g.emitDeserialize(f, "", "")
			}
			g.e.EmitLine("break;")
			g.e.EndIndent()
		}
		g.e.EmitLine("default:")
		g.e.StartIndent()
		g.e.EmitLine("r.Skip();")
		g.e.EmitLine("break;")
		g.e.EndIndent()
		g.e.EmitLine("}")
	})

}

func (g *csGenerator) emitMapDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	key, value := g.data.getMapKeyValue(f)
	g.e.EmitLine("if(r.IsNull())")
	g.e.Bracket("", func() {
		g.e.EmitLine("r.ReadNil();")
		g.e.EmitLine("this.%s = null;", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("var %s = this.%s;", filedName, filedName)
	g.e.EmitLine("if(!overridable || %s == null)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("%s = Provider.New<System.Collections.Generic.Dictionary<%s, %s>>();", filedName, g.conv.GetType(key), g.conv.GetType(value))
	})
	g.e.EmitLine("else")
	g.e.Bracket("", func() {
		g.e.EmitLine("%s.Clear();", filedName)
	})
	g.e.EmitLine("this.%s = %s;", filedName, filedName)
	g.e.EmitLine("var _%sLen = r.ReadMapHeader();", filedName)
	g.e.EmitLine("for(int mapIndex = 0; mapIndex < _%sLen; mapIndex++)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("%s _%sKey = default(%s);", g.conv.GetType(key), filedName, g.conv.GetType(key))
		g.e.EmitLine("%s _%sValue = default(%s);", g.conv.GetType(value), filedName, g.conv.GetType(value))
		g.emitDeserialize(key, "_"+filedName, "")
		g.emitDeserialize(value, "_"+filedName, "")
		g.e.EmitLine("%s[_%sKey] = _%sValue;", filedName, filedName, filedName)
	})
}

func (g *csGenerator) emitRepeatedDeserialize(f *protokit.FieldDescriptor) {
	filedName := g.conv.GetFieldName(f.GetName())
	g.e.EmitLine("if(r.IsNull())")
	g.e.Bracket("", func() {
		g.e.EmitLine("r.ReadNil();")
		g.e.EmitLine("this.%s = null;", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("var %s = this.%s;", filedName, filedName)
	g.e.EmitLine("var _%sLen = r.ReadArrayHeader();", filedName)
	g.e.EmitLine("if(!overridable || %s == null)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("%s = InstanceProvider.NewArray<%s>(_%sLen);", filedName, g.conv.GetTypeImpl(f), filedName)
	})
	g.e.EmitLine("else if(%s.Length != _%sLen)", filedName, filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("System.Array.Resize(ref %s, _%sLen);", filedName, filedName)
	})
	g.e.EmitLine("this.%s = %s;", filedName, filedName)
	g.e.EmitLine("for(int arrayIndex = 0; arrayIndex < _%sLen; arrayIndex++)", filedName)
	g.e.Bracket("", func() {
		g.emitDeserialize(f, "", "[arrayIndex]")
	})
}

func (g *csGenerator) emitDeserialize(f *protokit.FieldDescriptor, prefix string, suffix string) {
	filedName := g.conv.GetFieldName(f.GetName())
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.conv.GetType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.EmitLine("if(r.IsNull())")
		g.e.Bracket("", func() {
			g.e.EmitLine("r.ReadNil();")
			g.e.EmitLine("%s = default(%s);", prefix+filedName+suffix, g.conv.GetTypeImpl(f))
		})
		g.e.EmitLine(" else ")
		g.e.Bracket("", func() {
			g.e.EmitLine("if(!overridable || %s == default(%s))", prefix+filedName+suffix, g.conv.GetTypeImpl(f))
			g.e.Bracket("", func() {
				g.e.EmitLine("%s = InstanceProvider.New<%s>();", prefix+filedName+suffix, msgType)
			})
			g.e.EmitLine("%s.Read(r, overridable);", prefix+filedName+suffix)
		})
		break
	case "TYPE_ENUM":
		g.e.EmitLine("%s = (%s)r.ReadInt();", prefix+filedName+suffix, g.conv.GetTypeImpl(f))
		break
	case "TYPE_BYTES":
		g.e.EmitLine("if(r.IsNull())")
		g.e.Bracket("", func() {
			g.e.EmitLine("r.ReadNil();")
			if prefix == "" {
				g.e.EmitLine("this.%s = null;", filedName+suffix)
			} else {
				g.e.EmitLine("%s = null;", prefix+filedName+suffix)
			}
		})
		g.e.EmitLine(" else ")
		g.e.Bracket("", func() {
			if prefix == "" {
				g.e.EmitLine("var %s = this.%s;", filedName+suffix, filedName+suffix)
				g.e.EmitLine("r.ReadBytes(ref %s, overridable);", filedName+suffix)
				g.e.EmitLine("this.%s = %s;", filedName+suffix, filedName+suffix)
			} else {
				g.e.EmitLine("r.ReadBytes(ref %s, overridable);", prefix+filedName+suffix)
			}
		})
		break
	default:
		g.e.EmitLine("%s = r.Read%s();", prefix+g.conv.GetFieldName(f.GetName())+suffix, strings.Title(g.conv.GetTypeImpl(f)))
		break
	}
}
