package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email"
	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/email/mockemail"
	"github.com/gloveboxhq/glovebox-go-code-challenge/comms/handlers"
)

func TestAddPolicyVehicle(t *testing.T) {

	t.Parallel()

	type testCase struct {
		method       string
		payload      handlers.AddPolicyVehicleReq
		expectTplID  email.TplID
		expectStatus int
	}

	testCases := map[string]testCase{
		"pass": {
			method: http.MethodPost,
			payload: handlers.AddPolicyVehicleReq{
				EmailTo: "foo@bar.com",
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyVehicle,
			expectStatus: http.StatusOK,
		},
		"fail invalid method": {
			method:       http.MethodGet,
			payload:      handlers.AddPolicyVehicleReq{},
			expectStatus: http.StatusMethodNotAllowed,
		},
	}

	testFactory := func(tc testCase) func(*testing.T) {
		return func(t *testing.T) {

			payload, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatalf("could nor marshal payload to json: %v", err)
			}

			testEmail := mockemail.NewClient()
			defer testEmail.FlushSendLogs()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/comms/add-policy-vehicle", bytes.NewReader(payload))

			handlers.AddPolicyVehicle(testEmail)(w, req)

			resp := w.Result()

			if resp.StatusCode != tc.expectStatus {
				t.Fatalf("expected status %v but got %v", tc.expectStatus, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {

				if testEmail.SendLogs().IsEmpty() {
					t.Fatalf("expected email log but got empty")
				}

				lastEmail := testEmail.SendLogs().Last()

				tos := lastEmail.ExtractTos()
				expectedTos := []string{tc.payload.EmailTo}

				if len(tos) != len(expectedTos) {
					t.Fatalf("mismatched to lengths, expected %v but got %v", expectedTos, tos)
				}

				for i := range tos {
					to := tos[i]
					expectedTo := string(expectedTos[i])
					if to != expectedTo {
						t.Fatalf("expected to %v but got %v", expectedTo, to)
					}
				}

				if string(lastEmail.ExtractMessage()) != string(tc.payload.Message) {
					t.Fatalf("expected message %v but got %v", tc.payload.Message, lastEmail.ExtractMessage())
				}

				if lastEmail.ExtractTplID() != tc.expectTplID {
					t.Fatalf("expected tpl %v but got %v", tc.expectTplID, lastEmail.ExtractTplID())
				}
			}
		}
	}

	for name, tc := range testCases {
		t.Run(name, testFactory(tc))
	}
}

func TestAddPolicyDriver(t *testing.T) {

	t.Parallel()

	type testCase struct {
		method       string
		payload      handlers.AddPolicyDriverReq
		expectTplID  email.TplID
		expectStatus int
	}

	testCases := map[string]testCase{
		"pass": {
			method: http.MethodPost,
			payload: handlers.AddPolicyDriverReq{
				EmailTo: "foo@bar.com",
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyDriver,
			expectStatus: http.StatusOK,
		},
		"fail invalid method": {
			method:       http.MethodGet,
			payload:      handlers.AddPolicyDriverReq{},
			expectStatus: http.StatusMethodNotAllowed,
		},
	}

	testFactory := func(tc testCase) func(*testing.T) {
		return func(t *testing.T) {

			payload, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatalf("could nor marshal payload to json: %v", err)
			}

			testEmail := mockemail.NewClient()
			defer testEmail.FlushSendLogs()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/comms/add-policy-driver", bytes.NewReader(payload))

			handlers.AddPolicyDriver(testEmail)(w, req)

			resp := w.Result()

			if resp.StatusCode != tc.expectStatus {
				t.Fatalf("expected status %v but got %v", tc.expectStatus, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {

				if testEmail.SendLogs().IsEmpty() {
					t.Fatalf("expected email log but got empty")
				}

				lastEmail := testEmail.SendLogs().Last()

				tos := lastEmail.ExtractTos()
				expectedTos := []string{tc.payload.EmailTo}

				if len(tos) != len(expectedTos) {
					t.Fatalf("mismatched to lengths, expected %v but got %v", expectedTos, tos)
				}

				for i := range tos {
					to := tos[i]
					expectedTo := string(expectedTos[i])
					if to != expectedTo {
						t.Fatalf("expected to %v but got %v", expectedTo, to)
					}
				}

				if string(lastEmail.ExtractMessage()) != string(tc.payload.Message) {
					t.Fatalf("expected message %v but got %v", tc.payload.Message, lastEmail.ExtractMessage())
				}

				if lastEmail.ExtractTplID() != tc.expectTplID {
					t.Fatalf("expected tpl %v but got %v", tc.expectTplID, lastEmail.ExtractTplID())
				}
			}
		}
	}

	for name, tc := range testCases {
		t.Run(name, testFactory(tc))
	}
}

func TestAddPolicyCoverage(t *testing.T) {

	t.Parallel()

	type testCase struct {
		method       string
		payload      handlers.AddPolicyCoverageReq
		expectTplID  email.TplID
		expectStatus int
	}

	testCases := map[string]testCase{
		"pass": {
			method: http.MethodPost,
			payload: handlers.AddPolicyCoverageReq{
				EmailTo: []string{"foo@bar.com"},
				EmailCC: []string{"baz@bar.com"},
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyCoverage,
			expectStatus: http.StatusOK,
		},
		"pass 2 to": {
			method: http.MethodPost,
			payload: handlers.AddPolicyCoverageReq{
				EmailTo: []string{"foo@bar.com", "foo@baz.com"},
				EmailCC: []string{"baz@bar.com"},
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyCoverage,
			expectStatus: http.StatusOK,
		},
		"pass 2 cc": {
			method: http.MethodPost,
			payload: handlers.AddPolicyCoverageReq{
				EmailTo: []string{"baz@bar.com"},
				EmailCC: []string{"foo@bar.com", "foo@baz.com"},
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyCoverage,
			expectStatus: http.StatusOK,
		},
		"pass 2 to and 2 cc": {
			method: http.MethodPost,
			payload: handlers.AddPolicyCoverageReq{
				EmailTo: []string{"baz@bar.com", "fizz@buzz.com"},
				EmailCC: []string{"foo@bar.com", "foo@baz.com"},
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyCoverage,
			expectStatus: http.StatusOK,
		},
		"pass many": {
			method: http.MethodPost,
			payload: handlers.AddPolicyCoverageReq{
				EmailTo: []string{"baz@bar.com", "fizz@buzz.com", "beep@boop.com"},
				EmailCC: []string{"foo@bar.com", "foo@baz.com", "zig@zag.com"},
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyCoverage,
			expectStatus: http.StatusOK,
		},
		"fail invalid method": {
			method:       http.MethodGet,
			payload:      handlers.AddPolicyCoverageReq{},
			expectStatus: http.StatusMethodNotAllowed,
		},
	}

	testFactory := func(tc testCase) func(*testing.T) {
		return func(t *testing.T) {

			payload, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatalf("could nor marshal payload to json: %v", err)
			}

			testEmail := mockemail.NewClient()
			defer testEmail.FlushSendLogs()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/comms/add-policy-coverage", bytes.NewReader(payload))

			handlers.AddPolicyCoverage(testEmail)(w, req)

			resp := w.Result()

			if resp.StatusCode != tc.expectStatus {
				t.Fatalf("expected status %v but got %v", tc.expectStatus, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {

				if testEmail.SendLogs().IsEmpty() {
					t.Fatalf("expected email log but got empty")
				}

				lastEmail := testEmail.SendLogs().Last()

				{
					// Check all Tos
					tos := lastEmail.ExtractTos()
					expectedTos := tc.payload.EmailTo

					if len(tos) != len(expectedTos) {
						t.Fatalf("mismatched to lengths, expected %v but got %v", expectedTos, tos)
					}

					for i := range tos {
						to := tos[i]
						expectedTo := string(expectedTos[i])
						if to != expectedTo {
							t.Fatalf("expected to %v but got %v", expectedTo, to)
						}
					}
				}

				{
					// Check all CCs
					ccs := lastEmail.ExtractCCs()
					expectedCCs := tc.payload.EmailCC

					if len(ccs) != len(expectedCCs) {
						t.Fatalf("mismatched cc lengths, expected %v but got %v", expectedCCs, ccs)
					}

					for i := range ccs {
						cc := ccs[i]
						expectedCC := string(expectedCCs[i])
						if cc != expectedCC {
							t.Fatalf("expected cc %v but got %v", expectedCC, cc)
						}
					}
				}

				if string(lastEmail.ExtractMessage()) != string(tc.payload.Message) {
					t.Fatalf("expected message %v but got %v", tc.payload.Message, lastEmail.ExtractMessage())
				}

				if lastEmail.ExtractTplID() != tc.expectTplID {
					t.Fatalf("expected tpl %v but got %v", tc.expectTplID, lastEmail.ExtractTplID())
				}
			}
		}
	}

	for name, tc := range testCases {
		t.Run(name, testFactory(tc))
	}
}

func TestAddPolicyAddress(t *testing.T) {

	t.Parallel()

	type testCase struct {
		method       string
		payload      handlers.AddPolicyAddressReq
		expectTplID  email.TplID
		expectStatus int
	}

	testCases := map[string]testCase{
		"pass": {
			method: http.MethodPost,
			payload: handlers.AddPolicyAddressReq{
				EmailTo: "foo@bar.com",
				Message: json.RawMessage(`{"foo":"bar"}`),
			},
			expectTplID:  email.TplAddPolicyAddress,
			expectStatus: http.StatusOK,
		},
		"fail invalid method": {
			method:       http.MethodGet,
			payload:      handlers.AddPolicyAddressReq{},
			expectStatus: http.StatusMethodNotAllowed,
		},
	}

	testFactory := func(tc testCase) func(*testing.T) {
		return func(t *testing.T) {

			payload, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatalf("could nor marshal payload to json: %v", err)
			}

			testEmail := mockemail.NewClient()
			defer testEmail.FlushSendLogs()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/api/comms/add-policy-address", bytes.NewReader(payload))

			handlers.AddPolicyAddress(testEmail)(w, req)

			resp := w.Result()

			if resp.StatusCode != tc.expectStatus {
				t.Fatalf("expected status %v but got %v", tc.expectStatus, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {

				if testEmail.SendLogs().IsEmpty() {
					t.Fatalf("expected email log but got empty")
				}

				lastEmail := testEmail.SendLogs().Last()

				tos := lastEmail.ExtractTos()
				expectedTos := []string{tc.payload.EmailTo}

				if len(tos) != len(expectedTos) {
					t.Fatalf("mismatched to lengths, expected %v but got %v", expectedTos, tos)
				}

				for i := range tos {
					to := tos[i]
					expectedTo := string(expectedTos[i])
					if to != expectedTo {
						t.Fatalf("expected to %v but got %v", expectedTo, to)
					}
				}

				if string(lastEmail.ExtractMessage()) != string(tc.payload.Message) {
					t.Fatalf("expected message %v but got %v", tc.payload.Message, lastEmail.ExtractMessage())
				}

				if lastEmail.ExtractTplID() != tc.expectTplID {
					t.Fatalf("expected tpl %v but got %v", tc.expectTplID, lastEmail.ExtractTplID())
				}
			}
		}
	}

	for name, tc := range testCases {
		t.Run(name, testFactory(tc))
	}
}
