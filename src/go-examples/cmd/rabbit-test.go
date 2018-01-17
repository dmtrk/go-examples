package main

import (
	"fmt"
	"go-examples/pkg/util"
	"go-examples/pkg/rabbit"
	"go-examples/pkg/protobuf"
)

func main() {
	fmt.Println("main()")
	rabbitTestText()
	rabbitTestProto()
}

func rabbitTestText() {
	fmt.Println("rabbitTestText()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["amqp.url"] = "amqp://172.16.0.125"
	properties["amqp.routing_key"] = "hello"
	//
	publisher := rabbit.NewRabbitPublisher(properties)
	defer publisher.Disconnect()
	fmt.Println("publisher: ", publisher)
	//
	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	data := "test data"
	//
	err := publisher.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	} else {
		err = publisher.Publish(headers, []byte(data))
		if err != nil {
			fmt.Println("Publish failed: ", err)
		}
	}
}

func rabbitTestProto() {
	fmt.Println("rabbitTestProto()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["amqp.url"] = "amqp://172.16.0.125"
	properties["amqp.routing_key"] = "hello"
	//
	publisher := rabbit.NewRabbitPublisher(properties)
	defer publisher.Disconnect()
	fmt.Println("publisher: ", publisher)
	//
	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	data := "test proto data"
	message := protobuf.NewBytesMessage(headers, data)
	//
	err := publisher.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	} else {
		bytes, err2 := protobuf.SerializeBytesMessage(message)
		if err2 != nil {
			fmt.Println("SerializeBytesMessage failed: ", err)
		} else {
			err = publisher.Publish(headers, bytes)
			if err != nil {
				fmt.Println("Publish failed: ", err)
			}
		}
	}
}

