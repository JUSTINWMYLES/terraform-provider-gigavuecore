package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateInlineSslConfigAction_Invoke_Happy exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateInlineSslConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateInlineSslConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateInlineSslConfigAction_Invoke_NilClient exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateInlineSslConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateInlineSslConfigAction{}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateInlineSslConfigAction_Invoke_BuildError exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateInlineSslConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateInlineSslConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateInlineSslConfigAction_Invoke_SendError exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateInlineSslConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateInlineSslConfigAction{client: newTransportErrorClient(t)}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateInlineSslConfigAction_Invoke_APIError exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateInlineSslConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateInlineSslConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_inline_ssl_config")
}

// TestUpdateInlineSslConfigAction_Invoke_APIErrorReadBody exercises UpdateInlineSslConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateInlineSslConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateInlineSslConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateInlineSslConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
