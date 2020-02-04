package proto

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jhump/protoreflect/desc"
	ce "github.com/yazawa-ichio/protoc-gen-msgpack/code_emitter"
)

type JSGenerator struct {
	e                      *ce.CodeEmitter
	SkipSerializer         bool
	PackageNameToDirectory bool
}

func NewJSGenerator() *JSGenerator {
	return &JSGenerator{
		e:                      new(ce.CodeEmitter),
		PackageNameToDirectory: true,
	}
}

func (g *JSGenerator) Generate(files []string) ([]*GenerateFile, error) {
	output := make([]*GenerateFile, 0)
	parsed, err := parseFiles(files)
	if err != nil {
		return nil, err
	}
	packages := make(map[string][]desc.Descriptor, 0)
	for _, f := range parsed {
		list, ok := packages[f.GetPackage()]
		if !ok {
			list = make([]desc.Descriptor, 0)
		}
		for _, m := range f.GetMessageTypes() {
			list = append(list, m)
			output = append(output, g.genClass(m))
		}
		for _, e := range f.GetEnumTypes() {
			list = append(list, e)
			output = append(output, g.genEnum(e))
		}
		packages[f.GetPackage()] = list
	}
	output = append(output, g.genIndex(packages)...)
	return output, nil
}

func (g *JSGenerator) GenerateAndOutput(files []string, outputRoot string) error {
	return output(g, files, outputRoot)
}

func toJsClassCast(s string) string {
	return toCamel(s, false)
}

func toJsFieldCast(s string) string {
	return toCamel(s, true)
}

func (g *JSGenerator) toEnumName(d *desc.EnumValueDescriptor) string {
	return strings.ToUpper(d.GetName())
}

func (g *JSGenerator) getJsFieldTypeFileName(d *desc.FieldDescriptor) string {
	return strings.ToLower(toJsClassCast(g.getTypeImpl(d)) + ".js")
}

func (g *JSGenerator) getJsFileName(d desc.Descriptor) string {
	name := d.GetFullyQualifiedName()
	if g.PackageNameToDirectory && d.GetFile().GetPackage() != "" {
		name = name[len(d.GetFile().GetPackage())+1:]
	}
	return strings.ToLower(toJsClassCast(name) + ".js")
}

