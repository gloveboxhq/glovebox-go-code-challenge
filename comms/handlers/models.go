package handlers

import (
	"encoding/json"
)

type AddPolicyVehicleReq struct {
	EmailTo string          `json:"email_to"`
	Message json.RawMessage `json:"message"`
}

type AddPolicyDriverReq struct {
	EmailTo string          `json:"email_to"`
	Message json.RawMessage `json:"message"`
}

type AddPolicyAddressReq struct {
	EmailTo string `json:"email_to"`

	Message json.RawMessage `json:"message"`
}

type AddPolicyCoverageReq struct {
	EmailTo string          `json:"email_to"`
	EmailCc string          `json:"email_cc"`
	Message json.RawMessage `json:"message"`
}
