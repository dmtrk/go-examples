package main

import (
	"fmt"
	"log"
	"go-examples/pkg/protobuf"
)

func main() {
	fmt.Println("main()")
	test1()
	test2()
	//test3()
}

func test1()  {
	fmt.Println("test1()")
	headers := map[string]string {"rsc": "3711", "r":   "2138"}
	dataStr := "some data"
	message1 := protobuf.NewBytesMessage(headers,dataStr)
	fmt.Println("message1: ",message1)
	// Serialize
	buf, err := protobuf.SerializeBytesMessage(message1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// Deserialize
	message2, err := protobuf.DeserializeBytesMessage(buf)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("message2: ",message2)
	// compare
	if string(message1.Data) != string(message2.Data) {
		log.Fatalf("data mismatch %q != %q", message1.Data, message2.Data)
	}
}

func test2()  {
	fmt.Println("test2()")
	headers := map[string]string {"h1": "v1", "h2":   "v2"}
	fields := map[string]string {"f1": "v11", "f2":   "v22"}
	message1 := protobuf.NewFieldsMessage(headers,fields)
	fmt.Println("message1: ",message1)
	// Serialize
	buf, err := protobuf.SerializeFieldsMessage(message1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// Deserialize
	message2, err := protobuf.DeserializeFieldsMessage(buf)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("message2: ",message2)
	// compare
	if len(message1.Headers) != len(message2.Headers) {
		log.Fatalf("data mismatch %q != %q", message1.Headers, message2.Headers)
	}
}
