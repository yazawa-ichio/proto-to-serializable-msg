package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pseudomuto/protokit"
)

// TSConverter is Proto To TypeScript Language
type TSConverter struct {
	data *protoData
}

func newTSConverter(data *protoData) *TSConverter {
	return &TSConverter{data: data}
}

// GetFileName is Proto To TypeScript Language
func (c *TSConverter) GetFileName(name string) string {
	return strings.ToLower(c.GetClassName(name) + ".d.ts")
}

// GetPackageName is Proto To TypeScript Language
func (c *TSConverter) GetPackageName(name string) string {
	return strcase.ToCamel(name)
}

// GetClassName is Proto To TypeScript Language
func (c *TSConverter) GetClassName(name string) string {
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	names := strings.Split(name, ".")
	for i, n := range names {
		names[i] = strcase.ToCamel(n)
	}
	return strings.Join(names, ".")
}

// GetFieldName is Proto To TypeScript Language
func (c *TSConverter) GetFieldName(name string) string {
	return name
}

// GetEnumTypeName is Proto To TypeScript Language
func (c *TSConverter) GetEnumTypeName(name string) string {
	return name
}

// GetEnumName is Proto To TypeScript Language
func (c *TSConverter) GetEnumName(name string) string {
	return strings.ToUpper(name)
}

// GetType is Proto To TypeScript Language
func (c *TSConverter) GetType(f *protokit.FieldDescriptor) string {
	mapEntry := c.data.isMapEntry(f)
	if mapEntry {
		key, val := c.data.getMapKeyValue(f)
		return "Map<" + c.GetType(key) + ", " + c.GetType(val) + ">"
	}
	repeated := f.GetLabel().String() == "LABEL_REPEATED"
	if repeated {
		return "Array<" + c.GetTypeImpl(f) + ">"
	}
	return c.GetTypeImpl(f)
}

// GetTypeImpl is Proto To TypeScript Language
func (c *TSConverter) GetTypeImpl(f *protokit.FieldDescriptor) string {
	switch f.GetType().String() {
	case "TYPE_INT64", "TYPE_UINT64", "TYPE_INT32", "TYPE_FIXED64", "TYPE_FIXED32", "TYPE_UINT32", "TYPE_SFIXED32", "TYPE_SFIXED64", "TYPE_SINT32", "TYPE_SINT64", "TYPE_DOUBLE", "TYPE_FLOAT":
		return "number"
	case "TYPE_BOOL":
		return "bool"
	case "TYPE_STRING":
		return "string"
	case "TYPE_BYTES":
		return "Uint8Array"
	}
	typeName := c.GetClassName(f.GetTypeName())
	if c.data.isUserDefine(f) {
		typeName = c.importName(typeName)
	}
	if f.GetType().String() == "TYPE_MESSAGE" || f.GetType().String() == "TYPE_BYTES" {
		typeName = typeName + " | null"
	}
	return typeName
}

func (c *TSConverter) formFileName(name string) string {
	name = c.GetFileName(name)
	name = name[:len(name)-5]
	return name
}

func (c *TSConverter) importName(name string) string {
	return strings.Replace(c.formFileName(name), ".", "_", -1)
}
