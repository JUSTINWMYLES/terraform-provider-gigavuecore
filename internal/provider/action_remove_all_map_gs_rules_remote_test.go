package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapGsRulesAction_Invoke_Happy exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapGsRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapGsRulesAction_Invoke_NilClient exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapGsRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapGsRulesAction_Invoke_BuildError exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapGsRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapGsRulesAction_Invoke_SendError exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapGsRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapGsRulesAction_Invoke_APIError exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapGsRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_gs_rules")
}

// TestRemoveAllMapGsRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapGsRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapGsRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapGsRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapGsRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
