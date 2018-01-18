package main

import (
	"fmt"
	"go-examples/pkg/util"
	"go-examples/pkg/kafka"
	"go-examples/pkg/protobuf"
)

func main() {
	fmt.Println("main()")
	kafkaProducerTest1()
	kafkaProducerTest2()
	kafkaConsumerTest()
}

func kafkaProducerTest1() {
	fmt.Println("kafkaProducerTest1()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["kafka.brokers"] = "172.16.0.125:9092"
	properties["kafka.topic"] = "test"
	//
	publisher := kafka.NewKafkaProducer(properties)
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

func kafkaProducerTest2() {
	fmt.Println("kafkaProducerTest2()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["kafka.brokers"] = "172.16.0.125:9092"
	properties["kafka.topic"] = "test"
	//
	publisher := kafka.NewKafkaProducer(properties)
	defer publisher.Disconnect()
	fmt.Println("publisher: ", publisher)
	//
	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	data := "test data"
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

func kafkaConsumerTest() {
	fmt.Println("kafkaConsumerTest()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["kafka.brokers"] = "172.16.0.125:9092"
	properties["kafka.topic"] = "test"
	//
	consumer := kafka.NewKafkaConsumer(properties)
	defer consumer.Disconnect()
	fmt.Println("consumer: ", consumer)
	//
}