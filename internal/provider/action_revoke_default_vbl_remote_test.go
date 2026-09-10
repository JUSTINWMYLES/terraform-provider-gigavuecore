package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevokeDefaultVblAction_Invoke_Happy exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevokeDefaultVblAction_Invoke_Happy(t *testing.T) {
	r := &RevokeDefaultVblAction{client: newMockClientStatus(t, 200, "{}")}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevokeDefaultVblAction_Invoke_NilClient exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevokeDefaultVblAction_Invoke_NilClient(t *testing.T) {
	r := &RevokeDefaultVblAction{}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevokeDefaultVblAction_Invoke_BuildError exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevokeDefaultVblAction_Invoke_BuildError(t *testing.T) {
	r := &RevokeDefaultVblAction{client: newMalformedBaseURLClient(t)}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevokeDefaultVblAction_Invoke_SendError exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevokeDefaultVblAction_Invoke_SendError(t *testing.T) {
	r := &RevokeDefaultVblAction{client: newTransportErrorClient(t)}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevokeDefaultVblAction_Invoke_APIError exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevokeDefaultVblAction_Invoke_APIError(t *testing.T) {
	r := &RevokeDefaultVblAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revoke_default_vbl")
}

// TestRevokeDefaultVblAction_Invoke_APIErrorReadBody exercises RevokeDefaultVblAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevokeDefaultVblAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevokeDefaultVblAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevokeDefaultVblActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
