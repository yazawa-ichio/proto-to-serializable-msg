package proto

import (
	"strconv"
	"strings"

	"github.com/jhump/protoreflect/desc"
	ce "github.com/yazawa-ichio/proto-to-serializable-msg/code_emitter"
)

type GOGenerator struct {
	e              *ce.CodeEmitter
	PackageRoot    string
	SkipSerializer bool
	typeMap        map[string]*golangTypeInfo
}

type golangTypeInfo struct {
	packageDefine       bool
	goPackageName       string
	samePackageTypeName string
	dependTypeName      string
}

func NewGoGenerator() *GOGenerator {
	return &GOGenerator{
		e: new(ce.CodeEmitter),
	}
}

func (g *GOGenerator) Generate(files []string) ([]*GenerateFile, error) {
	output := make([]*GenerateFile, 0)
	parsed, err := parseFiles(files)
	if err != nil {
		return nil, err
	}
	// 型情報を収集
	g.createTypeMap(parsed)
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

func (g *GOGenerator) GenerateAndOutput(files []string, outputRoot string) error {
	return output(g, files, outputRoot)
}

func (g *GOGenerator) createTypeMap(files []*desc.FileDescriptor) {
	ret := make(map[string]*golangTypeInfo, 0)
	for _, f := range files {
		for _, m := range f.GetMessageTypes() {
			g.createMessageTypeMap(ret, m)
		}
		for _, e := range f.GetEnumTypes() {
			ret[e.GetFullyQualifiedName()] = g.createTypeInfo(e)
		}
	}
	g.typeMap = ret
}

func (g *GOGenerator) createMessageTypeMap(ret map[string]*golangTypeInfo, m *desc.MessageDescriptor) {
	ret[m.GetFullyQualifiedName()] = g.createTypeInfo(m)
	for _, nm := range m.GetNestedMessageTypes() {
		if !nm.IsMapEntry() {
			g.createMessageTypeMap(ret, nm)
		}
	}
}

func (g *GOGenerator) createTypeInfo(d desc.Descriptor) *golangTypeInfo {
	packageName := d.GetFile().GetPackage()
	packageDefine := packageName != ""
	goPackageName := g.getGoPackageName(d)
	name := d.GetFullyQualifiedName()
	if packageDefine && strings.HasPrefix(name, packageName) {
		name = name[len(packageName)+1:]
	}
	names := strings.Split(name, ".")
	for i, n := range names {
		names[i] = toCamel(n, false)
	}
	name = strings.Join(names, "_")
	samePackageTypeName := name
	dependTypeName := goPackageName + "." + name
	return &golangTypeInfo{
		packageDefine,
		goPackageName,
		samePackageTypeName,
		dependTypeName,
	}
}

func (c *GOGenerator) getGoPackageName(d desc.Descriptor) string {
	name := d.GetFile().GetPackage()
	if name == "" {
		name = "proto"
	}
	return strings.ToLower(name)
}

func (g *GOGenerator) getGoFileName(d desc.Descriptor) string {
	definePackageName := d.GetFile().GetPackage() != ""
	packageName := g.getGoPackageName(d)
	if g.PackageRoot == "" {
		return strings.ToLower(d.GetFullyQualifiedName() + ".go")
	} else {
		name := strings.ToLower(d.GetFullyQualifiedName())
		if definePackageName {
			if strings.HasPrefix(name, packageName) {
				name = name[len(packageName)+1:]
			}
		}
		if definePackageName {
			return strings.ToLower(packageName + "/" + name + ".go")
		} else {
			return strings.ToLower(name + ".go")
		}
	}
}

func (g *GOGenerator) toTypeName(fullName string, d desc.Descriptor) string {
	if strings.HasPrefix(fullName, ".") {
		fullName = fullName[1:]
	}
	if g.PackageRoot != "" {
		if g.typeMap[fullName].goPackageName == g.getGoPackageName(d) {
			return g.typeMap[fullName].samePackageTypeName
		} else {
			return g.typeMap[fullName].dependTypeName
		}
	}
	packageName := d.GetFile().GetPackage()
	names := strings.Split(fullName, ".")
	if g.PackageRoot != "" && names[0] == packageName {
		names = names[1:]
	}
	for i, n := range names {
		names[i] = toCamel(n, false)
	}
	return strings.Join(names, "_")
}

func (g *GOGenerator) getClassName(m *desc.MessageDescriptor) string {
	return g.toTypeName(m.GetFullyQualifiedName(), m)
}

func (g *GOGenerator) getEnumName(m *desc.EnumDescriptor) string {
	return g.toTypeName(m.GetFullyQualifiedName(), m)
}

func (g *GOGenerator) getFieldName(f *desc.FieldDescriptor) string {
	name := toCamel(f.GetName(), false)
	for _, m := range []string{"Id", "Rpc"} {
		name = strings.Replace(name, m, strings.ToUpper(m), -1)
	}
	return name
}

func (g *GOGenerator) toEnumName(d *desc.EnumValueDescriptor) string {
	return toCamel(d.GetName(), false)
}

func (g *GOGenerator) genClass(message *desc.MessageDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, message)
	g.e.NewLine()
	if g.PackageRoot == "" {
		g.e.EmitLine("package proto")
	} else {
		g.e.EmitLine("package %s", g.getGoPackageName(message))
	}
	g.e.NewLine()
	g.emitImport(message)
	g.emitClass(message)
	return &GenerateFile{
		Name:    g.getGoFileName(message),
		Content: g.e.String(),
	}
}

