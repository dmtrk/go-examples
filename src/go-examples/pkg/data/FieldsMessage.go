package data

import (
	"go-examples/pkg/protobuf"
	"github.com/golang/protobuf/proto"
)

type FieldsMessage struct  {
	Headers []*Header
	Fields  []*Field
}

type Header struct  {
	Name, Value string
}

type Field struct  {
	Name, Value string
}

func (self *FieldsMessage) AddHeader(name string, value string){
	h := &Header{name,value}
	self.Headers = append(self.Headers, h)
}

func (self *FieldsMessage) AddField(name string, value string){
	f := &Field{name,value}
	self.Fields = append(self.Fields, f)
}

func (self *FieldsMessage) Serialize() ([]byte, error) {
	msg := &protobuf.FieldsMessage{}
	for key, value := range self.Headers {
		header := new(protobuf.Header)
		header.Name = key
		header.Value = value
		msg.Headers = append(msg.Headers, header)
	}
	for key, value := range self.Fields {
		header := new(protobuf.Field)
		header.Name = key
		header.Value = value
		msg.Fields = append(msg.Fields, header)
	}
	return proto.Marshal(msg)
}

func DeserializeFieldsMessage(data []byte) (*FieldsMessage, error) {
	msg := &protobuf.FieldsMessage{}
	err := proto.Unmarshal(data, msg)
	if err == nil {
		message := &FieldsMessage{}
		message.Headers = make(map[string]string, len(msg.Headers))
		for _, element := range msg.Headers {
			message.Headers[element.Name]=element.Value
		}
		for _, element := range msg.Fields {
			message.Fields[element.Name]=element.Value
		}
		return message,nil;
	}
	return nil,err;
}