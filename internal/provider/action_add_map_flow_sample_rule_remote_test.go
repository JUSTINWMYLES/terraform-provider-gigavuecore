package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowSampleRuleAction_Invoke_Happy exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowSampleRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowSampleRuleAction_Invoke_NilClient exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowSampleRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowSampleRuleAction_Invoke_BuildError exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowSampleRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowSampleRuleAction_Invoke_SendError exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowSampleRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowSampleRuleAction_Invoke_APIError exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowSampleRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_sample_rule")
}

// TestAddMapFlowSampleRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowSampleRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowSampleRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowSampleRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowSampleRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
