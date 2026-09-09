package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetFmInstanceAction_Invoke_Happy exercises ResetFmInstanceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetFmInstanceAction_Invoke_Happy(t *testing.T) {
	r := &ResetFmInstanceAction{client: newMockClientStatus(t, 201, "{}")}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetFmInstanceAction_Invoke_NilClient exercises ResetFmInstanceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetFmInstanceAction_Invoke_NilClient(t *testing.T) {
	r := &ResetFmInstanceAction{}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetFmInstanceAction_Invoke_BuildError exercises ResetFmInstanceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetFmInstanceAction_Invoke_BuildError(t *testing.T) {
	r := &ResetFmInstanceAction{client: newMalformedBaseURLClient(t)}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetFmInstanceAction_Invoke_SendError exercises ResetFmInstanceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetFmInstanceAction_Invoke_SendError(t *testing.T) {
	r := &ResetFmInstanceAction{client: newTransportErrorClient(t)}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetFmInstanceAction_Invoke_APIError exercises ResetFmInstanceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetFmInstanceAction_Invoke_APIError(t *testing.T) {
	r := &ResetFmInstanceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_fm_instance")
}

// TestResetFmInstanceAction_Invoke_APIErrorReadBody exercises ResetFmInstanceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetFmInstanceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetFmInstanceAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
