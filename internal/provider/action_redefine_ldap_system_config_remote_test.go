package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineLdapSystemConfigAction_Invoke_Happy exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineLdapSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineLdapSystemConfigAction_Invoke_NilClient exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineLdapSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineLdapSystemConfigAction_Invoke_BuildError exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineLdapSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineLdapSystemConfigAction_Invoke_SendError exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineLdapSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{client: newTransportErrorClient(t)}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineLdapSystemConfigAction_Invoke_APIError exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineLdapSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_ldap_system_config")
}

// TestRedefineLdapSystemConfigAction_Invoke_APIErrorReadBody exercises RedefineLdapSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineLdapSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineLdapSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineLdapSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
