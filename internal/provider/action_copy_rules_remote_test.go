package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCopyRulesAction_Invoke_Happy exercises CopyRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCopyRulesAction_Invoke_Happy(t *testing.T) {
	r := &CopyRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCopyRulesAction_Invoke_NilClient exercises CopyRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCopyRulesAction_Invoke_NilClient(t *testing.T) {
	r := &CopyRulesAction{}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCopyRulesAction_Invoke_BuildError exercises CopyRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCopyRulesAction_Invoke_BuildError(t *testing.T) {
	r := &CopyRulesAction{client: newMalformedBaseURLClient(t)}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCopyRulesAction_Invoke_SendError exercises CopyRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCopyRulesAction_Invoke_SendError(t *testing.T) {
	r := &CopyRulesAction{client: newTransportErrorClient(t)}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCopyRulesAction_Invoke_APIError exercises CopyRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCopyRulesAction_Invoke_APIError(t *testing.T) {
	r := &CopyRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_copy_rules")
}

// TestCopyRulesAction_Invoke_APIErrorReadBody exercises CopyRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCopyRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CopyRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := CopyRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
