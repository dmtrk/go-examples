package main

import (
	"go-examples/pkg/util"
	"strings"
	"fmt"
)

func main() {
	fmt.Println("main()")
	//
	properties := util.ParseProperties(strings.NewReader("key1=val1 \n # \n key2=val2"))
	fmt.Println("properties: ", properties)
}
