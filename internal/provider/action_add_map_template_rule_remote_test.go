package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapTemplateRuleAction_Invoke_Happy exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapTemplateRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapTemplateRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapTemplateRuleAction_Invoke_NilClient exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapTemplateRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapTemplateRuleAction{}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapTemplateRuleAction_Invoke_BuildError exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapTemplateRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapTemplateRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapTemplateRuleAction_Invoke_SendError exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapTemplateRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapTemplateRuleAction{client: newTransportErrorClient(t)}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapTemplateRuleAction_Invoke_APIError exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapTemplateRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapTemplateRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_template_rule")
}

// TestAddMapTemplateRuleAction_Invoke_APIErrorReadBody exercises AddMapTemplateRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapTemplateRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapTemplateRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapTemplateRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
