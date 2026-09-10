package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevokeAllApiTokensAction_Invoke_Happy exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevokeAllApiTokensAction_Invoke_Happy(t *testing.T) {
	r := &RevokeAllApiTokensAction{client: newMockClientStatus(t, 200, "{}")}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevokeAllApiTokensAction_Invoke_NilClient exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevokeAllApiTokensAction_Invoke_NilClient(t *testing.T) {
	r := &RevokeAllApiTokensAction{}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevokeAllApiTokensAction_Invoke_BuildError exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevokeAllApiTokensAction_Invoke_BuildError(t *testing.T) {
	r := &RevokeAllApiTokensAction{client: newMalformedBaseURLClient(t)}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevokeAllApiTokensAction_Invoke_SendError exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevokeAllApiTokensAction_Invoke_SendError(t *testing.T) {
	r := &RevokeAllApiTokensAction{client: newTransportErrorClient(t)}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevokeAllApiTokensAction_Invoke_APIError exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevokeAllApiTokensAction_Invoke_APIError(t *testing.T) {
	r := &RevokeAllApiTokensAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revoke_all_api_tokens")
}

// TestRevokeAllApiTokensAction_Invoke_APIErrorReadBody exercises RevokeAllApiTokensAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevokeAllApiTokensAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevokeAllApiTokensAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevokeAllApiTokensActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
