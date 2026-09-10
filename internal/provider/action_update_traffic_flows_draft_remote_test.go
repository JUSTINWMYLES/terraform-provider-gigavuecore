package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateTrafficFlowsDraftAction_Invoke_Happy exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateTrafficFlowsDraftAction_Invoke_Happy(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateTrafficFlowsDraftAction_Invoke_NilClient exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateTrafficFlowsDraftAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateTrafficFlowsDraftAction_Invoke_BuildError exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateTrafficFlowsDraftAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{client: newMalformedBaseURLClient(t)}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateTrafficFlowsDraftAction_Invoke_SendError exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateTrafficFlowsDraftAction_Invoke_SendError(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{client: newTransportErrorClient(t)}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateTrafficFlowsDraftAction_Invoke_APIError exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateTrafficFlowsDraftAction_Invoke_APIError(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_traffic_flows_draft")
}

// TestUpdateTrafficFlowsDraftAction_Invoke_APIErrorReadBody exercises UpdateTrafficFlowsDraftAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateTrafficFlowsDraftAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateTrafficFlowsDraftAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateTrafficFlowsDraftActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