func (g *GOGenerator) genEnum(enum *desc.EnumDescriptor) *GenerateFile {
	g.e.Reset()
	emitFileInfo(g.e, enum)
	g.e.NewLine()
	if g.PackageRoot == "" {
		g.e.EmitLine("package proto")
	} else {
		g.e.EmitLine("package %s", g.getGoPackageName(enum))
	}
	g.e.EmitLine("")
	g.emitEnum(enum)
	return &GenerateFile{
		Name:    g.getGoFileName(enum),
		Content: g.e.String(),
	}
}

func (g *GOGenerator) emitImport(m *desc.MessageDescriptor) {
	hits := make(map[string]string)
	if g.PackageRoot != "" {
		g.getDepPackage(hits, m)
	}
	g.e.EmitLine("import (")
	g.e.StartIndent()
	g.e.EmitLine("protopack \"github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang\"")
	for _, value := range hits {
		g.e.EmitLine(value)
	}
	g.e.EndIndent()
	g.e.EmitLine(")")
	g.e.NewLine()
}

func (g *GOGenerator) getDepPackage(hits map[string]string, m *desc.MessageDescriptor) {
	for _, f := range m.GetFields() {
		if f.IsMap() {
			g.setDepPackage(hits, f.GetMapKeyType(), m)
			g.setDepPackage(hits, f.GetMapValueType(), m)
		} else {
			g.setDepPackage(hits, f, m)
		}
	}
	for _, nm := range m.GetNestedMessageTypes() {
		g.getDepPackage(hits, nm)
	}
}

func (g *GOGenerator) setDepPackage(hits map[string]string, f *desc.FieldDescriptor, m *desc.MessageDescriptor) {
	typeName := f.AsFieldDescriptorProto().GetTypeName()
	if typeName == "" {
		return
	}
	if strings.HasPrefix(typeName, ".") {
		typeName = typeName[1:]
	}
	if info, ok := g.typeMap[typeName]; ok {
		if g.getGoPackageName(m) == info.goPackageName {
			return
		}
		if _, ok = hits[info.goPackageName]; ok {
			return
		}
		if info.packageDefine {
			hits[info.goPackageName] = info.goPackageName + " \"" + g.PackageRoot + "/" + info.goPackageName + "\""
		} else {
			hits[info.goPackageName] = info.goPackageName + " \"" + g.PackageRoot + "\""
		}
	}
}

