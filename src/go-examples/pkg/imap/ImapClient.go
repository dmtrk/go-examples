package imap

import (
	"github.com/mxk/go-imap/imap"
	"time"
	"fmt"
	"net/url"
	"go-examples/pkg/util"
	"errors"
	"net/mail"
	"bytes"
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

func (self *ImapClient) CheckMail() (err error) {
	fmt.Println("CheckMail()")
	if !self.IsConnected() {
		return errors.New("Not connected")
	} else {
		var uids []uint32
		uids, err = self.uidSearch()
		fmt.Println("CheckMail(1) cmd: ", uids, err)
		if (err == nil) {
			for uid := range uids {
				if msg := self.uidFetch(uint32(uid)); msg != nil {
					fmt.Println("CheckMail(1) OK: ", uid, msg.Header.Get("Subject"))
				} else {
					fmt.Println("CheckMail(1) ERR: ", uid, msg)
				}
			}
		}
		return err
	}
}

func (self *ImapClient) uidSearch() ([]uint32, error) {
	cmd, err := self.client.UIDSearch("UNDELETED") //ALL
	if cmd.InProgress() {
		self.client.Recv(-1)
		//
		if (err == nil && len(cmd.Data) > 0) {
			results := cmd.Data[0].SearchResults()
			fmt.Println("uidSearch() results: ", results)
			return results, nil
		}
	}
	return make([]uint32, 0), err;
}

func (self *ImapClient) uidFetch(uid uint32) (msg *mail.Message) {
	set, _ := imap.NewSeqSet("")
	set.AddNum(uid)
	//
	cmd, err := self.client.UIDFetch(set, "RFC822")
	for cmd.InProgress() {
		self.client.Recv(-1)
		//
		if err == nil {
			var rsp *imap.Response
			for _, rsp = range cmd.Data {
				header := imap.AsBytes(rsp.MessageInfo().Attrs["RFC822"])
				if msg, err = mail.ReadMessage(bytes.NewReader(header)); msg != nil {
					return msg;
				}
			}
		} else {
			fmt.Println("uidFetch() err: ", err)
		}
	}
	return nil;
}

func shutdown(client *imap.Client) {
	if client != nil {
		client.Logout(5 * time.Second)
		client.Close(true)
	}
}
