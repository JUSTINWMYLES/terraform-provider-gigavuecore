package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_Happy exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_NilClient exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_BuildError exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_SendError exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_APIError exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_flow_sample_diameter_rule")
}

// TestUpdateMapFlowSampleDiameterRuleAction_Invoke_APIErrorReadBody exercises UpdateMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapFlowSampleDiameterRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapFlowSampleDiameterRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
