package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearPolicyAction_Invoke_Happy exercises ClearPolicyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearPolicyAction_Invoke_Happy(t *testing.T) {
	r := &ClearPolicyAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearPolicyAction_Invoke_NilClient exercises ClearPolicyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearPolicyAction_Invoke_NilClient(t *testing.T) {
	r := &ClearPolicyAction{}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearPolicyAction_Invoke_BuildError exercises ClearPolicyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearPolicyAction_Invoke_BuildError(t *testing.T) {
	r := &ClearPolicyAction{client: newMalformedBaseURLClient(t)}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearPolicyAction_Invoke_SendError exercises ClearPolicyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearPolicyAction_Invoke_SendError(t *testing.T) {
	r := &ClearPolicyAction{client: newTransportErrorClient(t)}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearPolicyAction_Invoke_APIError exercises ClearPolicyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearPolicyAction_Invoke_APIError(t *testing.T) {
	r := &ClearPolicyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_policy")
}

// TestClearPolicyAction_Invoke_APIErrorReadBody exercises ClearPolicyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearPolicyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearPolicyAction{client: newMockClientReadErrorBody(t, 500)}
	m := ClearPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
