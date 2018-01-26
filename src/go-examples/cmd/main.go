package main

import (
	"go-examples/pkg/util"
	"strings"
	"go-examples/pkg/docker"
	"os"
	"log"
)

func main() {
	log.Print("main()")
	//
	properties := util.ParseProperties(strings.NewReader("key1=val1 \n # \n key2=val2"))
	log.Printf("properties: %p", properties)

	//
	properties2, err := docker.FindAndParseProperties(os.Args)
	log.Printf("properties2: %v, %v", properties2, err)

}

