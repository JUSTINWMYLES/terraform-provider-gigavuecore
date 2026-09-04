package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_Happy exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_Happy(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{client: newMockClientStatus(t, 200, "{}")}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_NilClient exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_NilClient(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_BuildError exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_BuildError(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{client: newMalformedBaseURLClient(t)}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_SendError exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_SendError(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{client: newTransportErrorClient(t)}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_APIError exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_APIError(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revoke_api_tokens_for_privilege_users")
}

// TestRevokeApiTokensForPrivilegeUsersAction_Invoke_APIErrorReadBody exercises RevokeApiTokensForPrivilegeUsersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevokeApiTokensForPrivilegeUsersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevokeApiTokensForPrivilegeUsersAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevokeApiTokensForPrivilegeUsersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
