package mockemail

import (
	"encoding/json"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"
)

type Client struct {
	sendLogs SendLogs
}

func NewClient() *Client {
	return &Client{
		sendLogs: SendLogs{},
	}
}

func (c *Client) Send(tos []string, ccs []string, message json.RawMessage, tplID email.TplID) error {
	c.sendLogs = append(c.sendLogs, SendLog{
		tos:     tos,
		ccs:     ccs,
		message: message,
		tplID:   tplID,
	})

	return nil
}

func (c *Client) SendLogs() SendLogs {
	return c.sendLogs
}

func (c *Client) FlushSendLogs() {
	c.sendLogs = SendLogs{}
}
