package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapFlowRuleAction_Invoke_Happy exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapFlowRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapFlowRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapFlowRuleAction_Invoke_NilClient exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapFlowRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapFlowRuleAction{}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapFlowRuleAction_Invoke_BuildError exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapFlowRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapFlowRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapFlowRuleAction_Invoke_SendError exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapFlowRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapFlowRuleAction{client: newTransportErrorClient(t)}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapFlowRuleAction_Invoke_APIError exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapFlowRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapFlowRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_flow_rule")
}

// TestAddMapFlowRuleAction_Invoke_APIErrorReadBody exercises AddMapFlowRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapFlowRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapFlowRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapFlowRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
