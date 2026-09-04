package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowSampleRuleAction_Invoke_Happy exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowSampleRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowSampleRuleAction_Invoke_NilClient exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowSampleRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowSampleRuleAction_Invoke_BuildError exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowSampleRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowSampleRuleAction_Invoke_SendError exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowSampleRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowSampleRuleAction_Invoke_APIError exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowSampleRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_sample_rule")
}

// TestUpdateMapFlowSampleRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowSampleRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowSampleRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
