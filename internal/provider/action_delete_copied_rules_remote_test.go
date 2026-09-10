package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteCopiedRulesAction_Invoke_Happy exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteCopiedRulesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteCopiedRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteCopiedRulesAction_Invoke_NilClient exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteCopiedRulesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteCopiedRulesAction{}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteCopiedRulesAction_Invoke_BuildError exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteCopiedRulesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteCopiedRulesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteCopiedRulesAction_Invoke_SendError exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteCopiedRulesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteCopiedRulesAction{client: newTransportErrorClient(t)}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteCopiedRulesAction_Invoke_APIError exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteCopiedRulesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteCopiedRulesAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_copied_rules")
}

// TestDeleteCopiedRulesAction_Invoke_APIErrorReadBody exercises DeleteCopiedRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteCopiedRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteCopiedRulesAction{client: newMockClientReadErrorBody(t, 500)}
	m := DeleteCopiedRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
