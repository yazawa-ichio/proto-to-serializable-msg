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
	message.children = []*messageData{}
	message.enums = []*enumData{}
	p.messages[data.GetFullName()] = message
	for _, m := range data.Messages {
		message.children = append(message.children, setMessage(p, file, m, data))
	}
	for _, e := range data.Enums {
		message.enums = append(message.enums, setEnum(p, file, e, data))
	}
	return message
}

func setEnum(p *protoData, file *protokit.FileDescriptor, data *protokit.EnumDescriptor, parent *protokit.Descriptor) *enumData {
	enum := &enumData{file: file, data: data, parent: parent}
	p.enums[data.GetFullName()] = enum
	return enum
}
