package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapRulesFromTemplateAction_Invoke_Happy exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapRulesFromTemplateAction_Invoke_Happy(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapRulesFromTemplateAction_Invoke_NilClient exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapRulesFromTemplateAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapRulesFromTemplateAction_Invoke_BuildError exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapRulesFromTemplateAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{client: newMalformedBaseURLClient(t)}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapRulesFromTemplateAction_Invoke_SendError exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapRulesFromTemplateAction_Invoke_SendError(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{client: newTransportErrorClient(t)}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapRulesFromTemplateAction_Invoke_APIError exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapRulesFromTemplateAction_Invoke_APIError(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_rules_from_template")
}

// TestAddMapRulesFromTemplateAction_Invoke_APIErrorReadBody exercises AddMapRulesFromTemplateAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapRulesFromTemplateAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapRulesFromTemplateAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapRulesFromTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
