package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowWhitelistRuleAction_Invoke_Happy exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowWhitelistRuleAction_Invoke_NilClient exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowWhitelistRuleAction_Invoke_BuildError exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowWhitelistRuleAction_Invoke_SendError exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowWhitelistRuleAction_Invoke_APIError exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_whitelist_rule")
}

// TestUpdateMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowWhitelistRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
