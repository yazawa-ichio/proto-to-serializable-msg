package main

import (
	"github.com/pseudomuto/protokit"
)

//ProtoConverter is Convert Proto To Language
type ProtoConverter interface {
	GetPackageName(name string) string
	GetClassName(name string) string
	GetFieldName(name string) string
	GetEnumTypeName(name string) string
	GetEnumName(name string) string
	GetType(name *protokit.FieldDescriptor) string
}
