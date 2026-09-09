package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowRuleAction_Invoke_Happy exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowRuleAction_Invoke_NilClient exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowRuleAction{}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowRuleAction_Invoke_BuildError exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowRuleAction_Invoke_SendError exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowRuleAction_Invoke_APIError exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_rule")
}

// TestUpdateMapFlowRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
