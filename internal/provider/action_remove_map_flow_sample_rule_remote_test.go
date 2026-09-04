package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowSampleRuleAction_Invoke_Happy exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowSampleRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowSampleRuleAction_Invoke_NilClient exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowSampleRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowSampleRuleAction_Invoke_BuildError exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowSampleRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowSampleRuleAction_Invoke_SendError exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowSampleRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowSampleRuleAction_Invoke_APIError exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowSampleRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_sample_rule")
}

// TestRemoveMapFlowSampleRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowSampleRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowSampleRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
