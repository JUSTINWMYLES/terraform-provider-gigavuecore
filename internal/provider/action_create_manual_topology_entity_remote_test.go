package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateManualTopologyEntityAction_Invoke_Happy exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateManualTopologyEntityAction_Invoke_Happy(t *testing.T) {
	r := &CreateManualTopologyEntityAction{client: newMockClientStatus(t, 207, "{}")}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateManualTopologyEntityAction_Invoke_NilClient exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateManualTopologyEntityAction_Invoke_NilClient(t *testing.T) {
	r := &CreateManualTopologyEntityAction{}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateManualTopologyEntityAction_Invoke_BuildError exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateManualTopologyEntityAction_Invoke_BuildError(t *testing.T) {
	r := &CreateManualTopologyEntityAction{client: newMalformedBaseURLClient(t)}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateManualTopologyEntityAction_Invoke_SendError exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateManualTopologyEntityAction_Invoke_SendError(t *testing.T) {
	r := &CreateManualTopologyEntityAction{client: newTransportErrorClient(t)}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateManualTopologyEntityAction_Invoke_APIError exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateManualTopologyEntityAction_Invoke_APIError(t *testing.T) {
	r := &CreateManualTopologyEntityAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_manual_topology_entity")
}

// TestCreateManualTopologyEntityAction_Invoke_APIErrorReadBody exercises CreateManualTopologyEntityAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateManualTopologyEntityAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateManualTopologyEntityAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateManualTopologyEntityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
