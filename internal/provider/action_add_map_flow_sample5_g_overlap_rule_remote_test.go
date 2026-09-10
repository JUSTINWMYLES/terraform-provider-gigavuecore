package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_Happy exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_NilClient exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_BuildError exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_SendError exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_APIError exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_sample5_g_overlap_rule")
}

// TestAddMapFlowSample5GOverlapRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowSample5GOverlapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowSample5GOverlapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowSample5GOverlapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowSample5GOverlapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
