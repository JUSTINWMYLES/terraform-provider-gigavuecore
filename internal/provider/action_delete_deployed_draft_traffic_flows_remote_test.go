package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_Happy exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_NilClient exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_BuildError exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_SendError exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_APIError exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_deployed_draft_traffic_flows")
}

// TestDeleteDeployedDraftTrafficFlowsAction_Invoke_APIErrorReadBody exercises DeleteDeployedDraftTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteDeployedDraftTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteDeployedDraftTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteDeployedDraftTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
