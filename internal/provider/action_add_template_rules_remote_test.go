package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddTemplateRulesAction_Invoke_Happy exercises AddTemplateRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddTemplateRulesAction_Invoke_Happy(t *testing.T) {
	r := &AddTemplateRulesAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddTemplateRulesAction_Invoke_NilClient exercises AddTemplateRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddTemplateRulesAction_Invoke_NilClient(t *testing.T) {
	r := &AddTemplateRulesAction{}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddTemplateRulesAction_Invoke_BuildError exercises AddTemplateRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddTemplateRulesAction_Invoke_BuildError(t *testing.T) {
	r := &AddTemplateRulesAction{client: newMalformedBaseURLClient(t)}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddTemplateRulesAction_Invoke_SendError exercises AddTemplateRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddTemplateRulesAction_Invoke_SendError(t *testing.T) {
	r := &AddTemplateRulesAction{client: newTransportErrorClient(t)}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddTemplateRulesAction_Invoke_APIError exercises AddTemplateRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddTemplateRulesAction_Invoke_APIError(t *testing.T) {
	r := &AddTemplateRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_template_rules")
}

// TestAddTemplateRulesAction_Invoke_APIErrorReadBody exercises AddTemplateRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddTemplateRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddTemplateRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddTemplateRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
