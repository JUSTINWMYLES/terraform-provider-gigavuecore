package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowRulesAction_Invoke_Happy exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowRulesAction_Invoke_NilClient exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowRulesAction_Invoke_BuildError exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowRulesAction_Invoke_SendError exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowRulesAction_Invoke_APIError exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_rules")
}

// TestRemoveAllMapFlowRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
