package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowSampleOverlapRuleAction_Invoke_Happy exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowSampleOverlapRuleAction_Invoke_NilClient exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowSampleOverlapRuleAction_Invoke_BuildError exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowSampleOverlapRuleAction_Invoke_SendError exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowSampleOverlapRuleAction_Invoke_APIError exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_sample_overlap_rule")
}

// TestAddMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowSampleOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowSampleOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowSampleOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowSampleOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
