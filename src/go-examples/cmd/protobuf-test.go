package main

import (
	"fmt"
	"log"
	"go-examples/pkg/data"
)

func main() {
	fmt.Println("main()")
	test1()
	//test2()
	//test3()
}

func test1()  {
	fmt.Println("test1()")
	headers := map[string]string {"rsc": "3711", "r":   "2138"}
	dataStr := "some data"
	message1 := &data.BytesMessage{headers,dataStr}
	fmt.Println("message1: ",message1)
	// Serialize
	buf, err := message1.Serialize()
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// Deserialize
	message2, err := data.DeserializeBytesMessage(buf)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("message2: ",message2)
	// compare
	if string(message1.Data) != string(message2.Data) {
		log.Fatalf("data mismatch %q != %q", message1.Data, message2.Data)
	}
}

/*func test2()  {
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
}*/

/*func test3()  {
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
}*/
