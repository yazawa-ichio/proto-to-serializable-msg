package proto

import (
	"sort"
	"strings"

	"github.com/jhump/protoreflect/desc"
	ce "github.com/yazawa-ichio/proto-to-serializable-msg/code_emitter"
)

type TSGenerator struct {
	e                      *ce.CodeEmitter
	SkipSerializer         bool
	PackageNameToDirectory bool
}

func NewTSGenerator() *TSGenerator {
	return &TSGenerator{
		e:                      new(ce.CodeEmitter),
		PackageNameToDirectory: true,
	}
}

func (g *TSGenerator) Generate(files []string) ([]*GenerateFile, error) {
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
		}
		for _, e := range f.GetEnumTypes() {
			list = append(list, e)
		}
		packages[f.GetPackage()] = list
	}
	output = append(output, g.genIndex(packages)...)
	return output, nil
}

func (g *TSGenerator) GenerateAndOutput(files []string, outputRoot string) error {
	return output(g, files, outputRoot)
}

func (g *TSGenerator) getTsFileName(d desc.Descriptor) string {
	return strings.ToLower(toJsClassCast(d.GetFullyQualifiedName()) + ".d.ts")
}

func (g *TSGenerator) toEnumName(d *desc.EnumValueDescriptor) string {
	return strings.ToUpper(d.GetName())
}

func (g *TSGenerator) formFileName(f *desc.FieldDescriptor) string {
	name := f.AsFieldDescriptorProto().GetTypeName()
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	return strings.ToLower(toJsClassCast(name))
}

func (g *TSGenerator) importName(f *desc.FieldDescriptor) string {
	return strings.Replace(g.formFileName(f), ".", "_", -1)
}

func (g *TSGenerator) genIndex(packages map[string][]desc.Descriptor) []*GenerateFile {

	keys := make([]string, 0)
	for key := range packages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	output := make([]*GenerateFile, 0)

	for i, key := range keys {
		if g.PackageNameToDirectory || i == 0 {
			g.e.Reset()
			g.e.EmitLine("/// <reference types=\"node\" />")
			g.e.EmitLine("import * as packer from 'proto-msgpack'")
		}
		if !g.PackageNameToDirectory && key != "" {
			g.e.NewLine()
			g.e.StartBracket("export namespace %s ", key)
		}
		if g.PackageNameToDirectory {
			for _, pkgName := range keys {
				if pkgName != key {
					if key == "" {
						g.e.EmitLine("import * as %s from './%s'", toJsClassCast(pkgName), strings.ToLower(toJsClassCast(pkgName)))
					} else {
						if pkgName != "" {
							g.e.EmitLine("import * as %s from '../%s'", toJsClassCast(pkgName), strings.ToLower(toJsClassCast(pkgName)))
						} else {
							g.e.EmitLine("import * as proto from '../index'")
						}
					}
				}
			}
			g.e.NewLine()
		}
		for _, d := range packages[key] {
			if message, ok := d.(*desc.MessageDescriptor); ok {
				g.emitClass(message, key == "" || g.PackageNameToDirectory)
			}
			if enum, ok := d.(*desc.EnumDescriptor); ok {
				g.emitEnum(enum, key == "" || g.PackageNameToDirectory)
			}
		}
		if !g.PackageNameToDirectory && key != "" {
			g.e.EndBracket("")
		}

		if g.PackageNameToDirectory || len(keys)-1 == i {
			content := g.e.String()
			fileName := "index.d.ts"
			if g.PackageNameToDirectory && key != "" {
				fileName = strings.ToLower(toJsClassCast(key)) + "/index.d.ts"
			}
			output = append(output, &GenerateFile{
				Name:    fileName,
				Content: content,
			})
		}
	}

	return output
}

func (g *TSGenerator) genClass(message *desc.MessageDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, message)
	g.emitDeps(message)
	g.emitClass(message, true)
	fileName := g.getTsFileName(message)
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *TSGenerator) genEnum(enum *desc.EnumDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, enum)
	g.emitEnum(enum, true)
	fileName := g.getTsFileName(enum)
	content := g.e.String()
	return &GenerateFile{
		Name:    fileName,
		Content: content,
	}
}

func (g *TSGenerator) emitDeps(m *desc.MessageDescriptor) {
	g.e.EmitLine("/// <reference types=\"node\" />")
	g.e.EmitLine("import * as packer from 'proto-msgpack'")
	hits := make(map[string]string)
	for _, f := range m.GetFields() {
		if f.IsMap() {
			key, val := f.GetMapKeyType(), f.GetMapValueType()
			g.emitImport(hits, key)
			g.emitImport(hits, val)
		} else {
			g.emitImport(hits, f)
		}
	}
}

