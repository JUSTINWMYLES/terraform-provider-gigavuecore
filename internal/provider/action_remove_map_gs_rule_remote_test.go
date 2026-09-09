package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapGsRuleAction_Invoke_Happy exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapGsRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapGsRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapGsRuleAction_Invoke_NilClient exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapGsRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapGsRuleAction{}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapGsRuleAction_Invoke_BuildError exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapGsRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapGsRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapGsRuleAction_Invoke_SendError exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapGsRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapGsRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapGsRuleAction_Invoke_APIError exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapGsRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapGsRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_gs_rule")
}

// TestRemoveMapGsRuleAction_Invoke_APIErrorReadBody exercises RemoveMapGsRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapGsRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapGsRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
