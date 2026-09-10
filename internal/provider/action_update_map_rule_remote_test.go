package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapRuleAction_Invoke_Happy exercises UpdateMapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapRuleAction_Invoke_NilClient exercises UpdateMapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapRuleAction{}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapRuleAction_Invoke_BuildError exercises UpdateMapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapRuleAction_Invoke_SendError exercises UpdateMapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapRuleAction_Invoke_APIError exercises UpdateMapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_rule")
}

// TestUpdateMapRuleAction_Invoke_APIErrorReadBody exercises UpdateMapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
