package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_Happy exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_NilClient exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_BuildError exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_SendError exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_APIError exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_sample_overlap_rule")
}

// TestRemoveMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowSampleOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
