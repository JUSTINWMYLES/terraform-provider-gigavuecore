package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapRulesAction_Invoke_Happy exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapRulesAction_Invoke_NilClient exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapRulesAction{}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapRulesAction_Invoke_BuildError exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapRulesAction_Invoke_SendError exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapRulesAction_Invoke_APIError exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_rules")
}

// TestRemoveAllMapRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
