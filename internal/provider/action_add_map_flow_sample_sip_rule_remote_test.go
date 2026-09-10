package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowSampleSipRuleAction_Invoke_Happy exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowSampleSipRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowSampleSipRuleAction_Invoke_NilClient exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowSampleSipRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowSampleSipRuleAction_Invoke_BuildError exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowSampleSipRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowSampleSipRuleAction_Invoke_SendError exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowSampleSipRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowSampleSipRuleAction_Invoke_APIError exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowSampleSipRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_sample_sip_rule")
}

// TestAddMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowSampleSipRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
