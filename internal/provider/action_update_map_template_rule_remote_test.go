package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapTemplateRuleAction_Invoke_Happy exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapTemplateRuleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapTemplateRuleAction_Invoke_NilClient exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapTemplateRuleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapTemplateRuleAction_Invoke_BuildError exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapTemplateRuleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapTemplateRuleAction_Invoke_SendError exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapTemplateRuleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{client: newTransportErrorClient(t)}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapTemplateRuleAction_Invoke_APIError exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapTemplateRuleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_template_rule")
}

// TestUpdateMapTemplateRuleAction_Invoke_APIErrorReadBody exercises UpdateMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapTemplateRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapTemplateRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
