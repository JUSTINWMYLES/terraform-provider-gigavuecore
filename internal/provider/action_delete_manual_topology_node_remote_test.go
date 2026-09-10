package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteManualTopologyNodeAction_Invoke_Happy exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteManualTopologyNodeAction_Invoke_Happy(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteManualTopologyNodeAction_Invoke_NilClient exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteManualTopologyNodeAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteManualTopologyNodeAction_Invoke_BuildError exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteManualTopologyNodeAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{client: newMalformedBaseURLClient(t)}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteManualTopologyNodeAction_Invoke_SendError exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteManualTopologyNodeAction_Invoke_SendError(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{client: newTransportErrorClient(t)}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteManualTopologyNodeAction_Invoke_APIError exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteManualTopologyNodeAction_Invoke_APIError(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_manual_topology_node")
}

// TestDeleteManualTopologyNodeAction_Invoke_APIErrorReadBody exercises DeleteManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteManualTopologyNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteManualTopologyNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