func (g *GOGenerator) emitComment(comment string) {
	if comment == "" {
		return
	}
	if comment[len(comment)-1] == '\n' {
		comment = comment[:len(comment)-1]
	}
	for _, c := range strings.Split(comment, "\n") {
		g.e.EmitLine("//%s" + c)
	}
	//g.e.EmitLine("//%s", comment)
	return
}

func (g *GOGenerator) emitParamComment(param, comment string) {
	if comment == "" {
		return
	}
	if comment[len(comment)-1] == '\n' {
		comment = comment[:len(comment)-1]
	}
	for i, c := range strings.Split(comment, "\n") {
		if i == 0 {
			g.e.EmitLine("// %s %s", param, c)
		} else {
			g.e.EmitLine("//%s", c)
		}
	}
	return
}

func (g *GOGenerator) emitCommentField(field *desc.FieldDescriptor) {
	comment := field.GetSourceInfo().GetLeadingComments()
	if comment == "" {
		return
	}
	if comment[len(comment)-1] == '\n' {
		comment = comment[:len(comment)-1]
	}
	g.e.EmitLine("// %s %s", field.GetName(), comment)
	return
}

func (g *GOGenerator) emitClass(message *desc.MessageDescriptor) {

	typeName := g.getClassName(message)
	//emit Class
	g.emitParamComment(typeName, message.GetSourceInfo().GetLeadingComments())
	g.e.StartBracket("type %s struct ", typeName)
	maxFieldLen := 0
	for _, f := range message.GetFields() {
		len := len(g.getFieldName(f))
		if len > maxFieldLen {
			maxFieldLen = len
		}
	}
	for _, f := range message.GetFields() {
		fieldType := g.getType(f, message, "*")
		fieldName := g.getFieldName(f)
		g.emitCommentField(f)
		for len(fieldName) < maxFieldLen {
			fieldName += " "
		}
		g.e.EmitLine("%s %s", fieldName, fieldType)
	}
	g.e.EndBracket("")

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
		g.emitWriter(message)
		g.emitReader(message)
	}

}

func (g *GOGenerator) emitEnum(enum *desc.EnumDescriptor) {

	typeName := g.getEnumName(enum)

	//emit Enum
	g.emitParamComment(typeName, enum.GetSourceInfo().GetLeadingComments())
	g.e.EmitLine("type %s int32", typeName)
	g.e.EmitLine("")

	g.e.EmitLine("const (")
	g.e.StartIndent()
	for _, f := range enum.GetValues() {
		g.emitParamComment(typeName+"_"+g.toEnumName(f), f.GetSourceInfo().GetLeadingComments())
		g.e.EmitLine("%s_%s %s = %s", typeName, g.toEnumName(f), typeName, strconv.Itoa(int(f.GetNumber())))
	}
	g.e.EndIndent()
	g.e.EmitLine(")")
	g.e.NewLine()

	g.emitParamComment("String", typeName+" to string")
	g.e.StartBracket("func (x %s) String() string ", typeName)
	g.e.EmitLine("switch x {")
	for _, f := range enum.GetValues() {
		g.e.EmitLine("case %s_%s:", typeName, g.toEnumName(f))
		g.e.EmitIndentLine("return \"%s\"", g.toEnumName(f))
	}
	g.e.EmitLine("default:")
	g.e.EmitIndentLine("return \"Unknown\"")
	g.e.EmitLine("}")
	g.e.EndBracket("")

	g.e.NewLine()
	g.e.EmitLine("// Parse%s string to %s", typeName, typeName)
	g.e.StartBracket("func Parse%s(val string) (%s , bool) ", typeName, typeName)
	g.e.EmitLine("switch val {")
	for _, f := range enum.GetValues() {
		g.e.EmitLine("case \"%s\":", g.toEnumName(f))
		g.e.EmitIndentLine("return %s_%s, true", typeName, g.toEnumName(f))
	}
	g.e.EmitLine("default:")
	g.e.EmitIndentLine("return %s(0), false", typeName)
	g.e.EmitLine("}")
	g.e.EndBracket("")

}
func (g *GOGenerator) emitErrorCheck() {
	g.e.StartBracket("if err != nil ")
	g.e.EmitLine("return err")
	g.e.EndBracket("")
}

