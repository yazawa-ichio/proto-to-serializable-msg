package proto

import (
	"strconv"
	"strings"

	"github.com/jhump/protoreflect/desc"
	ce "github.com/yazawa-ichio/protoc-gen-msgpack/code_emitter"
)

type CSGenerator struct {
	e              *ce.CodeEmitter
	SkipSerializer bool
	Property       bool
	Serializable   bool
}

func toCSharpCast(s string) string {
	return toCamel(s, false)
}

func NewCSGenerator() *CSGenerator {
	return &CSGenerator{
		e:            new(ce.CodeEmitter),
		Serializable: true,
	}
}

func (g *CSGenerator) Generate(files []string) ([]*GenerateFile, error) {
	output := make([]*GenerateFile, 0)
	parsed, err := parseFiles(files)
	if err != nil {
		return nil, err
	}
	for _, f := range parsed {
		for _, m := range f.GetMessageTypes() {
			output = append(output, g.genClass(m))
		}
		for _, e := range f.GetEnumTypes() {
			output = append(output, g.genEnum(e))
		}
	}
	return output, nil
}

func (g *CSGenerator) GenerateAndOutput(files []string, outputRoot string) error {
	return output(g, files, outputRoot)
}

func (g *CSGenerator) genClass(message *desc.MessageDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, message)
	g.e.EmitLine("using ILib.ProtoPack;")
	g.e.EmitLine("using System.Collections.Generic;")
	if message.GetFile().GetPackage() != "" {
		g.e.NewLine()
		g.e.EmitLine("namespace %s", toCSharpCast(message.GetFile().GetPackage()))
		g.e.StartBracket("")
	}
	g.emitClass(message)
	if message.GetFile().GetPackage() != "" {
		g.e.EndBracket("")
	}
	var fileName string
	if message.GetFile().GetPackage() == "" {
		fileName = toCSharpCast(message.GetName()) + ".cs"
	} else {
		fileName = toCSharpCast(message.GetFile().GetPackage()+"."+message.GetName()) + ".cs"
	}
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *CSGenerator) genEnum(enum *desc.EnumDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, enum)
	if enum.GetFile().GetPackage() != "" {
		g.e.NewLine()
		g.e.EmitLine("namespace %s", toCSharpCast(enum.GetFile().GetPackage()))
		g.e.StartBracket("")
	}
	g.emitEnum(enum)
	if enum.GetFile().GetPackage() != "" {
		g.e.EndBracket("")
	}
	var fileName string
	if enum.GetFile().GetPackage() == "" {
		fileName = toCSharpCast(enum.GetName()) + ".cs"
	} else {
		fileName = toCSharpCast(enum.GetFile().GetPackage()+"."+enum.GetName()) + ".cs"
	}
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *CSGenerator) emitEnum(enum *desc.EnumDescriptor) {
	//emit Enum
	g.emitComment(enum)
	g.e.EmitLine("public enum %s", toCSharpCast(enum.GetName()))
	g.e.StartBracket("")
	for _, f := range enum.GetValues() {
		g.emitComment(f)
		g.e.EmitLine("%s = %v,", toCSharpCast(f.GetName()), f.GetNumber())
	}
	g.e.EndBracket("")
}

func (g *CSGenerator) emitComment(descriptor desc.Descriptor) bool {
	return g.emitSummary(descriptor.GetSourceInfo().GetLeadingComments())
}

func (g *CSGenerator) emitSummary(comment string) bool {
	if len(comment) == 0 {
		return false
	}
	if comment[len(comment)-1] == '\n' {
		comment = comment[:len(comment)-1]
	}
	g.e.EmitLine("/// <summary>")
	for _, c := range strings.Split(comment, "\n") {
		g.e.EmitLine("/// " + c)
	}
	g.e.EmitLine("/// </summary>")
	return true
}

