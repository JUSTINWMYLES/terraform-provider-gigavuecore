package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_Happy exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_Happy(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_NilClient exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_BuildError exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{client: newMalformedBaseURLClient(t)}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_SendError exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_SendError(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{client: newTransportErrorClient(t)}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_APIError exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_APIError(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_traffic_flows_deployed_draft")
}

// TestUpdateTrafficFlowsDeployedDraftAction_Invoke_APIErrorReadBody exercises UpdateTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateTrafficFlowsDeployedDraftAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateTrafficFlowsDeployedDraftAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
