package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_Happy exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_NilClient exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_BuildError exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_SendError exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_APIError exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_sample5_g_rules")
}

// TestRemoveAllMapFlowSample5GRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowSample5GRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowSample5GRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowSample5GRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowSample5GRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
