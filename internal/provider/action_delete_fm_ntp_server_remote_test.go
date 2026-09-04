package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteFmNtpServerAction_Invoke_Happy exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteFmNtpServerAction_Invoke_Happy(t *testing.T) {
	r := &DeleteFmNtpServerAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteFmNtpServerAction_Invoke_NilClient exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteFmNtpServerAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteFmNtpServerAction{}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteFmNtpServerAction_Invoke_BuildError exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteFmNtpServerAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteFmNtpServerAction{client: newMalformedBaseURLClient(t)}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteFmNtpServerAction_Invoke_SendError exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteFmNtpServerAction_Invoke_SendError(t *testing.T) {
	r := &DeleteFmNtpServerAction{client: newTransportErrorClient(t)}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteFmNtpServerAction_Invoke_APIError exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteFmNtpServerAction_Invoke_APIError(t *testing.T) {
	r := &DeleteFmNtpServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_fm_ntp_server")
}

// TestDeleteFmNtpServerAction_Invoke_APIErrorReadBody exercises DeleteFmNtpServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteFmNtpServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteFmNtpServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
