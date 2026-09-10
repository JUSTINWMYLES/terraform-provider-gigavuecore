package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSshCiphersAction_Invoke_Happy exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSshCiphersAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSshCiphersAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSshCiphersAction_Invoke_NilClient exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSshCiphersAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSshCiphersAction{}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSshCiphersAction_Invoke_BuildError exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSshCiphersAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSshCiphersAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSshCiphersAction_Invoke_SendError exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSshCiphersAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSshCiphersAction{client: newTransportErrorClient(t)}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSshCiphersAction_Invoke_APIError exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSshCiphersAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSshCiphersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_ssh_ciphers")
}

// TestUpdateSshCiphersAction_Invoke_APIErrorReadBody exercises UpdateSshCiphersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSshCiphersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSshCiphersAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSshCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
