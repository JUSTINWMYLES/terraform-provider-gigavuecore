package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateManualTopologyNodeAction_Invoke_Happy exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateManualTopologyNodeAction_Invoke_Happy(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateManualTopologyNodeAction_Invoke_NilClient exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateManualTopologyNodeAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateManualTopologyNodeAction_Invoke_BuildError exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateManualTopologyNodeAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{client: newMalformedBaseURLClient(t)}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateManualTopologyNodeAction_Invoke_SendError exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateManualTopologyNodeAction_Invoke_SendError(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{client: newTransportErrorClient(t)}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateManualTopologyNodeAction_Invoke_APIError exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateManualTopologyNodeAction_Invoke_APIError(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_manual_topology_node")
}

// TestUpdateManualTopologyNodeAction_Invoke_APIErrorReadBody exercises UpdateManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateManualTopologyNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateManualTopologyNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
