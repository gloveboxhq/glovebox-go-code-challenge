package mail

import "encoding/json"

// This package is a stand-in for the official [sendgrid-go](https://github.com/sendgrid/sendgrid-go)
// library. It represents a 3rd party vendored library for purposes of the code challenge,
// please do not modify.

func NewSendClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

type Client struct {
	apiKey string
}

func (c *Client) Send(v3mail *V3Mail) error {
	return nil
}

func NewV3Mail() *V3Mail {
	return &V3Mail{}
}

type V3Mail struct {
	message          json.RawMessage
	tplID            string
	fromName         string
	fromAddress      string
	personalizations []Personalization
}

func (m *V3Mail) SetMessage(msg json.RawMessage) *V3Mail {
	m.message = msg
	return m
}

func (m *V3Mail) SetTemplateID(tplID string) *V3Mail {
	m.tplID = tplID
	return m
}

func (m *V3Mail) SetFromName(fromName string) *V3Mail {
	m.fromName = fromName
	return m
}

func (m *V3Mail) SetFromAddress(fromAddress string) *V3Mail {
	m.fromAddress = fromAddress
	return m
}

func (m *V3Mail) AddPersonalization(p *Personalization) *V3Mail {
	m.personalizations = append(m.personalizations, *p)
	return m
}

func NewPersonalization() *Personalization {
	return &Personalization{}
}

type Personalization struct {
	tos  []string
	ccs  []string
	bccs []string
}

func (p *Personalization) AddTos(tos ...string) *Personalization {
	p.tos = append(p.tos, tos...)
	return p
}

func (p *Personalization) AddCCs(ccs ...string) *Personalization {
	p.ccs = append(p.ccs, ccs...)
	return p
}

func (p *Personalization) AddBCCs(bccs ...string) *Personalization {
	p.bccs = append(p.bccs, bccs...)
	return p
}
