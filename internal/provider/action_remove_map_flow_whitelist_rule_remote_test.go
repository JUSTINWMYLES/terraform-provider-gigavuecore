package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowWhitelistRuleAction_Invoke_Happy exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowWhitelistRuleAction_Invoke_NilClient exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowWhitelistRuleAction_Invoke_BuildError exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowWhitelistRuleAction_Invoke_SendError exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowWhitelistRuleAction_Invoke_APIError exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_whitelist_rule")
}

// TestRemoveMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowWhitelistRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
