package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSecurityConfigAction_Invoke_Happy exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSecurityConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSecurityConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSecurityConfigAction_Invoke_NilClient exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSecurityConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSecurityConfigAction{}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSecurityConfigAction_Invoke_BuildError exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSecurityConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSecurityConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSecurityConfigAction_Invoke_SendError exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSecurityConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSecurityConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSecurityConfigAction_Invoke_APIError exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSecurityConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSecurityConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_security_config")
}

// TestUpdateSecurityConfigAction_Invoke_APIErrorReadBody exercises UpdateSecurityConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSecurityConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSecurityConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
