package email

import "encoding/json"

type TplID string

const (
	TplAddPolicyVehicle  TplID = "add-policy-vehicle"
	TplAddPolicyDriver   TplID = "add-policy-driver"
	TplAddPolicyAddress  TplID = "add-policy-address"
	TplAddPolicyCoverage TplID = "add-policy-coverage"
)

type MailProvider interface {
	Send(to []string, cc []string, message json.RawMessage, tpl TplID) error
}
