package email

import "encoding/json"

type TplID string

const (
	TplAddPolicyVehicle TplID = "add-policy-vehicle"
	TplAddPolicyDriver  TplID = "add-policy-driver"
	TplAddPolicyAddress TplID = "add-policy-address"
)

type MailProvider interface {
	Send(to []string, message json.RawMessage, tpl TplID) error
}
