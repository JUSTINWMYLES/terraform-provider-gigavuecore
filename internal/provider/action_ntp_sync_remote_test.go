package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestNtpSyncAction_Invoke_Happy exercises NtpSyncAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestNtpSyncAction_Invoke_Happy(t *testing.T) {
	r := &NtpSyncAction{client: newMockClientStatus(t, 200, "{}")}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNtpSyncAction_Invoke_NilClient exercises NtpSyncAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNtpSyncAction_Invoke_NilClient(t *testing.T) {
	r := &NtpSyncAction{}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNtpSyncAction_Invoke_BuildError exercises NtpSyncAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNtpSyncAction_Invoke_BuildError(t *testing.T) {
	r := &NtpSyncAction{client: newMalformedBaseURLClient(t)}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNtpSyncAction_Invoke_SendError exercises NtpSyncAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestNtpSyncAction_Invoke_SendError(t *testing.T) {
	r := &NtpSyncAction{client: newTransportErrorClient(t)}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNtpSyncAction_Invoke_APIError exercises NtpSyncAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNtpSyncAction_Invoke_APIError(t *testing.T) {
	r := &NtpSyncAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_ntp_sync")
}

// TestNtpSyncAction_Invoke_APIErrorReadBody exercises NtpSyncAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNtpSyncAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &NtpSyncAction{client: newMockClientReadErrorBody(t, 501)}
	m := NtpSyncActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
