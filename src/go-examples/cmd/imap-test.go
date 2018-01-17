package main

import (
	"fmt"
	"go-examples/pkg/imap"
	"go-examples/pkg/util"
)

func main() {
	fmt.Println("main()")
	fileName := ""
	properties, _ := util.ParsePropertiesFromFile(fileName)
	properties["imap.url"] = "imap://172.16.0.125:143/Inbox"
	properties["imap.username"] = "user1"
	properties["imap.password"] = "123456"
	//
	client := imap.NewImapClient(properties)
	defer client.Disconnect()
	fmt.Println("client: ", client)
	//
	err := client.Connect()
	if err != nil {
		fmt.Println("Connect failed: ", err)
	}else{
		client.CheckMail()
	}


/*	headers := map[string]string{
		"sessionName": "3711",
		"sessionSeq":  "2138",
	}
	publisher.Post(headers, []byte("test data"))*/
}