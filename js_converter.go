package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pseudomuto/protokit"
)

// JSConverter is Proto To JavaScript Language
type JSConverter struct {
}

// GetFileName is Proto To JavaScript Language
func (c *JSConverter) GetFileName(name string) string {
	return strings.ToLower(c.GetClassName(name) + ".js")
}

// GetPackageName is Proto To JavaScript Language
func (c *JSConverter) GetPackageName(name string) string {
	return strcase.ToCamel(name)
}

// GetClassName is Proto To JavaScript Language
func (c *JSConverter) GetClassName(name string) string {
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	names := strings.Split(name, ".")
	for i, n := range names {
		names[i] = strcase.ToCamel(n)
	}
	return strings.Join(names, ".")
}

// GetFieldName is Proto To JavaScript Language
func (c *JSConverter) GetFieldName(name string) string {
	return name
}

// GetEnumTypeName is Proto To JavaScript Language
func (c *JSConverter) GetEnumTypeName(name string) string {
	return name
}

// GetEnumName is Proto To JavaScript Language
func (c *JSConverter) GetEnumName(name string) string {
	return strings.ToUpper(name)
}

// GetType is Proto To JavaScript Language
func (c *JSConverter) GetType(f *protokit.FieldDescriptor) string {
	repeated := f.GetLabel().String() == "LABEL_REPEATED"
	if repeated {
		return c.GetTypeImpl(f) + "[]"
	}
	return c.GetTypeImpl(f)
}

// GetTypeImpl is Proto To JavaScript Language
func (c *JSConverter) GetTypeImpl(f *protokit.FieldDescriptor) string {
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
	return c.GetClassName(f.GetTypeName())
}
