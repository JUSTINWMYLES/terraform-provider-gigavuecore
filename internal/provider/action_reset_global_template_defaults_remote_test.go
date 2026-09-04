package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetGlobalTemplateDefaultsAction_Invoke_Happy exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetGlobalTemplateDefaultsAction_Invoke_Happy(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{client: newMockClientStatus(t, 200, "{}")}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetGlobalTemplateDefaultsAction_Invoke_NilClient exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetGlobalTemplateDefaultsAction_Invoke_NilClient(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetGlobalTemplateDefaultsAction_Invoke_BuildError exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetGlobalTemplateDefaultsAction_Invoke_BuildError(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{client: newMalformedBaseURLClient(t)}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetGlobalTemplateDefaultsAction_Invoke_SendError exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetGlobalTemplateDefaultsAction_Invoke_SendError(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{client: newTransportErrorClient(t)}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetGlobalTemplateDefaultsAction_Invoke_APIError exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetGlobalTemplateDefaultsAction_Invoke_APIError(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_global_template_defaults")
}

// TestResetGlobalTemplateDefaultsAction_Invoke_APIErrorReadBody exercises ResetGlobalTemplateDefaultsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetGlobalTemplateDefaultsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetGlobalTemplateDefaultsAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetGlobalTemplateDefaultsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
