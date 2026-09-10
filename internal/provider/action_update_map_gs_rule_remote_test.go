package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapGsRuleAction_Invoke_Happy exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapGsRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapGsRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapGsRuleAction_Invoke_NilClient exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapGsRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapGsRuleAction{}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapGsRuleAction_Invoke_BuildError exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapGsRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapGsRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapGsRuleAction_Invoke_SendError exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapGsRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapGsRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapGsRuleAction_Invoke_APIError exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapGsRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapGsRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_gs_rule")
}

// TestUpdateMapGsRuleAction_Invoke_APIErrorReadBody exercises UpdateMapGsRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapGsRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapGsRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
