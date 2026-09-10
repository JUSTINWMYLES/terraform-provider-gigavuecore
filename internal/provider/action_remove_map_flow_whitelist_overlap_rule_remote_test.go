package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_Happy exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_NilClient exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_BuildError exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_SendError exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_APIError exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_whitelist_overlap_rule")
}

// TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowWhitelistOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
