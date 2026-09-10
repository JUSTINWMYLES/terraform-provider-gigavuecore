package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetUserAccountLocksAction_Invoke_Happy exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetUserAccountLocksAction_Invoke_Happy(t *testing.T) {
	r := &ResetUserAccountLocksAction{client: newMockClientStatus(t, 204, "{}")}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetUserAccountLocksAction_Invoke_NilClient exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetUserAccountLocksAction_Invoke_NilClient(t *testing.T) {
	r := &ResetUserAccountLocksAction{}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetUserAccountLocksAction_Invoke_BuildError exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetUserAccountLocksAction_Invoke_BuildError(t *testing.T) {
	r := &ResetUserAccountLocksAction{client: newMalformedBaseURLClient(t)}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetUserAccountLocksAction_Invoke_SendError exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetUserAccountLocksAction_Invoke_SendError(t *testing.T) {
	r := &ResetUserAccountLocksAction{client: newTransportErrorClient(t)}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetUserAccountLocksAction_Invoke_APIError exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetUserAccountLocksAction_Invoke_APIError(t *testing.T) {
	r := &ResetUserAccountLocksAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_user_account_locks")
}

// TestResetUserAccountLocksAction_Invoke_APIErrorReadBody exercises ResetUserAccountLocksAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetUserAccountLocksAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetUserAccountLocksAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetUserAccountLocksActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
