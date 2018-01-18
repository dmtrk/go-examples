package main

import (
	"fmt"
	"go-examples/pkg/util"
	"go-examples/pkg/rabbit"
	"go-examples/pkg/protobuf"
	"time"
)

func main() {
	fmt.Println("main()")
	rabbitTestText()
	rabbitTestProto()
	rabbitTestConsume()
}

func rabbitTestConsume() {
	fmt.Println("rabbitTestConsume()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["amqp.url"] = "amqp://172.16.0.125"
	properties["amqp.queue"] = "hello"
	//
	client := rabbit.NewRabbitClient(properties)
	defer client.Disconnect()
	fmt.Println("client: ", client)
	//
	err := client.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	} else {
		listener := new(Listener)
		err = client.Consume(listener)
		if err != nil {
			fmt.Println("Consume failed: ", err)
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}

func rabbitTestText() {
	fmt.Println("rabbitTestText()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["amqp.url"] = "amqp://172.16.0.125"
	properties["amqp.routing_key"] = "hello"
	//
	client := rabbit.NewRabbitClient(properties)
	defer client.Disconnect()
	fmt.Println("client: ", client)
	//
	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	data := "test data"
	//
	err := client.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	} else {
		err = client.Publish(headers, []byte(data))
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
	client := rabbit.NewRabbitClient(properties)
	defer client.Disconnect()
	fmt.Println("client: ", client)
	//
	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	data := "test proto data"
	message := protobuf.NewBytesMessage(headers, data)
	//
	err := client.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	} else {
		bytes, err2 := protobuf.SerializeBytesMessage(message)
		if err2 != nil {
			fmt.Println("SerializeBytesMessage failed: ", err)
		} else {
			err = client.Publish(headers, bytes)
			if err != nil {
				fmt.Println("Publish failed: ", err)
			}
		}
	}
}

type Listener struct {

}

func (listener *Listener) GetId() string {
	return "client1"
}

func (listener *Listener) OnMessage(headers map[string]string, data []byte) error {
	fmt.Println("OnMessage() data: ", string(data))

	return nil
}
