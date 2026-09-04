package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearPtpCountersByAliasAction_Invoke_Happy exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearPtpCountersByAliasAction_Invoke_Happy(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearPtpCountersByAliasAction_Invoke_NilClient exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearPtpCountersByAliasAction_Invoke_NilClient(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearPtpCountersByAliasAction_Invoke_BuildError exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearPtpCountersByAliasAction_Invoke_BuildError(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{client: newMalformedBaseURLClient(t)}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearPtpCountersByAliasAction_Invoke_SendError exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearPtpCountersByAliasAction_Invoke_SendError(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{client: newTransportErrorClient(t)}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearPtpCountersByAliasAction_Invoke_APIError exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearPtpCountersByAliasAction_Invoke_APIError(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_ptp_counters_by_alias")
}

// TestClearPtpCountersByAliasAction_Invoke_APIErrorReadBody exercises ClearPtpCountersByAliasAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearPtpCountersByAliasAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearPtpCountersByAliasAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearPtpCountersByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
