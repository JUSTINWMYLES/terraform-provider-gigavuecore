package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteManualTopologyAction_Invoke_Happy exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteManualTopologyAction_Invoke_Happy(t *testing.T) {
	r := &DeleteManualTopologyAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteManualTopologyAction_Invoke_NilClient exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteManualTopologyAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteManualTopologyAction{}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteManualTopologyAction_Invoke_BuildError exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteManualTopologyAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteManualTopologyAction{client: newMalformedBaseURLClient(t)}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteManualTopologyAction_Invoke_SendError exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteManualTopologyAction_Invoke_SendError(t *testing.T) {
	r := &DeleteManualTopologyAction{client: newTransportErrorClient(t)}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteManualTopologyAction_Invoke_APIError exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteManualTopologyAction_Invoke_APIError(t *testing.T) {
	r := &DeleteManualTopologyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_manual_topology")
}

// TestDeleteManualTopologyAction_Invoke_APIErrorReadBody exercises DeleteManualTopologyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteManualTopologyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteManualTopologyAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteManualTopologyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
