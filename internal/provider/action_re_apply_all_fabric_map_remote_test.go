package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReApplyAllFabricMapAction_Invoke_Happy exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReApplyAllFabricMapAction_Invoke_Happy(t *testing.T) {
	r := &ReApplyAllFabricMapAction{client: newMockClientStatus(t, 202, "{}")}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReApplyAllFabricMapAction_Invoke_NilClient exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReApplyAllFabricMapAction_Invoke_NilClient(t *testing.T) {
	r := &ReApplyAllFabricMapAction{}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReApplyAllFabricMapAction_Invoke_BuildError exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReApplyAllFabricMapAction_Invoke_BuildError(t *testing.T) {
	r := &ReApplyAllFabricMapAction{client: newMalformedBaseURLClient(t)}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReApplyAllFabricMapAction_Invoke_SendError exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReApplyAllFabricMapAction_Invoke_SendError(t *testing.T) {
	r := &ReApplyAllFabricMapAction{client: newTransportErrorClient(t)}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReApplyAllFabricMapAction_Invoke_APIError exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReApplyAllFabricMapAction_Invoke_APIError(t *testing.T) {
	r := &ReApplyAllFabricMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_re_apply_all_fabric_map")
}

// TestReApplyAllFabricMapAction_Invoke_APIErrorReadBody exercises ReApplyAllFabricMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReApplyAllFabricMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReApplyAllFabricMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReApplyAllFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
