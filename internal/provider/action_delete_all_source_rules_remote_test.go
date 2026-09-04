package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllSourceRulesAction_Invoke_Happy exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllSourceRulesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllSourceRulesAction{client: newMockClientStatus(t, 202, "{}")}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllSourceRulesAction_Invoke_NilClient exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllSourceRulesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllSourceRulesAction{}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllSourceRulesAction_Invoke_BuildError exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllSourceRulesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllSourceRulesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllSourceRulesAction_Invoke_SendError exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllSourceRulesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllSourceRulesAction{client: newTransportErrorClient(t)}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllSourceRulesAction_Invoke_APIError exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllSourceRulesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllSourceRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_source_rules")
}

// TestDeleteAllSourceRulesAction_Invoke_APIErrorReadBody exercises DeleteAllSourceRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllSourceRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllSourceRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
