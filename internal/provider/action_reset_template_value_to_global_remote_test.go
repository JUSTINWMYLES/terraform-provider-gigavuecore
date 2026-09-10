package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetTemplateValueToGlobalAction_Invoke_Happy exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetTemplateValueToGlobalAction_Invoke_Happy(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{client: newMockClientStatus(t, 200, "{}")}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetTemplateValueToGlobalAction_Invoke_NilClient exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetTemplateValueToGlobalAction_Invoke_NilClient(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetTemplateValueToGlobalAction_Invoke_BuildError exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetTemplateValueToGlobalAction_Invoke_BuildError(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{client: newMalformedBaseURLClient(t)}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetTemplateValueToGlobalAction_Invoke_SendError exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetTemplateValueToGlobalAction_Invoke_SendError(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{client: newTransportErrorClient(t)}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetTemplateValueToGlobalAction_Invoke_APIError exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetTemplateValueToGlobalAction_Invoke_APIError(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_template_value_to_global")
}

// TestResetTemplateValueToGlobalAction_Invoke_APIErrorReadBody exercises ResetTemplateValueToGlobalAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetTemplateValueToGlobalAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetTemplateValueToGlobalAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetTemplateValueToGlobalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
