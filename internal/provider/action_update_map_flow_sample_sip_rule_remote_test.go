package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowSampleSipRuleAction_Invoke_Happy exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowSampleSipRuleAction_Invoke_NilClient exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowSampleSipRuleAction_Invoke_BuildError exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowSampleSipRuleAction_Invoke_SendError exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowSampleSipRuleAction_Invoke_APIError exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_sample_sip_rule")
}

// TestUpdateMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowSampleSipRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
