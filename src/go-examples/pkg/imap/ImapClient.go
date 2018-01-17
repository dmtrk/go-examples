package imap

import (
	"github.com/mxk/go-imap/imap"
	"time"
	"fmt"
	"net/url"
	"go-examples/pkg/util"
	"errors"
)

const (
	DEFAULT_URL = "imap://localhost:143/Inbox"
)

type ImapClient struct {
	client   *imap.Client
	Folder   string
	//
	Url      *url.URL
	Username string
	Password string
	UseSsl   string
}

func NewImapClient(properties map[string]string) *ImapClient {
	instance := new(ImapClient)
	//
	instance.Url = ParseUrl(util.GetStr(properties, "imap.url", DEFAULT_URL))
	instance.Username = util.GetStr(properties, "imap.username", "")
	instance.Password = util.GetStr(properties, "imap.password", "")
	instance.UseSsl = util.GetStr(properties, "imap.use_ssl", "")
	//
	instance.Folder = GetFolderName(instance.Url)
	return instance
}

func (self *ImapClient) IsConnected() bool {
	return self.client != nil
}

func (self *ImapClient) Connect() error {
	shutdown(self.client)
	//
	var err error
	var c *imap.Client
	c, err = Dial(self.Url, self.UseSsl)
	if (err == nil) {
		if c.Caps["STARTTLS"] {
			_, err = c.StartTLS(nil)
		}
		if (err == nil) {
			_, err = c.Login(self.Username, self.Password)
			if (err == nil) {
				_, err = c.Select(self.Folder, false)
			}
		}
	}
	if err != nil {
		return err
	}
	self.client = c
	fmt.Println("Connected() client:", self.client)
	//
	return nil
}

func (self *ImapClient) Disconnect() {
	fmt.Println("Disconnect()")
	shutdown(self.client)
	self.client = nil;
}

func (self *ImapClient) CheckMail() (error) {
	fmt.Println("CheckMail()")
	if !self.IsConnected() {
		return errors.New("Not connected")
	} else {
		cmd, err := self.client.UIDSearch("UNDELETED")
		fmt.Println("CheckMail() ", cmd, err)

		return err
	}
}

func shutdown(client *imap.Client) {
	if client != nil {
		client.Logout(5 * time.Second)
		client.Close(true)
	}
}
