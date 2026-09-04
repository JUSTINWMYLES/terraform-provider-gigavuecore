package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowSampleSipRuleAction_Invoke_Happy exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowSampleSipRuleAction_Invoke_NilClient exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowSampleSipRuleAction_Invoke_BuildError exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowSampleSipRuleAction_Invoke_SendError exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowSampleSipRuleAction_Invoke_APIError exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_sample_sip_rule")
}

// TestRemoveMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowSampleSipRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowSampleSipRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowSampleSipRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowSampleSipRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
