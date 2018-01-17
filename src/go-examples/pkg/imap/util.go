package imap

import (
	"net/url"
	"fmt"
	"strings"
	"github.com/mxk/go-imap/imap"
)

func ParseUrl(urlStr string) (*url.URL) {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Errorf("Failed to parse URL object from string '%v'", urlStr)
	}
	return u
}

func GetFolderName(urlObj *url.URL) (string) {
	if len(urlObj.Path) > 0 {
		norm := strings.TrimSpace(urlObj.Path)
		return strings.TrimPrefix(norm, "/")
	}
	return "Inbox"
}

func Dial(urlObj *url.URL, useSsl string) (c *imap.Client, err error) {
	fmt.Println("Dial() ", urlObj, useSsl)

	if (strings.ToLower(useSsl) == "true") {
		c, err = imap.DialTLS(urlObj.Host, nil)
	} else {
		c, err = imap.Dial(urlObj.Host)
	}
	return c, err
}