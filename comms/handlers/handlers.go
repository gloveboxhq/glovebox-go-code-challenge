package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"
)

func AddPolicyVehicle(emailsvc email.MailProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
			return
		}

		payload := AddPolicyVehicleReq{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if err := emailsvc.Send([]string{payload.EmailTo}, nil, payload.Message, email.TplAddPolicyVehicle); err != nil {
			http.Error(w, fmt.Sprintf("error sending email: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func AddPolicyDriver(emailsvc email.MailProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		payload := AddPolicyDriverReq{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if err := emailsvc.Send([]string{payload.EmailTo}, nil, payload.Message, email.TplAddPolicyDriver); err != nil {
			http.Error(w, fmt.Sprintf("error sending email: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func AddPolicyAddress(emailsvc email.MailProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		payload := AddPolicyAddressReq{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if err := emailsvc.Send([]string{payload.EmailTo}, nil, payload.Message, email.TplAddPolicyAddress); err != nil {
			http.Error(w, fmt.Sprintf("error sending email: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func AddPolicyCoverage(emailsvc email.MailProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		payload := AddPolicyCoverageReq{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if err := emailsvc.Send([]string{payload.EmailTo}, []string{payload.EmailCc}, payload.Message, email.TplAddPolicyCoverage); err != nil {
			http.Error(w, fmt.Sprintf("error sending email: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