func (g *CSGenerator) emitClass(message *desc.MessageDescriptor) {

	//emit Class
	g.emitComment(message)
	if g.Serializable {
		g.e.EmitLine("[System.Serializable]")
	}
	if g.SkipSerializer {
		g.e.EmitLine("public partial class %s", toCSharpCast(message.GetName()))
	} else {
		g.e.EmitLine("public partial class %s : IMessage", toCSharpCast(message.GetName()))
	}
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	for i, f := range message.GetFields() {
		if i > 0 {
			g.e.NewLine()
		}
		g.emitComment(f)
		g.e.EmitTab()
		g.e.EmitAppend("public " + g.getType(f))
		g.e.EmitAppend(" " + toCSharpCast(f.GetName()))
		if g.Property {
			g.e.EmitAppend(" { get; set; }")
		} else {
			g.e.EmitAppend(";")
		}
		g.e.NewLine()
	}

	for _, e := range message.GetNestedEnumTypes() {
		g.e.NewLine()
		g.emitEnum(e)
	}

	for _, m := range message.GetNestedMessageTypes() {
		if !m.IsMapEntry() {
			g.e.NewLine()
			g.emitClass(m)
		}
	}

	if !g.SkipSerializer {
		g.e.NewLine()
		g.e.EmitLine("#region Serialization")
		g.emitWriter(message)
		g.emitReader(message)
		g.e.EmitLine("#endregion")
		g.e.NewLine()
	}

}

func (g *CSGenerator) getType(f *desc.FieldDescriptor) string {
	if f.IsMap() {
		key := f.GetMapKeyType()
		val := f.GetMapValueType()
		return "Dictionary<" + g.getType(key) + ", " + g.getType(val) + ">"
	}
	if f.IsRepeated() {
		return g.getTypeImpl(f) + "[]"
	}
	return g.getTypeImpl(f)
}

func (g *CSGenerator) getTypeImpl(f *desc.FieldDescriptor) string {
	switch f.GetType().String() {
	case "TYPE_DOUBLE":
		return "double"
	case "TYPE_FLOAT":
		return "float"
	case "TYPE_INT64":
		return "long"
	case "TYPE_UINT64":
		return "ulong"
	case "TYPE_INT32":
		return "int"
	case "TYPE_FIXED64":
		return "ulong"
	case "TYPE_FIXED32":
		return "uint"
	case "TYPE_BOOL":
		return "bool"
	case "TYPE_STRING":
		return "string"
	case "TYPE_BYTES":
		return "byte[]"
	case "TYPE_UINT32":
		return "uint"
	case "TYPE_SFIXED32":
		return "int"
	case "TYPE_SFIXED64":
		return "long"
	case "TYPE_SINT32":
		return "int"
	case "TYPE_SINT64":
		return "long"
	}
	s := f.AsFieldDescriptorProto().GetTypeName()
	if strings.HasPrefix(s, ".") {
		s = s[1:]
	}
	return toCSharpCast(s)
}

