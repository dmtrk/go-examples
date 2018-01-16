package main

import (
	"go-examples/pkg/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("main()")
	headers := map[string]string {"rsc": "3711", "r":   "2138"}
	//
	msg := protobuf.NewBytesMessage(headers,"some data")
	fmt.Println("msg:    ",msg)
	//
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	//
	newMsg := &protobuf.BytesMessage{}
	err = proto.Unmarshal(data, newMsg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("newMsg: ",newMsg)

	if string(msg.GetData()) != string(newMsg.GetData()) {
		log.Fatalf("data mismatch %q != %q", msg.GetData(), newMsg.GetData())
	}
}
