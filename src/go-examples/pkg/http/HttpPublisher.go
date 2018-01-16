package http

import (
	"fmt"
	"go-examples/pkg/protobuf"
)

type HttpPublisher struct  {

}

func (self *HttpPublisher) Publish(headers map[string]string, data string){
	fmt.Println("Publish()")
	msg := protobuf.NewBytesMessage(headers,data)
	fmt.Println("msg: ",msg.String())
	fmt.Println("msg.Headers: ",msg.Headers)
	fmt.Println("msg.Data:    ",msg.Data)
}