package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGetPolicyRulesAction_Invoke_Happy exercises GetPolicyRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGetPolicyRulesAction_Invoke_Happy(t *testing.T) {
	r := &GetPolicyRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetPolicyRulesAction_Invoke_NilClient exercises GetPolicyRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetPolicyRulesAction_Invoke_NilClient(t *testing.T) {
	r := &GetPolicyRulesAction{}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetPolicyRulesAction_Invoke_BuildError exercises GetPolicyRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetPolicyRulesAction_Invoke_BuildError(t *testing.T) {
	r := &GetPolicyRulesAction{client: newMalformedBaseURLClient(t)}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetPolicyRulesAction_Invoke_SendError exercises GetPolicyRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetPolicyRulesAction_Invoke_SendError(t *testing.T) {
	r := &GetPolicyRulesAction{client: newTransportErrorClient(t)}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetPolicyRulesAction_Invoke_APIError exercises GetPolicyRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetPolicyRulesAction_Invoke_APIError(t *testing.T) {
	r := &GetPolicyRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_get_policy_rules")
}

// TestGetPolicyRulesAction_Invoke_APIErrorReadBody exercises GetPolicyRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetPolicyRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GetPolicyRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := GetPolicyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
