package data

import (
	"github.com/golang/protobuf/proto"
	"go-examples/pkg/protobuf"
)

type BytesMessage struct  {
	Headers []*Header
	Data string
}

type Header struct  {
	Name, Value string
}

type Field struct  {
	Name, Value string
}
func (self *BytesMessage) AddHeader(name string, value string){
	h := &Header{name,value}
	self.Headers = append(self.Headers, h)
}

func (self *BytesMessage) Serialize() ([]byte, error) {
	msg := &protobuf.BytesMessage{}
	msg.Data = []byte(self.Data)
	for key, value := range self.Headers {
		header := new(protobuf.Header)
		header.Name = key
		header.Value = value
		msg.Headers = append(msg.Headers, header)
	}
	return proto.Marshal(msg)
}

func DeserializeBytesMessage(data []byte) (*BytesMessage, error) {
	msg := &protobuf.BytesMessage{}
	err := proto.Unmarshal(data, msg)
	if err == nil {
		message := &BytesMessage{}
		message.Data = string(msg.Data);
		message.Headers = make(map[string]string, len(msg.Headers))
		for _, element := range msg.Headers {
			message.Headers[element.Name]=element.Value
		}
		return message,nil;
	}
	return nil,err;
}
