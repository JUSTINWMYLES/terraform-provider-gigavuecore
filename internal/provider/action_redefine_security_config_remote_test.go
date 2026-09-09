package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSecurityConfigAction_Invoke_Happy exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSecurityConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSecurityConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSecurityConfigAction_Invoke_NilClient exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSecurityConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSecurityConfigAction{}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSecurityConfigAction_Invoke_BuildError exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSecurityConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSecurityConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSecurityConfigAction_Invoke_SendError exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSecurityConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSecurityConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSecurityConfigAction_Invoke_APIError exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSecurityConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSecurityConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_security_config")
}

// TestRedefineSecurityConfigAction_Invoke_APIErrorReadBody exercises RedefineSecurityConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSecurityConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSecurityConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSecurityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
