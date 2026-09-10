package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSystemHostnameAction_Invoke_Happy exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSystemHostnameAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSystemHostnameAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSystemHostnameAction_Invoke_NilClient exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSystemHostnameAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSystemHostnameAction{}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSystemHostnameAction_Invoke_BuildError exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSystemHostnameAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSystemHostnameAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSystemHostnameAction_Invoke_SendError exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSystemHostnameAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSystemHostnameAction{client: newTransportErrorClient(t)}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSystemHostnameAction_Invoke_APIError exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSystemHostnameAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSystemHostnameAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_system_hostname")
}

// TestUpdateSystemHostnameAction_Invoke_APIErrorReadBody exercises UpdateSystemHostnameAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSystemHostnameAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSystemHostnameAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSystemHostnameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
