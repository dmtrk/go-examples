package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"go-examples/pkg/util"
	"bytes"
)

type HttpClient struct {
	client   *http.Client
	Url      string
	Username string
	Password string
}

func NewHttpClient(properties map[string]string) *HttpClient {
	instance := new(HttpClient)
	instance.client = &http.Client{}
	instance.Url = util.GetStr(properties, "http.url", "https://httpbin.org/post")
	instance.Username = util.GetStr(properties, "http.username", "")
	instance.Password = util.GetStr(properties, "http.password", "")
	//
	return instance
}

func (self *HttpClient) Post(headers map[string]string, data []byte) (error) {
	return self.Do("POST", headers, data)
}

func (self *HttpClient) Get(headers map[string]string, data []byte) (error) {
	return self.Do("GET", headers, data)
}

func (self *HttpClient) Do(method string, headers map[string]string, data []byte) (error) {
	fmt.Println("Publish()")
	var err error
	req, err := http.NewRequest(method, self.Url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	for key, val := range headers {
		fmt.Println("Publish() adding header ",key,val)
		req.Header.Add(key, val)
	}
	if (len(self.Username) > 0 && len(self.Password) > 0) {
		req.SetBasicAuth(self.Username, self.Password)
	}
	if len(req.Header.Get("Content-Type")) == 0 {
		req.Header.Add("Content-Type", "text/plain")
	}

	// Do POST
	resp, err := self.client.Do(req)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Publish() response: ", string(body))
	return err
}

func (self *HttpClient) Disconnect(){
	fmt.Println("Disconnect()")
}
