package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowSample5GRuleAction_Invoke_Happy exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowSample5GRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowSample5GRuleAction_Invoke_NilClient exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowSample5GRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowSample5GRuleAction_Invoke_BuildError exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowSample5GRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowSample5GRuleAction_Invoke_SendError exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowSample5GRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowSample5GRuleAction_Invoke_APIError exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowSample5GRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_sample5_g_rule")
}

// TestRemoveMapFlowSample5GRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowSample5GRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowSample5GRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowSample5GRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowSample5GRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
