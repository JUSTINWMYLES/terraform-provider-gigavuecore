package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllFlowRulesAction_Invoke_Happy exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllFlowRulesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllFlowRulesAction{client: newMockClientStatus(t, 202, "{}")}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllFlowRulesAction_Invoke_NilClient exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllFlowRulesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllFlowRulesAction{}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllFlowRulesAction_Invoke_BuildError exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllFlowRulesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllFlowRulesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllFlowRulesAction_Invoke_SendError exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllFlowRulesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllFlowRulesAction{client: newTransportErrorClient(t)}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllFlowRulesAction_Invoke_APIError exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllFlowRulesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllFlowRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_flow_rules")
}

// TestDeleteAllFlowRulesAction_Invoke_APIErrorReadBody exercises DeleteAllFlowRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllFlowRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllFlowRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
