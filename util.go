package proto

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	ce "github.com/yazawa-ichio/proto-to-serializable-msg/code_emitter"
)

func toCamel(s string, disableFirstToUpper bool) string {
	s = strings.Trim(s, " ")
	ret := ""
	changeUpper := !disableFirstToUpper
	prevUppercase := false
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if prevUppercase || (i == 0 && disableFirstToUpper) {
				ret += strings.ToLower(string(c))
			} else {
				ret += string(c)
			}
			prevUppercase = true
			changeUpper = false
			continue
		}
		prevUppercase = false
		if c >= 'a' && c <= 'z' {
			if changeUpper {
				ret += strings.ToUpper(string(c))
			} else {
				ret += strings.ToLower(string(c))
			}
			changeUpper = false
			continue
		}
		skip := (c == '_' || c == ' ' || c == '-')
		if !skip {
			ret += string(c)
		}
		changeUpper = (skip || c == '.' || (c >= '0' && c <= '9'))
	}
	return ret
}

func emitFileInfo(e *ce.CodeEmitter, d desc.Descriptor) {
	e.EmitLine("// %s", strings.Replace(d.GetFile().GetName(), "\\", "/", -1))
}

func isUserDefine(f *desc.FieldDescriptor) bool {
	_type := f.GetType().String()
	if _type == "TYPE_MESSAGE" {
		if f.IsMap() {
			return false
		}
		return true
	} else if _type == "TYPE_ENUM" {
		return true
	}
	return false
}

func parseFiles(files []string) ([]*desc.FileDescriptor, error) {
	sort.Strings(files)
	p := &protoparse.Parser{}
	p.IncludeSourceCodeInfo = true
	return p.ParseFiles(files...)
}

func FindProtoFiles(path string) ([]string, error) {
	f, err := os.Stat(path)
	if err == nil {
		if f.IsDir() {
			files := make([]string, 0)
			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return err
				}
				if filepath.Ext(path) == ".proto" {
					files = append(files, filepath.ToSlash(path))
				}
				return nil
			})
			return files, err
		} else {
			return []string{path}, nil
		}
	}
	return nil, fmt.Errorf("input not found %s, use filepath or dirpath", path)
}
