package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapApRuleAction_Invoke_Happy exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapApRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapApRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapApRuleAction_Invoke_NilClient exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapApRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapApRuleAction{}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapApRuleAction_Invoke_BuildError exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapApRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapApRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapApRuleAction_Invoke_SendError exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapApRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapApRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapApRuleAction_Invoke_APIError exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapApRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapApRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_ap_rule")
}

// TestUpdateMapApRuleAction_Invoke_APIErrorReadBody exercises UpdateMapApRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapApRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapApRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
