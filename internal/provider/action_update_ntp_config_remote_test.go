package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateNtpConfigAction_Invoke_Happy exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateNtpConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateNtpConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateNtpConfigAction_Invoke_NilClient exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateNtpConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateNtpConfigAction{}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateNtpConfigAction_Invoke_BuildError exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateNtpConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateNtpConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateNtpConfigAction_Invoke_SendError exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateNtpConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateNtpConfigAction{client: newTransportErrorClient(t)}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateNtpConfigAction_Invoke_APIError exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateNtpConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateNtpConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_ntp_config")
}

// TestUpdateNtpConfigAction_Invoke_APIErrorReadBody exercises UpdateNtpConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateNtpConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateNtpConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
