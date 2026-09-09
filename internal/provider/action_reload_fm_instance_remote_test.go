package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReloadFmInstanceAction_Invoke_Happy exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReloadFmInstanceAction_Invoke_Happy(t *testing.T) {
	r := &ReloadFmInstanceAction{client: newMockClientStatus(t, 201, "{}")}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReloadFmInstanceAction_Invoke_NilClient exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReloadFmInstanceAction_Invoke_NilClient(t *testing.T) {
	r := &ReloadFmInstanceAction{}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReloadFmInstanceAction_Invoke_BuildError exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReloadFmInstanceAction_Invoke_BuildError(t *testing.T) {
	r := &ReloadFmInstanceAction{client: newMalformedBaseURLClient(t)}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReloadFmInstanceAction_Invoke_SendError exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReloadFmInstanceAction_Invoke_SendError(t *testing.T) {
	r := &ReloadFmInstanceAction{client: newTransportErrorClient(t)}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReloadFmInstanceAction_Invoke_APIError exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReloadFmInstanceAction_Invoke_APIError(t *testing.T) {
	r := &ReloadFmInstanceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reload_fm_instance")
}

// TestReloadFmInstanceAction_Invoke_APIErrorReadBody exercises ReloadFmInstanceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReloadFmInstanceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReloadFmInstanceAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReloadFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