func (g *GOGenerator) emitWriter(message *desc.MessageDescriptor) {
	g.e.NewLine()

	typeName := g.getClassName(message)

	g.emitParamComment("Pack", "Serialize Message")
	g.e.StartBracket("func (m *%s) Pack() ([]byte, error) ", typeName)
	g.e.EmitLine("return protopack.Pack(m)")
	g.e.EndBracket("")

	g.e.NewLine()
	g.emitParamComment("Write", "Serialize Message")
	g.e.StartBracket("func (m *%s) Write(w protopack.Writer) error ", typeName)
	defer g.e.EndBracket("")
	defer g.e.EmitLine("return err")

	g.e.EmitLine("// Write Map Length")
	g.e.EmitLine("err := w.WriteMapHeader(%s)", strconv.Itoa(len(message.GetFields())))
	g.emitErrorCheck()

	for _, f := range message.GetFields() {
		g.e.EmitLine("")
		g.e.EmitLine("// Write " + f.GetName())
		filedName := g.getFieldName(f)

		g.e.EmitLine("err = w.WriteTag(%s)", strconv.Itoa(int(f.GetNumber())))
		g.emitErrorCheck()

		if f.IsMap() {
			g.e.StartBracket("if m.%s == nil ", filedName)
			{
				g.e.EmitLine("err = w.WriteNil()")
				g.emitErrorCheck()
			}
			g.e.EndAndStartBracket(" else ")
			{
				g.e.EmitLine("mapLen := len(m.%s)", filedName)
				g.e.EmitLine("err = w.WriteMapHeader(mapLen)")
				g.emitErrorCheck()

				mapKey, mapVal := f.GetMapKeyType(), f.GetMapValueType()
				g.e.StartBracket("for map%sKey, map%sValue := range m.%s ", filedName, filedName, filedName)
				g.emitSerialize(mapKey, "map"+filedName, "")
				g.emitSerialize(mapVal, "map"+filedName, "")
				g.e.EndBracket("")
			}
			g.e.EndBracket("")

		} else if f.GetLabel().String() == "LABEL_REPEATED" {
			g.e.StartBracket("if m.%s == nil", filedName)
			{
				g.e.EmitLine("err = w.WriteNil()")
				g.emitErrorCheck()
			}
			g.e.EndAndStartBracket(" else ")
			g.e.EmitLine("arrayLen := len(m.%s)", filedName)
			g.e.EmitLine("err = w.WriteArrayHeader(arrayLen)")
			g.emitErrorCheck()
			g.e.StartBracket("for arrayIndex, _ := range m.%s ", filedName)
			{
				g.emitSerialize(f, "m.", "[arrayIndex]")
			}
			g.e.EndBracket("")
			g.e.EndBracket("")
		} else {
			g.emitSerialize(f, "m.", "")
		}
	}

}

func (g *GOGenerator) emitSerialize(f *desc.FieldDescriptor, prefix string, suffix string) {
	fieldName := g.getFieldName(f)
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		g.e.EmitLine("err = w.WriteMessage(%s)", prefix+fieldName+suffix)
		break
	case "TYPE_ENUM":
		g.e.EmitLine("err = w.WriteInt32(int32(%s))", prefix+fieldName+suffix)
		break
	case "TYPE_BYTES":
		g.e.EmitLine("err = w.WriteBytes(%s)", prefix+fieldName+suffix)
		break
	default:
		typeName := strings.Title(g.getTypeImpl(f, f.GetOwner(), ""))
		g.e.EmitLine("err = w.Write%s(%s)", typeName, prefix+fieldName+suffix)
		break
	}
	g.emitErrorCheck()
}

