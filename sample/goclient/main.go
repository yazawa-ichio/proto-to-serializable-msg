package main

import (
	"log"

	proto "github.com/yazawa-ichio/protoc-gen-msgpack/sample/goclient/proto"
)

func main() {
	//Pack
	data := &proto.Forum_User{
		ID:   32,
		Name: "test",
		Roll: proto.Forum_Roll_Master,
	}
	log.Println(data)
	buf, err := data.Pack()
	log.Println(buf, err)

	//Unpack
	dst := &proto.Forum_User{}
	err = dst.Unpack(buf)
	log.Println(dst, err)
	log.Println(&proto.DependTest{
		&proto.Forum_PostData{ID: 10},
	})
}
