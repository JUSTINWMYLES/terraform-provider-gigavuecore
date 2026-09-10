package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateTopologyVizConfigsAction_Invoke_Happy exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateTopologyVizConfigsAction_Invoke_Happy(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{client: newMockClientStatus(t, 204, "{}")}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateTopologyVizConfigsAction_Invoke_NilClient exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateTopologyVizConfigsAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateTopologyVizConfigsAction_Invoke_BuildError exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateTopologyVizConfigsAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{client: newMalformedBaseURLClient(t)}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateTopologyVizConfigsAction_Invoke_SendError exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateTopologyVizConfigsAction_Invoke_SendError(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{client: newTransportErrorClient(t)}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateTopologyVizConfigsAction_Invoke_APIError exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateTopologyVizConfigsAction_Invoke_APIError(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_topology_viz_configs")
}

// TestUpdateTopologyVizConfigsAction_Invoke_APIErrorReadBody exercises UpdateTopologyVizConfigsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateTopologyVizConfigsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateTopologyVizConfigsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateTopologyVizConfigsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
