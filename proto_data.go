package main

import (
	"github.com/pseudomuto/protokit"
)

type protoData struct {
	messages map[string]*messageData
	enums    map[string]*enumData
}

type messageData struct {
	file     *protokit.FileDescriptor
	data     *protokit.Descriptor
	parent   *protokit.Descriptor
	children []*messageData
	enums    []*enumData
	mapEntry bool
}

type enumData struct {
	file   *protokit.FileDescriptor
	data   *protokit.EnumDescriptor
	parent *protokit.Descriptor
}

func newProtoData(fds []*protokit.FileDescriptor) *protoData {
	data := &protoData{}
	data.messages = map[string]*messageData{}
	data.enums = map[string]*enumData{}
	for _, f := range fds {
		for _, m := range f.Messages {
			setMessage(data, f, m, nil)
		}
		for _, e := range f.Enums {
			setEnum(data, f, e, nil)
		}
	}
	return data
}

func setMessage(p *protoData, file *protokit.FileDescriptor, data *protokit.Descriptor, parent *protokit.Descriptor) *messageData {
	message := &messageData{file: file, data: data, parent: parent}
	message.mapEntry = data.GetOptions().GetMapEntry()
	message.children = make([]*messageData, len(data.Messages))
	message.enums = make([]*enumData, len(data.Enums))
	fullName := data.GetFullName()
	if fullName[0] == '.' {
		fullName = fullName[1:]
	}
	p.messages[fullName] = message
	for i, m := range data.Messages {
		message.children[i] = setMessage(p, file, m, data)
	}
	for i, e := range data.Enums {
		message.enums[i] = setEnum(p, file, e, data)
	}
	return message
}

func setEnum(p *protoData, file *protokit.FileDescriptor, data *protokit.EnumDescriptor, parent *protokit.Descriptor) *enumData {
	enum := &enumData{file: file, data: data, parent: parent}
	p.enums[data.GetFullName()] = enum
	return enum
}

func (d *protoData) isUserDefine(f *protokit.FieldDescriptor) bool {
	_type := f.GetType().String()
	if _type == "TYPE_MESSAGE" {
		if d.getMessageData(f).mapEntry {
			return false
		}
		return true
	} else if _type == "TYPE_ENUM" {
		return true
	}
	return false
}

func (d *protoData) getMessageData(f *protokit.FieldDescriptor) *messageData {
	if f.GetType().String() == "TYPE_MESSAGE" {
		typeName := f.GetTypeName()
		if typeName[0] == '.' {
			typeName = typeName[1:]
		}
		m := d.messages[typeName]
		return m
	}
	return nil
}

func (d *protoData) isMapEntry(f *protokit.FieldDescriptor) bool {
	if f.GetType().String() == "TYPE_MESSAGE" {
		m := d.getMessageData(f)
		if m != nil {
			return m.mapEntry
		}
	}
	return false
}

func (d *protoData) getMapKeyValue(f *protokit.FieldDescriptor) (*protokit.FieldDescriptor, *protokit.FieldDescriptor) {
	m := d.getMessageData(f)
	var key *protokit.FieldDescriptor
	var value *protokit.FieldDescriptor
	for _, field := range m.data.GetMessageFields() {
		if field.GetName() == "key" {
			key = field
		}
		if field.GetName() == "value" {
			value = field
		}
	}
	return key, value
}
