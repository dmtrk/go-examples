package protobuf

import (
	"github.com/golang/protobuf/proto"
)

//####### BytesMessage
func NewBytesMessage(headers map[string]string, data string) *BytesMessage {
	message := &BytesMessage{}
	message.Data = []byte(data)
	for key, value := range headers {
		header := new(Header)
		header.Name = key
		header.Value = value
		message.Headers = append(message.Headers, header)
	}
	return message;
}

func AddHeaderBytesMessage(message *BytesMessage, name string, value string) {
	h := &Header{name, value}
	message.Headers = append(message.GetHeaders(), h)
}

func SerializeBytesMessage(message *BytesMessage) ([]byte, error) {
	return proto.Marshal(message)
}

func DeserializeBytesMessage(data []byte) (*BytesMessage, error) {
	message := &BytesMessage{}
	return message, proto.Unmarshal(data, message)
}

//####### FieldsMessage
func NewFieldsMessage(headers map[string]string, fields map[string]string) *FieldsMessage {
	message := &FieldsMessage{}
	for key, value := range headers {
		header := new(Header)
		header.Name = key
		header.Value = value
		message.Headers = append(message.Headers, header)
	}
	for key, value := range fields {
		field := new(Field)
		field.Name = key
		field.Value = value
		message.Fields = append(message.Fields, field)
	}
	return message;
}

func AddHeaderFieldsMessage(message *FieldsMessage, name string, value string) {
	h := &Header{name, value}
	message.Headers = append(message.GetHeaders(), h)
}

func SerializeFieldsMessage(message *FieldsMessage) ([]byte, error) {
	return proto.Marshal(message)
}

func DeserializeFieldsMessage(data []byte) (*FieldsMessage, error) {
	message := &FieldsMessage{}
	return message, proto.Unmarshal(data, message)
}