func (g *CSGenerator) emitWriter(message *desc.MessageDescriptor) {
	g.e.NewLine()
	g.emitSummary("Serialize Message")
	g.e.EmitLine("public void Write(IWriter w)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Write Map Length")
	g.e.EmitLine("w.WriteMapHeader(%s);", strconv.Itoa(len(message.GetFields())))

	for _, f := range message.GetFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := toCSharpCast(f.GetName())

		g.e.EmitLine("w.WriteTag(%s);", strconv.Itoa(int(f.GetNumber())))

		if f.IsMap() {
			g.e.EmitLine("if (%s == null)", filedName)
			g.e.Bracket("", func() {
				g.e.EmitLine("w.WriteNil();")
			})
			g.e.EmitLine("else")
			g.e.Bracket("", func() {
				g.e.EmitLine("var mapLen = %s.Count;", filedName)
				g.e.EmitLine("w.WriteMapHeader(mapLen);")
				g.e.Bracket("foreach(var _%sEntry in %s)", func() {
					mapKey, mapVal := f.GetMapKeyType(), f.GetMapValueType()
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
	}

}

func (g *CSGenerator) emitSerialize(f *desc.FieldDescriptor, prefix string, suffix string) {
	name := prefix + toCSharpCast(f.GetName()) + suffix
	switch f.GetType().String() {
	case "TYPE_ENUM":
		g.e.EmitLine("w.Write((int)%s);", name)
		break
	default:
		g.e.EmitLine("w.Write(%s);", name)
		break
	}
}

func (g *CSGenerator) emitReader(message *desc.MessageDescriptor) {
	g.e.NewLine()
	g.emitSummary("Deserialize Message")
	g.e.EmitLine("public void Read(IReader r)")
	g.e.StartBracket("")
	defer g.e.EndBracket("")

	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("var len = r.ReadMapHeader();")
	g.e.NewLine()
	g.e.EmitLine("for (var i = 0; i < len; i++)")
	g.e.Bracket("", func() {
		g.e.EmitLine("var tag = r.ReadTag();")
		g.e.EmitLine("switch(tag) {")
		for _, f := range message.GetFields() {
			g.e.EmitLine("case %d: // Read %s", f.GetNumber(), f.GetName())
			g.e.StartIndent()
			if f.IsMap() {
				g.emitMapDeserialize(f)
			} else if f.IsRepeated() {
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

func (g *CSGenerator) emitMapDeserialize(f *desc.FieldDescriptor) {
	filedName := toCSharpCast(f.GetName())
	key, value := f.GetMapKeyType(), f.GetMapValueType()
	g.e.EmitLine("if(r.NextFormatIsNull())")
	g.e.Bracket("", func() {
		g.e.EmitLine("r.ReadNil();")
		g.e.EmitLine("%s = null;", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("var _%sLen = r.ReadMapHeader();", filedName)
	g.e.EmitLine("%s = new Dictionary<%s, %s>(_%sLen);", filedName, g.getType(key), g.getType(value), filedName)
	g.e.EmitLine("for(int mapIndex = 0; mapIndex < _%sLen; mapIndex++)", filedName)
	g.e.Bracket("", func() {
		g.e.EmitLine("var _%sKey = default(%s);", filedName, g.getType(key))
		g.e.EmitLine("var _%sValue = default(%s);", filedName, g.getType(value))
		g.emitDeserialize(key, "_"+filedName, "")
		g.emitDeserialize(value, "_"+filedName, "")
		g.e.EmitLine("%s[_%sKey] = _%sValue;", filedName, filedName, filedName)
	})
}

func (g *CSGenerator) emitRepeatedDeserialize(f *desc.FieldDescriptor) {
	filedName := toCSharpCast(f.GetName())
	g.e.EmitLine("if(r.NextFormatIsNull())")
	g.e.Bracket("", func() {
		g.e.EmitLine("r.ReadNil();")
		g.e.EmitLine("this.%s = null;", filedName)
		g.e.EmitLine("continue;")
	})
	g.e.EmitLine("var _%sLen = r.ReadArrayHeader();", filedName)
	if g.getTypeImpl(f) != "byte[]" {
		g.e.EmitLine("%s = new %s[_%sLen];", filedName, g.getTypeImpl(f), filedName)
	} else {
		g.e.EmitLine("%s = new byte[_%sLen][];", filedName, filedName)
	}
	g.e.EmitLine("for(int arrayIndex = 0; arrayIndex < _%sLen; arrayIndex++)", filedName)
	g.e.Bracket("", func() {
		g.emitDeserialize(f, "", "[arrayIndex]")
	})
}

func (g *CSGenerator) emitDeserialize(f *desc.FieldDescriptor, prefix string, suffix string) {
	filedName := toCSharpCast(f.GetName())
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.getType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.EmitLine("%s = r.ReadMessage<%s>();", prefix+filedName+suffix, g.getTypeImpl(f))
		break
	case "TYPE_ENUM":
		g.e.EmitLine("%s = (%s)r.ReadInt();", prefix+filedName+suffix, g.getTypeImpl(f))
		break
	case "TYPE_BYTES":
		g.e.EmitLine("%s = r.ReadBytes();", prefix+filedName+suffix)
		break
	default:
		g.e.EmitLine("%s = r.Read%s();", prefix+toCSharpCast(f.GetName())+suffix, strings.Title(g.getTypeImpl(f)))
		break
	}
}
