package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_Happy exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_NilClient exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_BuildError exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_SendError exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_APIError exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_sample_overlap_rules")
}

// TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowSampleOverlapRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowSampleOverlapRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowSampleOverlapRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowSampleOverlapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
