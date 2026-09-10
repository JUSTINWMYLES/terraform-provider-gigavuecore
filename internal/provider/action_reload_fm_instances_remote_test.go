package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReloadFmInstancesAction_Invoke_Happy exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReloadFmInstancesAction_Invoke_Happy(t *testing.T) {
	r := &ReloadFmInstancesAction{client: newMockClientStatus(t, 201, "{}")}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReloadFmInstancesAction_Invoke_NilClient exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReloadFmInstancesAction_Invoke_NilClient(t *testing.T) {
	r := &ReloadFmInstancesAction{}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReloadFmInstancesAction_Invoke_BuildError exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReloadFmInstancesAction_Invoke_BuildError(t *testing.T) {
	r := &ReloadFmInstancesAction{client: newMalformedBaseURLClient(t)}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReloadFmInstancesAction_Invoke_SendError exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReloadFmInstancesAction_Invoke_SendError(t *testing.T) {
	r := &ReloadFmInstancesAction{client: newTransportErrorClient(t)}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReloadFmInstancesAction_Invoke_APIError exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReloadFmInstancesAction_Invoke_APIError(t *testing.T) {
	r := &ReloadFmInstancesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reload_fm_instances")
}

// TestReloadFmInstancesAction_Invoke_APIErrorReadBody exercises ReloadFmInstancesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReloadFmInstancesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReloadFmInstancesAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReloadFmInstancesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
