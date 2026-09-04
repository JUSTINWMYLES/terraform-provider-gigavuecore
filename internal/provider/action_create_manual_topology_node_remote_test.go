package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateManualTopologyNodeAction_Invoke_Happy exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateManualTopologyNodeAction_Invoke_Happy(t *testing.T) {
	r := &CreateManualTopologyNodeAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateManualTopologyNodeAction_Invoke_NilClient exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateManualTopologyNodeAction_Invoke_NilClient(t *testing.T) {
	r := &CreateManualTopologyNodeAction{}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateManualTopologyNodeAction_Invoke_BuildError exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateManualTopologyNodeAction_Invoke_BuildError(t *testing.T) {
	r := &CreateManualTopologyNodeAction{client: newMalformedBaseURLClient(t)}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateManualTopologyNodeAction_Invoke_SendError exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateManualTopologyNodeAction_Invoke_SendError(t *testing.T) {
	r := &CreateManualTopologyNodeAction{client: newTransportErrorClient(t)}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateManualTopologyNodeAction_Invoke_APIError exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateManualTopologyNodeAction_Invoke_APIError(t *testing.T) {
	r := &CreateManualTopologyNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_manual_topology_node")
}

// TestCreateManualTopologyNodeAction_Invoke_APIErrorReadBody exercises CreateManualTopologyNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateManualTopologyNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateManualTopologyNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateManualTopologyNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
