package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapRuleAction_Invoke_Happy exercises RemoveMapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapRuleAction_Invoke_NilClient exercises RemoveMapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapRuleAction{}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapRuleAction_Invoke_BuildError exercises RemoveMapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapRuleAction_Invoke_SendError exercises RemoveMapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapRuleAction_Invoke_APIError exercises RemoveMapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_rule")
}

// TestRemoveMapRuleAction_Invoke_APIErrorReadBody exercises RemoveMapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
