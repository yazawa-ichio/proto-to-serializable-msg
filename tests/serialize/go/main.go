package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
	proto "github.com/yazawa-ichio/proto-to-serializable-msg/tests/serialize/go/proto"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("http listen :9001")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Print(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	model := r.URL.Path[1:]
	var m protopack.Message
	switch model {
	case "AllParameter":
		m = &proto.AllParameter{}
		break
	case "AllRepeatedParameter":
		m = &proto.AllRepeatedParameter{}
		break
	case "DependMessage":
		m = &proto.DependMessage{}
		break
	case "DependTest":
		m = &proto.DependTest{}
		break
	case "LowerCamelCase":
		m = &proto.LowerCamelCase{}
		break
	case "UpperCamelCase":
		m = &proto.UpperCamelCase{}
		break
	}
	err := m.Read(protopack.NewReader(r.Body))
	if err != nil && err != io.EOF {
		fmt.Printf("error %+v\n", err)
	}
	writer := protopack.NewWriter(bytes.Buffer{})
	if err = m.Write(writer); err != nil {
		fmt.Printf("error %+v\n", err)
	}
	w.Write(writer.Bytes())

}
