package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllTrafficFlowsAction_Invoke_Happy exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{client: newMockClientStatus(t, 202, "{}")}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllTrafficFlowsAction_Invoke_NilClient exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllTrafficFlowsAction_Invoke_BuildError exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllTrafficFlowsAction_Invoke_SendError exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllTrafficFlowsAction_Invoke_APIError exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_traffic_flows")
}

// TestDeleteAllTrafficFlowsAction_Invoke_APIErrorReadBody exercises DeleteAllTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