func (g *TSGenerator) emitImport(hits map[string]string, f *desc.FieldDescriptor) {
	if !isUserDefine(f) {
		return
	}
	name := f.AsFieldDescriptorProto().GetTypeName()
	if _, hit := hits[name]; hit {
		return
	}
	hits[name] = name
	g.e.EmitLine("import %s from './%s';", g.importName(f), g.formFileName(f))
}

func (g *TSGenerator) getType(f *desc.FieldDescriptor) string {
	if f.IsMap() {
		key, val := f.GetMapKeyType(), f.GetMapValueType()
		return "Map<" + g.getType(key) + ", " + g.getType(val) + ">"
	}
	repeated := f.GetLabel().String() == "LABEL_REPEATED"
	if repeated {
		return "Array<" + g.getTypeImpl(f) + ">"
	}
	return g.getTypeImpl(f)
}

func (g *TSGenerator) getTypeImpl(f *desc.FieldDescriptor) string {
	typePackageName := ""
	switch f.GetType().String() {
	case "TYPE_INT64", "TYPE_UINT64", "TYPE_INT32", "TYPE_FIXED64", "TYPE_FIXED32", "TYPE_UINT32", "TYPE_SFIXED32", "TYPE_SFIXED64", "TYPE_SINT32", "TYPE_SINT64", "TYPE_DOUBLE", "TYPE_FLOAT":
		return "number"
	case "TYPE_BOOL":
		return "boolean"
	case "TYPE_STRING":
		return "string"
	case "TYPE_BYTES":
		return "Uint8Array | null"
	case "TYPE_ENUM":
		typePackageName = f.GetEnumType().GetFile().GetPackage()
		break
	case "TYPE_MESSAGE":
		typePackageName = f.GetMessageType().GetFile().GetPackage()
		break
	}
	samePackage := typePackageName == f.GetFile().GetPackage()
	name := f.AsFieldDescriptorProto().GetTypeName()
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	if g.PackageNameToDirectory && samePackage && typePackageName != "" {
		name = name[len(typePackageName)+1:]
	}
	name = toJsClassCast(name)
	if g.PackageNameToDirectory && !samePackage && typePackageName == "" {
		name = "proto." + name
	}
	if f.GetType().String() == "TYPE_MESSAGE" {
		name = name + " | null"
	}
	return name
}

func (g *TSGenerator) emitClass(message *desc.MessageDescriptor, root bool) {
	g.emitDescComment(message)
	if root {
		g.e.StartBracket("export class %s", toJsClassCast(message.GetName()))
	} else {
		g.e.StartBracket("class %s", toJsClassCast(message.GetName()))
	}
	for _, f := range message.GetFields() {
		g.emitDescComment(f)
		typeName := g.getType(f)
		g.e.EmitLine("%s: %s;", toJsFieldCast(f.GetName()), typeName)
	}
	g.e.EmitLine("constructor(init?: boolean | Buffer, pos?: number) ")
	if !g.SkipSerializer {
		g.e.EmitLine("pack(): Buffer;")
		g.e.EmitLine("unpack(buf: Buffer, pos?: number): void;")
		g.e.EmitLine("write(w: packer.ProtoWriter): void;")
		g.e.EmitLine("read(r: packer.ProtoReader): void;")
	}

	g.e.EndBracket("")

	if root {
		g.e.StartBracket("export namespace %s ", toJsClassCast(message.GetName()))
	} else {
		g.e.StartBracket("namespace %s ", toJsClassCast(message.GetName()))
	}

	for _, nm := range message.GetNestedMessageTypes() {
		if !nm.IsMapEntry() {
			g.emitClass(nm, false)
		}
	}
	for _, ne := range message.GetNestedEnumTypes() {
		g.emitEnum(ne, false)
	}
	g.e.EndBracket("")

	//g.e.EmitLine("export = " + toJsClassCast(message.GetName()) + ";")

}

func (g *TSGenerator) emitEnum(enum *desc.EnumDescriptor, root bool) {
	g.emitDescComment(enum)
	prefix := ""
	if root {
		prefix = "export "
	}
	g.e.Bracket("%senum %s", func() {
		vals := enum.GetValues()
		for i, v := range vals {
			g.emitDescComment(v)
			if i < len(vals)-1 {
				g.e.EmitLine("%s = %d,", g.toEnumName(v), v.GetNumber())
			} else {
				g.e.EmitLine("%s = %d", g.toEnumName(v), v.GetNumber())
			}
		}
	}, prefix, enum.GetName())
}

func (g *TSGenerator) emitDescComment(d desc.Descriptor) bool {
	return g.emitComment(d.GetSourceInfo().GetLeadingComments())
}

func (g *TSGenerator) emitComment(comment string) bool {
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
