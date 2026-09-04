package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowWhitelistRuleAction_Invoke_Happy exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowWhitelistRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowWhitelistRuleAction_Invoke_NilClient exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowWhitelistRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowWhitelistRuleAction_Invoke_BuildError exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowWhitelistRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowWhitelistRuleAction_Invoke_SendError exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowWhitelistRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowWhitelistRuleAction_Invoke_APIError exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowWhitelistRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_whitelist_rule")
}

// TestAddMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowWhitelistRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowWhitelistRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowWhitelistRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowWhitelistRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
