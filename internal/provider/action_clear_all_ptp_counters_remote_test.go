package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearAllPtpCountersAction_Invoke_Happy exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearAllPtpCountersAction_Invoke_Happy(t *testing.T) {
	r := &ClearAllPtpCountersAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearAllPtpCountersAction_Invoke_NilClient exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearAllPtpCountersAction_Invoke_NilClient(t *testing.T) {
	r := &ClearAllPtpCountersAction{}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearAllPtpCountersAction_Invoke_BuildError exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearAllPtpCountersAction_Invoke_BuildError(t *testing.T) {
	r := &ClearAllPtpCountersAction{client: newMalformedBaseURLClient(t)}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearAllPtpCountersAction_Invoke_SendError exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearAllPtpCountersAction_Invoke_SendError(t *testing.T) {
	r := &ClearAllPtpCountersAction{client: newTransportErrorClient(t)}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearAllPtpCountersAction_Invoke_APIError exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearAllPtpCountersAction_Invoke_APIError(t *testing.T) {
	r := &ClearAllPtpCountersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_all_ptp_counters")
}

// TestClearAllPtpCountersAction_Invoke_APIErrorReadBody exercises ClearAllPtpCountersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearAllPtpCountersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearAllPtpCountersAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearAllPtpCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
