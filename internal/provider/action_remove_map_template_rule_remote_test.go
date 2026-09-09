package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapTemplateRuleAction_Invoke_Happy exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapTemplateRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapTemplateRuleAction_Invoke_NilClient exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapTemplateRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapTemplateRuleAction_Invoke_BuildError exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapTemplateRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapTemplateRuleAction_Invoke_SendError exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapTemplateRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapTemplateRuleAction_Invoke_APIError exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapTemplateRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_template_rule")
}

// TestRemoveMapTemplateRuleAction_Invoke_APIErrorReadBody exercises RemoveMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapTemplateRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapTemplateRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
