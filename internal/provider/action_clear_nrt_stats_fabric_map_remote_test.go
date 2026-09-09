package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearNrtStatsFabricMapAction_Invoke_Happy exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearNrtStatsFabricMapAction_Invoke_Happy(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearNrtStatsFabricMapAction_Invoke_NilClient exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearNrtStatsFabricMapAction_Invoke_NilClient(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearNrtStatsFabricMapAction_Invoke_BuildError exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearNrtStatsFabricMapAction_Invoke_BuildError(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{client: newMalformedBaseURLClient(t)}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearNrtStatsFabricMapAction_Invoke_SendError exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearNrtStatsFabricMapAction_Invoke_SendError(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{client: newTransportErrorClient(t)}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearNrtStatsFabricMapAction_Invoke_APIError exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearNrtStatsFabricMapAction_Invoke_APIError(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_nrt_stats_fabric_map")
}

// TestClearNrtStatsFabricMapAction_Invoke_APIErrorReadBody exercises ClearNrtStatsFabricMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearNrtStatsFabricMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearNrtStatsFabricMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
