package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearPolicyStatusAction_Invoke_Happy exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearPolicyStatusAction_Invoke_Happy(t *testing.T) {
	r := &ClearPolicyStatusAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearPolicyStatusAction_Invoke_NilClient exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearPolicyStatusAction_Invoke_NilClient(t *testing.T) {
	r := &ClearPolicyStatusAction{}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearPolicyStatusAction_Invoke_BuildError exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearPolicyStatusAction_Invoke_BuildError(t *testing.T) {
	r := &ClearPolicyStatusAction{client: newMalformedBaseURLClient(t)}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearPolicyStatusAction_Invoke_SendError exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearPolicyStatusAction_Invoke_SendError(t *testing.T) {
	r := &ClearPolicyStatusAction{client: newTransportErrorClient(t)}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearPolicyStatusAction_Invoke_APIError exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearPolicyStatusAction_Invoke_APIError(t *testing.T) {
	r := &ClearPolicyStatusAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_policy_status")
}

// TestClearPolicyStatusAction_Invoke_APIErrorReadBody exercises ClearPolicyStatusAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearPolicyStatusAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearPolicyStatusAction{client: newMockClientReadErrorBody(t, 500)}
	m := ClearPolicyStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
