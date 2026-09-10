package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapTemplateRulesAction_Invoke_Happy exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapTemplateRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapTemplateRulesAction_Invoke_NilClient exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapTemplateRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapTemplateRulesAction_Invoke_BuildError exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapTemplateRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapTemplateRulesAction_Invoke_SendError exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapTemplateRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapTemplateRulesAction_Invoke_APIError exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapTemplateRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_template_rules")
}

// TestRemoveAllMapTemplateRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapTemplateRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapTemplateRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapTemplateRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
