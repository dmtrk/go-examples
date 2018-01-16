package main

import (
	"go-examples/pkg/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("main()")
	test1()
	test2()
	test3()
}

func test1()  {
	fmt.Println("test1()")
	headers := map[string]string {"rsc": "3711", "r":   "2138"}
	data := "some data"
	msg := protobuf.NewBytesMessage(headers,data)
	fmt.Println("msg:    ",msg)
	//
	buf, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	//
	newMsg := &protobuf.BytesMessage{}
	err = proto.Unmarshal(buf, newMsg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("newMsg: ",newMsg)

	if string(msg.GetData()) != string(newMsg.GetData()) {
		log.Fatalf("data mismatch %q != %q", msg.GetData(), newMsg.GetData())
	}
}

func test2()  {
	fmt.Println("test2()")
	headers := map[string]string {"rsc": "3711", "r":   "2138"}
	data := "some data"
	//
	msg := protobuf.NewBytesMessage(headers,data)
	buf, err := protobuf.BytesMessageToBytes(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newMsg, err := protobuf.BytesMessageFromBytes(buf)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("newMsg: ",newMsg)
}
func test3()  {
	fmt.Println("test3()")
	headers := map[string]string {"h1": "v1", "h2":   "v2"}
	fields := map[string]string {"f1": "v11", "f2":   "v22"}
	//
	msg := protobuf.NewFieldsMessage(headers,fields)
	buf, err := protobuf.FieldsMessageToBytes(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newMsg, err := protobuf.FieldsMessageFromBytes(buf)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("newMsg: ",newMsg)
}