func (g *GOGenerator) emitReader(message *desc.MessageDescriptor) {

	typeName := g.getClassName(message)

	g.e.NewLine()
	g.emitParamComment("Unpack", "Serialize Message")
	g.e.StartBracket("func (m *%s) Unpack(buf []byte) (error) ", typeName)
	g.e.EmitLine("return protopack.Unpack(m, buf)")
	g.e.EndBracket("")

	g.e.NewLine()
	g.emitParamComment("Read", "Deserialize Message")
	g.e.StartBracket("func (m *%s) Read(r protopack.Reader) error ", typeName)
	defer g.e.EndBracket("")
	defer g.e.EmitLine("return err")

	g.e.EmitLine("// Read Map Length")
	g.e.EmitLine("len, err := r.ReadMapHeader()")
	g.emitErrorCheck()
	g.e.NewLine()
	g.e.StartBracket("for i := uint(0); i < len; i++ ")
	defer g.e.EndBracket("")

	g.e.NewLine()
	g.e.EmitLine("// Read Tag")
	g.e.EmitLine("tag, err := r.ReadTag()")
	g.emitErrorCheck()
	g.e.NewLine()

	g.e.EmitLine("switch tag {")
	for _, f := range message.GetFields() {
		g.e.EmitLine("case %d: // Read %s", f.GetNumber(), f.GetName())
		g.e.StartIndent()
		if f.IsMap() {
			g.emitMapDeserialize(f)
		} else if f.GetLabel().String() == "LABEL_REPEATED" {
			g.emitRepeatedDeserialize(f)
		} else {
			g.emitDeserialize(f, "m.", "")
		}
		g.e.EmitLine("break")
		g.e.EndIndent()
	}
	{
		g.e.EmitLine("default:")
		g.e.StartIndent()
		g.e.EmitLine("err = r.Skip()")
		g.emitErrorCheck()
		g.e.EmitLine("break")
	}
	g.e.EndIndent()
	g.e.EmitLine("}")
}

func (g *GOGenerator) emitMapDeserialize(f *desc.FieldDescriptor) {
	filedName := g.getFieldName(f)
	key, value := f.GetMapKeyType(), f.GetMapValueType()
	keyType := g.getType(key, f.GetOwner(), "*")
	valueType := g.getType(value, f.GetOwner(), "*")
	g.e.EmitLine("isMapNil, err := r.NextFormatIsNull() ")
	g.emitErrorCheck()
	g.e.Bracket("if isMapNil ", func() {
		g.e.EmitLine("r.ReadNil()")
		g.e.EmitLine("m.%s = nil", filedName)
		g.e.EmitLine("continue")
	})

	g.e.EmitLine("map%sLen, err := r.ReadMapHeader();", filedName)
	g.emitErrorCheck()

	g.e.EmitLine("m.%s = make(map[%s]%s, map%sLen)", filedName, keyType, valueType, filedName)
	g.e.StartBracket("for mapIndex := uint(0); mapIndex < map%sLen; mapIndex++ ", filedName)
	{
		g.e.EmitLine("var map%sKey %s", filedName, keyType)
		g.e.EmitLine("var map%sValue %s", filedName, valueType)
		g.emitDeserialize(key, "map"+filedName, "")
		g.emitDeserialize(value, "map"+filedName, "")
		g.e.EmitLine("m.%s[map%sKey] = map%sValue;", filedName, filedName, filedName)
	}
	g.e.EndBracket("")
}

func (g *GOGenerator) emitRepeatedDeserialize(f *desc.FieldDescriptor) {
	filedName := g.getFieldName(f)
	g.e.EmitLine("isArrayNil, err := r.NextFormatIsNull() ")
	g.emitErrorCheck()
	g.e.Bracket("if isArrayNil ", func() {
		g.e.EmitLine("r.ReadNil()")
		g.e.EmitLine("m.%s = nil", filedName)
		g.e.EmitLine("continue")
	})
	g.e.EmitLine("_%sLen, err := r.ReadArrayHeader();", filedName)
	g.emitErrorCheck()
	if g.getTypeImpl(f, f.GetOwner(), "*") != "byte[]" {
		g.e.EmitLine("m.%s = make([]%s, _%sLen);", filedName, g.getTypeImpl(f, f.GetOwner(), "*"), filedName)
	} else {
		g.e.EmitLine("m.%s = new byte[_%sLen][];", filedName, filedName)
	}
	g.e.StartBracket("for arrayIndex := uint(0); arrayIndex < _%sLen; arrayIndex++ ", filedName)
	g.emitDeserialize(f, "m.", "[arrayIndex]")
	g.e.EndBracket("")
}

