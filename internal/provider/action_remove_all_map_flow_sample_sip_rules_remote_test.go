package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_Happy exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_NilClient exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_BuildError exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_SendError exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_APIError exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_sample_sip_rules")
}

// TestRemoveAllMapFlowSampleSipRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowSampleSipRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowSampleSipRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowSampleSipRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowSampleSipRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
