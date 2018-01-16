package main

import (
	"fmt"
	"go-examples/pkg/http"
)


func main() {
	fmt.Println("main()")

	//publisher := new (HttpPublisher)
	publisher := http.HttpPublisher{}
	fmt.Sprintf("publisher: %v",publisher)


	headers := map[string]string {
		"rsc": "3711",
		//"r":   "2138",
		//"gri": "1908",
		//"adg": "912",
	}
	publisher.Publish(headers, "test data")


}
