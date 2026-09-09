package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateFmInstanceAction_Invoke_Happy exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateFmInstanceAction_Invoke_Happy(t *testing.T) {
	r := &UpdateFmInstanceAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateFmInstanceAction_Invoke_NilClient exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateFmInstanceAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateFmInstanceAction{}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateFmInstanceAction_Invoke_BuildError exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateFmInstanceAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateFmInstanceAction{client: newMalformedBaseURLClient(t)}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateFmInstanceAction_Invoke_SendError exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateFmInstanceAction_Invoke_SendError(t *testing.T) {
	r := &UpdateFmInstanceAction{client: newTransportErrorClient(t)}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateFmInstanceAction_Invoke_APIError exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateFmInstanceAction_Invoke_APIError(t *testing.T) {
	r := &UpdateFmInstanceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_fm_instance")
}

// TestUpdateFmInstanceAction_Invoke_APIErrorReadBody exercises UpdateFmInstanceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateFmInstanceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateFmInstanceAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
