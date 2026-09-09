package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowRuleAction_Invoke_Happy exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowRuleAction_Invoke_NilClient exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowRuleAction{}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowRuleAction_Invoke_BuildError exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowRuleAction_Invoke_SendError exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowRuleAction_Invoke_APIError exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_rule")
}

// TestRemoveMapFlowRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
