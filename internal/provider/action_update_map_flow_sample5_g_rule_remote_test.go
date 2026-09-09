package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowSample5GRuleAction_Invoke_Happy exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowSample5GRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowSample5GRuleAction_Invoke_NilClient exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowSample5GRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowSample5GRuleAction_Invoke_BuildError exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowSample5GRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowSample5GRuleAction_Invoke_SendError exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowSample5GRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowSample5GRuleAction_Invoke_APIError exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowSample5GRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_sample5_g_rule")
}

// TestUpdateMapFlowSample5GRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowSample5GRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowSample5GRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowSample5GRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
