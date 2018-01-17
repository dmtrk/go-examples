package http

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

type HttpPublisher struct {
	client   *http.Client
	Url      string
	Username string
	Password string
}

func NewHttpPublisher(properties map[string]string) *HttpPublisher {
	instance := new(HttpPublisher)
	instance.client = &http.Client{}
	instance.Url = GetString(properties, "http.url", "https://httpbin.org/post")
	instance.Username = GetString(properties, "http.username", "")
	instance.Password = GetString(properties, "http.password", "")
	//
	return instance
}

func (self *HttpPublisher) Post(headers map[string]string, data string) (error) {
	return self.Do("POST", headers, data)
}

func (self *HttpPublisher) Get(headers map[string]string, data string) (error) {
	return self.Do("GET", headers, data)
}

func (self *HttpPublisher) Do(method string, headers map[string]string, data string) (error) {
	fmt.Println("Publish()")
	var err error
	req, err := http.NewRequest(method, self.Url, strings.NewReader(data))
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

func GetString(properties map[string]string, key string, defaultValue string) string {
	value := properties[key]
	if len(value) > 0 {
		return strings.TrimSpace(value)
	}
	return defaultValue
}
