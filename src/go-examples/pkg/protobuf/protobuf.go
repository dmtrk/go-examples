package protobuf

import "github.com/golang/protobuf/proto"

type ProtoMessage2 interface {
	GetHeaders() []*Header
}


func AddHeader(message *ProtoMessage2, name string, value string){
	h := &Header{name,value}
	message.GetHeaders() = append(message.GetHeaders(), h)
}

func SerializeBytesMessage(message *BytesMessage) ([]byte, error) {
	return proto.Marshal(message)
}


func DeserializeBytesMessage(data []byte) (*BytesMessage, error) {
	msg := &BytesMessage{}
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


func SerializeFieldsMessage(message *FieldsMessage) ([]byte, error) {
	return proto.Marshal(message)
}

func DeserializeFieldsMessage(data []byte) (*FieldsMessage, error) {
	protoMessage := &FieldsMessage{}
	err := proto.Unmarshal(data, protoMessage)
	if err == nil {
		message := &FieldsMessage{}
		message.Headers = make(map[string]string, len(protoMessage.Headers))
		for _, element := range protoMessage.Headers {
			message.Headers[element.Name]=element.Value
		}
		for _, element := range protoMessage.Fields {
			message.Fields[element.Name]=element.Value
		}
		return message,nil;
	}
	return nil,err;
}
