package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowSampleRulesAction_Invoke_Happy exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowSampleRulesAction_Invoke_NilClient exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowSampleRulesAction_Invoke_BuildError exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowSampleRulesAction_Invoke_SendError exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowSampleRulesAction_Invoke_APIError exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_sample_rules")
}

// TestRemoveAllMapFlowSampleRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowSampleRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowSampleRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowSampleRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowSampleRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
