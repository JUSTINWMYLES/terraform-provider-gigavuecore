package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRegisterNrtStatsFabricMapAction_Invoke_Happy exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRegisterNrtStatsFabricMapAction_Invoke_Happy(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{client: newMockClientStatus(t, 200, "{}")}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRegisterNrtStatsFabricMapAction_Invoke_NilClient exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRegisterNrtStatsFabricMapAction_Invoke_NilClient(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRegisterNrtStatsFabricMapAction_Invoke_BuildError exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRegisterNrtStatsFabricMapAction_Invoke_BuildError(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{client: newMalformedBaseURLClient(t)}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRegisterNrtStatsFabricMapAction_Invoke_SendError exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRegisterNrtStatsFabricMapAction_Invoke_SendError(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{client: newTransportErrorClient(t)}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRegisterNrtStatsFabricMapAction_Invoke_APIError exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRegisterNrtStatsFabricMapAction_Invoke_APIError(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_register_nrt_stats_fabric_map")
}

// TestRegisterNrtStatsFabricMapAction_Invoke_APIErrorReadBody exercises RegisterNrtStatsFabricMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRegisterNrtStatsFabricMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RegisterNrtStatsFabricMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := RegisterNrtStatsFabricMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
