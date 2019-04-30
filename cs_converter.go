package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pseudomuto/protokit"
)

// CSConverter is Proto To CSharp Language
type CSConverter struct {
}

// GetPackageName is Proto To CSharp Language
func (c *CSConverter) GetPackageName(name string) string {
	return strcase.ToCamel(name)
}

// GetFileName is Proto To CSharp Language
func (c *CSConverter) GetFileName(name string) string {
	return c.GetClassName(name) + ".cs"
}

// GetClassName is Proto To CSharp Language
func (c *CSConverter) GetClassName(name string) string {
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}
	names := strings.Split(name, ".")
	for i, n := range names {
		names[i] = strcase.ToCamel(n)
	}
	return strings.Join(names, ".")
}

// GetFieldName is Proto To CSharp Language
func (c *CSConverter) GetFieldName(name string) string {
	return strcase.ToCamel(name)
}

// GetEnumTypeName is Proto To CSharp Language
func (c *CSConverter) GetEnumTypeName(name string) string {
	return name
}

// GetEnumName is Proto To CSharp Language
func (c *CSConverter) GetEnumName(name string) string {
	return strcase.ToCamel(name)
}

// GetType is Proto To CSharp Language
func (c *CSConverter) GetType(f *protokit.FieldDescriptor) string {
	repeated := f.GetLabel().String() == "LABEL_REPEATED"
	if repeated {
		return c.GetTypeImpl(f) + "[]"
	}
	return c.GetTypeImpl(f)
}

// GetTypeImpl is Proto To CSharp Language
func (c *CSConverter) GetTypeImpl(f *protokit.FieldDescriptor) string {
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
	return c.GetClassName(f.GetTypeName())
}
