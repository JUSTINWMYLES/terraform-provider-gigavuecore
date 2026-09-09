package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteManualTopologyEntitiesAction_Invoke_Happy exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteManualTopologyEntitiesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{client: newMockClientStatus(t, 207, "{}")}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteManualTopologyEntitiesAction_Invoke_NilClient exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteManualTopologyEntitiesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteManualTopologyEntitiesAction_Invoke_BuildError exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteManualTopologyEntitiesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteManualTopologyEntitiesAction_Invoke_SendError exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteManualTopologyEntitiesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{client: newTransportErrorClient(t)}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteManualTopologyEntitiesAction_Invoke_APIError exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteManualTopologyEntitiesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_manual_topology_entities")
}

// TestDeleteManualTopologyEntitiesAction_Invoke_APIErrorReadBody exercises DeleteManualTopologyEntitiesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteManualTopologyEntitiesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteManualTopologyEntitiesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteManualTopologyEntitiesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
