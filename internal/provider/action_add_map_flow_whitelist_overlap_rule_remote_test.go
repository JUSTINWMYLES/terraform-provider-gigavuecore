package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_Happy exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_NilClient exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_BuildError exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_SendError exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_APIError exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_whitelist_overlap_rule")
}

// TestAddMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowWhitelistOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowWhitelistOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowWhitelistOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowWhitelistOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
