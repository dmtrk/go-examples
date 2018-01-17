package main

import (
	"fmt"
	"go-examples/pkg/http"
	"go-examples/pkg/util"
)

func main() {
	fmt.Println("main()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	publisher := http.NewHttpPublisher(properties)
	fmt.Sprintf("publisher: %v", publisher)

	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	publisher.Post(headers, "test data")

}
