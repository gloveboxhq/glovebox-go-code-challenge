package sendgrid

import (
	"encoding/json"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"

	// This is a stand-in for the 3rd party sendgrid library
	// In a real project this would be vendored as github.com/sendgrid/sendgrid-go
	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email/sendgrid/mail"
)

type Config struct {
	APIKey      string
	FromAddress string
	FromName    string
}

func NewSvc(cfg Config) *Client {
	return &Client{
		client:      mail.NewSendClient(cfg.APIKey),
		fromAddress: cfg.FromAddress,
		fromName:    cfg.FromName,
	}
}

type Client struct {
	client      *mail.Client
	fromAddress string
	fromName    string
}

func (c *Client) Send(to []string, cc []string, message json.RawMessage, tpl email.TplID) error {

	// create personalization
	p := mail.NewPersonalization().
		AddTos(to...).
		AddCCs(cc...)

	// create the email
	m := mail.NewV3Mail().
		SetMessage(message).
		SetTemplateID(string(tpl)).
		SetFromName(c.fromName).
		SetFromAddress(c.fromAddress).
		AddPersonalization(p)

	// send the email
	if err := c.client.Send(m); err != nil {
		return err
	}

	return nil
}