func (g *JSGenerator) genClass(message *desc.MessageDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, message)
	g.e.EmitLine("\"use strict\";")
	g.e.EmitLine("const _packer = require('proto-msgpack');")
	packageName := message.GetFile().GetPackage()
	if !g.PackageNameToDirectory || packageName == "" {
		g.e.EmitLine("const _proto = require('./index.js');")
	} else {
		g.e.EmitLine("const _proto = require('./../index.js');")
	}
	//	g.emitMesssageDeps(message)
	g.emitClass(message, "")
	fileName := g.getJsFileName(message)
	if g.PackageNameToDirectory && packageName != "" {
		fileName = strings.ToLower(toJsClassCast(packageName)) + "/" + fileName
	}
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *JSGenerator) genEnum(enum *desc.EnumDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, enum)
	g.e.EmitLine("\"use strict\";")
	g.emitEnum(enum, "")
	fileName := g.getJsFileName(enum)
	packageName := enum.GetFile().GetPackage()
	if g.PackageNameToDirectory && packageName != "" {
		fileName = strings.ToLower(toJsClassCast(packageName)) + "/" + fileName
	}
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *JSGenerator) genIndex(packages map[string][]desc.Descriptor) []*GenerateFile {
	keys := make([]string, 0)
	for key := range packages {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	output := make([]*GenerateFile, 0)

	for i, key := range keys {
		if g.PackageNameToDirectory || i == 0 {
			g.e.Reset()
			g.e.EmitLine("\"use strict\";")
			g.e.NewLine()
		}

		if !g.PackageNameToDirectory && key != "" {
			g.e.EmitLine("exports.%s = {}", key)
			g.e.NewLine()
		}
		for _, d := range packages[key] {
			g.emitDescComment(d)
			if !g.PackageNameToDirectory && key != "" {
				g.e.EmitLine("exports.%s.%s = require(\"./%s\")", key, toJsClassCast(d.GetName()), g.getJsFileName(d))
				g.e.NewLine()
			} else {
				g.e.EmitLine("exports.%s = require(\"./%s\")", toJsClassCast(d.GetName()), g.getJsFileName(d))
				g.e.NewLine()
			}
		}
		if g.PackageNameToDirectory && key == "" {
			for _, pkgName := range keys {
				if pkgName != "" {
					g.e.EmitLine("exports.%s = require(\"./%s\")", pkgName, strings.ToLower(toJsClassCast(pkgName)))
					g.e.NewLine()
				}
			}
		}
		if g.PackageNameToDirectory || len(keys)-1 == i {
			content := g.e.String()
			fileName := "index.js"
			if g.PackageNameToDirectory && key != "" {
				fileName = strings.ToLower(toJsClassCast(key)) + "/index.js"
			}
			output = append(output, &GenerateFile{
				Name:    fileName,
				Content: content,
			})
		}
	}

	return output
}

func (g *JSGenerator) emitMesssageDeps(m *desc.MessageDescriptor) {
	hits := make(map[string]string)
	for _, f := range m.GetFields() {
		if f.IsMap() {
			key, val := f.GetMapKeyType(), f.GetMapValueType()
			g.emitRequire(hits, key)
			g.emitRequire(hits, val)
		} else {
			g.emitRequire(hits, f)
		}
	}
	for _, nm := range m.GetNestedMessageTypes() {
		if !nm.IsMapEntry() {
			g.e.EmitLine("require('./%s');", g.getJsFileName(nm))
		}
	}
	for _, ne := range m.GetNestedEnumTypes() {
		g.e.EmitLine("require('./%s');", g.getJsFileName(ne))
	}
}

func (g *JSGenerator) emitRequire(hits map[string]string, f *desc.FieldDescriptor) {
	if !isUserDefine(f) {
		return
	}
	name := f.AsFieldDescriptorProto().GetTypeName()
	if _, hit := hits[name]; hit {
		return
	}
	hits[name] = name
	g.e.EmitLine("require('./%s');", g.getJsFieldTypeFileName(f))
}

func (g *JSGenerator) emitClass(message *desc.MessageDescriptor, parent string) {
	g.emitDescComment(message)
	name := toJsClassCast(message.GetName())
	if parent == "" {
		g.e.EmitLine("class %s {", name)
	} else {
		g.e.EmitLine("%s.%s = class %s {", parent, name, name)
	}
	g.e.StartIndent()
	g.e.Bracket("constructor(init, pos) ", func() {
		for _, f := range message.GetFields() {
			fieldName := toJsFieldCast(f.GetName())
			if f.GetLabel().String() == "LABEL_REPEATED" {
				g.e.EmitLine("this.%s = null;", fieldName)
				continue
			}
			switch f.GetType().String() {
			case "TYPE_ENUM", "TYPE_FLOAT", "TYPE_DOUBLE":
				g.e.EmitLine("this.%s = 0;", fieldName)
				break
			case "TYPE_MESSAGE":
				g.e.Bracket("if(init == null || init == true)", func() {
					g.e.EmitLine("this.%s = new _proto.%s();", fieldName, g.getType(f))
					g.e.EndAndStartBracket(" else ")
					g.e.EmitLine("this.%s = null;", fieldName)
				})
				break
			default:
				if g.getTypeImpl(f) == "number" {
					g.e.EmitLine("this.%s = 0;", fieldName)
				} else {
					g.e.EmitLine("this.%s = null;", fieldName)
				}
				break
			}
		}
		g.e.Bracket("if(Buffer.isBuffer(init))", func() {
			g.e.EmitLine("this.read(new _packer.ProtoReader(init, pos));")
		})
	})

	if !g.SkipSerializer {
		g.e.Bracket("pack() ", func() {
			g.e.EmitLine("const w = _packer.defaultWriter;")
			g.e.EmitLine("w.clear();")
			g.e.EmitLine("this.write(w);")
			g.e.EmitLine("return w.toBuffer();")
		})
		g.e.Bracket("unpack(buf, pos) ", func() {
			g.e.Bracket("if(!Buffer.isBuffer(buf))", func() {
				g.e.EmitLine("this.read(buf);")
				g.e.EndAndStartBracket(" else ")
				g.e.EmitLine("this.read(new _packer.ProtoReader(buf, pos));")
			})
		})
		g.e.Bracket("write(w) ", func() {
			g.emitWriter(message)
		})
		g.e.Bracket("read(r) ", func() {
			g.emitReader(message)
		})
	}

	g.e.EndBracket("")

	if parent == "" {
		g.e.EmitLine("module.exports = %s", name)
		parent = name
	} else {
		parent = parent + "." + name
	}

	for _, nm := range message.GetNestedMessageTypes() {
		if !nm.IsMapEntry() {
			g.emitClass(nm, parent)
		}
	}
	for _, ne := range message.GetNestedEnumTypes() {
		g.emitEnum(ne, parent)
	}

}

func (g *JSGenerator) emitEnum(enum *desc.EnumDescriptor, parent string) {
	g.emitDescComment(enum)
	//g.e.StartIndent()
	if parent == "" {
		g.e.StartBracket("const %s = ", enum.GetName())
	} else {
		g.e.StartBracket("%s.%s = ", parent, enum.GetName())
	}
	vals := enum.GetValues()
	for i, v := range vals {
		g.emitDescComment(v)
		if i < len(vals)-1 {
			g.e.EmitLine("%s: %d,", g.toEnumName(v), v.GetNumber())
		} else {
			g.e.EmitLine("%s: %d", g.toEnumName(v), v.GetNumber())
		}
	}
	g.e.EndBracket(";")
	if parent == "" {
		g.emitDescComment(enum)
		g.e.EmitLine("module.exports = %s", enum.GetName())
	}
}

func (g *JSGenerator) emitDescComment(d desc.Descriptor) bool {
	return g.emitComment(d.GetSourceInfo().GetLeadingComments())
}

func (g *JSGenerator) emitComment(comment string) bool {
	if comment == "" {
		return false
	}
	if comment[len(comment)-1] == '\n' {
		comment = comment[:len(comment)-1]
	}
	g.e.EmitLine("/**")
	for _, c := range strings.Split(comment, "\n") {
		g.e.EmitLine(" * %s", c)
	}
	g.e.EmitLine(" */")
	return true
}

func (g *JSGenerator) getType(f *desc.FieldDescriptor) string {
	if f.IsRepeated() {
		return g.getTypeImpl(f) + "[]"
	}
	return g.getTypeImpl(f)
}

func (g *JSGenerator) getTypeImpl(f *desc.FieldDescriptor) string {
	switch f.GetType().String() {
	case "TYPE_INT64", "TYPE_UINT64", "TYPE_INT32", "TYPE_FIXED64", "TYPE_FIXED32", "TYPE_UINT32", "TYPE_SFIXED32", "TYPE_SFIXED64", "TYPE_SINT32", "TYPE_SINT64":
		return "number"
	case "TYPE_DOUBLE":
		return "double"
	case "TYPE_FLOAT":
		return "float"
	case "TYPE_BOOL":
		return "bool"
	case "TYPE_STRING":
		return "string"
	case "TYPE_BYTES":
		return "Uint8Array"
	}
	name := f.AsFieldDescriptorProto().GetTypeName()
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	return toJsClassCast(name)
}

func (g *JSGenerator) emitWriter(message *desc.MessageDescriptor) {
	g.e.EmitLine("// Write Map Length")
	g.e.EmitLine("w.writeMapHeader(%s);", strconv.Itoa(len(message.GetFields())))
	for _, f := range message.GetFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := toJsFieldCast(f.GetName())
		g.e.EmitLine("w.writeTag(%s);", strconv.Itoa(int(f.GetNumber())))

		if f.IsMap() {
			g.e.Bracket("if (this.%s == null) ", func() {
				g.e.EmitLine("w.writeNil();")
				g.e.EndAndStartBracket(" else ")
				g.e.EmitLine("const mapLen = this.%s.size;", filedName)
				g.e.EmitLine("w.writeMapHeader(mapLen);")
				g.e.StartBracket("this.%s.forEach(function(value, key)", filedName)
				mapKey, mapVal := f.GetMapKeyType(), f.GetMapValueType()
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

func (g *JSGenerator) emitSerialize(f *desc.FieldDescriptor, suffix string, field bool) {
	prefix := ""
	if field {
		prefix = "this."
	}
	filedName := prefix + toJsFieldCast(f.GetName()) + suffix
	typeName := strings.Title(g.getTypeImpl(f))
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

func (g *JSGenerator) emitReader(message *desc.MessageDescriptor) {
	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("const mapLen = r.readMapHeader();")
	g.e.NewLine()
	g.e.Bracket("for(let i = 0; i < mapLen; i++) ", func() {
		g.e.EmitLine("const tag = r.readTag();")
		g.e.EmitLine("switch(tag) {")
		for _, f := range message.GetFields() {
			g.e.EmitLine("case %d:", f.GetNumber())
			g.e.StartIndent()
			if f.IsMap() {
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

func (g *JSGenerator) emitMapDeserialize(f *desc.FieldDescriptor) {
	filedName := toJsFieldCast(f.GetName())
	key, value := f.GetMapKeyType(), f.GetMapValueType()
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

func (g *JSGenerator) emitRepeatedDeserialize(f *desc.FieldDescriptor) {
	filedName := toJsFieldCast(f.GetName())
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

func (g *JSGenerator) emitDeserialize(f *desc.FieldDescriptor, suffix string, field bool) {
	prefix := ""
	if field {
		prefix = "this."
	}
	filedName := prefix + toJsFieldCast(f.GetName()) + suffix
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		msgType := g.getType(f)
		if strings.HasSuffix(msgType, "[]") {
			msgType = msgType[0 : len(msgType)-2]
		}
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("r.readNil();")
			g.e.EmitLine("%s = null;", filedName)
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("%s = new _proto.%s(false);", filedName, msgType)
			g.e.EmitLine("%s.read(r);", filedName)
		})
		break
	case "TYPE_ENUM":
		g.e.EmitLine("%s = r.readNumber();", filedName)
		break
	case "TYPE_BYTES":
		g.e.Bracket("if(r.isNull()) ", func() {
			g.e.EmitLine("r.readNil();")
			g.e.EmitLine("%s = null;", filedName)
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("%s = r.readBytes();", filedName)
		})
		break
	default:
		g.e.EmitLine("%s = r.read%s();", filedName, strings.Title(g.getTypeImpl(f)))
		break
	}
}
