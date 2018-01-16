package protobuf

import "github.com/golang/protobuf/proto"


func NewBytesMessage(headers map[string]string, data string) *BytesMessage {
	msg := &BytesMessage{}
	msg.Data = []byte(data)
	for key, value := range headers {
		header := new(Header)
		header.Name = key
		header.Value = value
		msg.Headers = append(msg.Headers, header)
	}
	return msg;
}

func BytesMessageToBytes(msg *BytesMessage) ([]byte, error) {
	return proto.Marshal(msg)
}

func BytesMessageFromBytes(data []byte) (*BytesMessage, error) {
	msg := &BytesMessage{}
	err := proto.Unmarshal(data, msg)
	return msg,err;
}


func NewFieldsMessage(headers map[string]string, fields map[string]string) *FieldsMessage {
	msg := &FieldsMessage{}
	for key, value := range headers {
		header := new(Header)
		header.Name = key
		header.Value = value
		msg.Headers = append(msg.Headers, header)
	}
	for key, value := range fields {
		field := new(Field)
		field.Name = key
		field.Value = value
		msg.Fields = append(msg.Fields, field)
	}
	return msg;
}

func FieldsMessageToBytes(msg *FieldsMessage) ([]byte, error) {
	return proto.Marshal(msg)
}

func FieldsMessageFromBytes(data []byte) (*FieldsMessage, error) {
	msg := &FieldsMessage{}
	err := proto.Unmarshal(data, msg)
	return msg,err;
}