package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetTopologyVizAction_Invoke_Happy exercises ResetTopologyVizAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetTopologyVizAction_Invoke_Happy(t *testing.T) {
	r := &ResetTopologyVizAction{client: newMockClientStatus(t, 200, "{}")}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetTopologyVizAction_Invoke_NilClient exercises ResetTopologyVizAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetTopologyVizAction_Invoke_NilClient(t *testing.T) {
	r := &ResetTopologyVizAction{}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetTopologyVizAction_Invoke_BuildError exercises ResetTopologyVizAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetTopologyVizAction_Invoke_BuildError(t *testing.T) {
	r := &ResetTopologyVizAction{client: newMalformedBaseURLClient(t)}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetTopologyVizAction_Invoke_SendError exercises ResetTopologyVizAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetTopologyVizAction_Invoke_SendError(t *testing.T) {
	r := &ResetTopologyVizAction{client: newTransportErrorClient(t)}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetTopologyVizAction_Invoke_APIError exercises ResetTopologyVizAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetTopologyVizAction_Invoke_APIError(t *testing.T) {
	r := &ResetTopologyVizAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_topology_viz")
}

// TestResetTopologyVizAction_Invoke_APIErrorReadBody exercises ResetTopologyVizAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetTopologyVizAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetTopologyVizAction{client: newMockClientReadErrorBody(t, 500)}
	m := ResetTopologyVizActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
