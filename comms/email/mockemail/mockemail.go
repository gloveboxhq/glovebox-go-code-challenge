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

func (c *Client) Send(to []string, cc []string, message json.RawMessage, tplID email.TplID) error {

	for _, v := range to {
		c.sendLogs = append(c.sendLogs, SendLog{
			to:      v,
			cc:      cc,
			message: message,
			tplID:   tplID,
		})
	}

	return nil
}

func (c *Client) SendLogs() SendLogs {
	return c.sendLogs
}

func (c *Client) FlushSendLogs() {
	c.sendLogs = SendLogs{}
}
