package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_Happy exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_NilClient exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_BuildError exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_SendError exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_APIError exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_whitelist_rules")
}

// TestRemoveAllMapFlowWhitelistRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowWhitelistRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowWhitelistRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowWhitelistRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowWhitelistRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