func (g *GOGenerator) emitDeserialize(f *desc.FieldDescriptor, prefix string, suffix string) {
	filedName := g.getFieldName(f)
	switch f.GetType().String() {
	case "TYPE_MESSAGE":
		g.e.EmitLine("isNil, err := r.NextFormatIsNull()")
		g.emitErrorCheck()
		g.e.StartBracket("if isNil ")
		{
			g.e.EmitLine("%s = nil", prefix+filedName+suffix)
			g.e.EmitLine("err = r.ReadNil()")
		}
		g.e.EndAndStartBracket(" else ")
		{
			g.e.EmitLine("%s = %s{}", prefix+filedName+suffix, g.getTypeImpl(f, f.GetOwner(), "&"))
			g.e.EmitLine("err = %s.Read(r)", prefix+filedName+suffix)
		}
		g.e.EndBracket("")
		break
	case "TYPE_ENUM":
		g.e.EmitLine("val, err := r.ReadInt32()")
		g.e.EmitLine("%s = %s(val)", prefix+filedName+suffix, g.getTypeImpl(f, f.GetOwner(), "*"))
		break
	case "TYPE_BYTES":
		g.e.EmitLine("%s, err = r.ReadBytes();", prefix+filedName+suffix)
		break
	default:
		g.e.EmitLine("%s, err = r.Read%s();", prefix+filedName+suffix, strings.Title(g.getTypeImpl(f, f.GetOwner(), "*")))
		break
	}
	g.emitErrorCheck()
}

// GetType is Proto To Go Language
func (g *GOGenerator) getType(f *desc.FieldDescriptor, owner *desc.MessageDescriptor, classPrefix string) string {
	//packageName := g.getPackageName(owner)
	if f.IsMap() {
		key, val := f.GetMapKeyType(), f.GetMapValueType()
		return "map[" + g.getType(key, owner, "") + "]" + g.getType(val, owner, "*")
	}
	repeated := f.GetLabel().String() == "LABEL_REPEATED"
	if repeated {
		return "[]" + g.getTypeImpl(f, owner, "*")
	}
	return g.getTypeImpl(f, owner, classPrefix)
}

func (g *GOGenerator) getTypeImpl(f *desc.FieldDescriptor, owner *desc.MessageDescriptor, classPrefix string) string {
	switch f.GetType().String() {
	case "TYPE_INT64":
		return "int64"
	case "TYPE_UINT64":
		return "uint64"
	case "TYPE_INT32":
		return "int32"
	case "TYPE_UINT32":
		return "uint32"
	case "TYPE_FIXED64":
		return "uint64"
	case "TYPE_FIXED32":
		return "uint32"
	case "TYPE_SFIXED32":
		return "int32"
	case "TYPE_SFIXED64":
		return "int64"
	case "TYPE_SINT32":
		return "int32"
	case "TYPE_SINT64":
		return "int64"
	case "TYPE_DOUBLE":
		return "float64"
	case "TYPE_FLOAT":
		return "float32"
	case "TYPE_BOOL":
		return "bool"
	case "TYPE_STRING":
		return "string"
	case "TYPE_BYTES":
		return "[]byte"
	case "TYPE_ENUM":
		name := f.AsFieldDescriptorProto().GetTypeName()
		if strings.HasPrefix(name, ".") {
			name = name[1:]
		}
		return g.toTypeName(name, owner)
	}
	name := f.AsFieldDescriptorProto().GetTypeName()
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	return classPrefix + g.toTypeName(name, owner)
}
