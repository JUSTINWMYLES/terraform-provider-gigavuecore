package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_Happy exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_NilClient exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_BuildError exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_SendError exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_APIError exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_sample_overlap_rule")
}

// TestUpdateMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowSampleOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
