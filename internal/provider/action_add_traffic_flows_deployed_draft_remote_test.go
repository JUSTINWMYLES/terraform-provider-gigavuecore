package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddTrafficFlowsDeployedDraftAction_Invoke_Happy exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_Happy(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{client: newMockClientStatus(t, 200, "{}")}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddTrafficFlowsDeployedDraftAction_Invoke_NilClient exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_NilClient(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddTrafficFlowsDeployedDraftAction_Invoke_BuildError exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_BuildError(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{client: newMalformedBaseURLClient(t)}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddTrafficFlowsDeployedDraftAction_Invoke_SendError exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_SendError(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{client: newTransportErrorClient(t)}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddTrafficFlowsDeployedDraftAction_Invoke_APIError exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_APIError(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_traffic_flows_deployed_draft")
}

// TestAddTrafficFlowsDeployedDraftAction_Invoke_APIErrorReadBody exercises AddTrafficFlowsDeployedDraftAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddTrafficFlowsDeployedDraftAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddTrafficFlowsDeployedDraftAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddTrafficFlowsDeployedDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
