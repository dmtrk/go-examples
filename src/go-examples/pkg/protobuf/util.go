package protobuf

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