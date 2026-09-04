package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_Happy exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_NilClient exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_BuildError exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_SendError exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_APIError exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_whitelist_overlap_rule")
}

// TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowWhitelistOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
