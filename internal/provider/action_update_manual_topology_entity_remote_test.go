package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateManualTopologyEntityAction_Invoke_Happy exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateManualTopologyEntityAction_Invoke_Happy(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{client: newMockClientStatus(t, 207, "{}")}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateManualTopologyEntityAction_Invoke_NilClient exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateManualTopologyEntityAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateManualTopologyEntityAction_Invoke_BuildError exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateManualTopologyEntityAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{client: newMalformedBaseURLClient(t)}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateManualTopologyEntityAction_Invoke_SendError exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateManualTopologyEntityAction_Invoke_SendError(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{client: newTransportErrorClient(t)}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateManualTopologyEntityAction_Invoke_APIError exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateManualTopologyEntityAction_Invoke_APIError(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_manual_topology_entity")
}

// TestUpdateManualTopologyEntityAction_Invoke_APIErrorReadBody exercises UpdateManualTopologyEntityAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateManualTopologyEntityAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateManualTopologyEntityAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
