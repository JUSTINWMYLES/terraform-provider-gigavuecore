package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestChangeTunnelAuthModeAction_Invoke_Happy exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestChangeTunnelAuthModeAction_Invoke_Happy(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{client: newMockClientStatus(t, 200, "{}")}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestChangeTunnelAuthModeAction_Invoke_NilClient exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestChangeTunnelAuthModeAction_Invoke_NilClient(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestChangeTunnelAuthModeAction_Invoke_BuildError exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestChangeTunnelAuthModeAction_Invoke_BuildError(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{client: newMalformedBaseURLClient(t)}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestChangeTunnelAuthModeAction_Invoke_SendError exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestChangeTunnelAuthModeAction_Invoke_SendError(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{client: newTransportErrorClient(t)}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestChangeTunnelAuthModeAction_Invoke_APIError exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestChangeTunnelAuthModeAction_Invoke_APIError(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_change_tunnel_auth_mode")
}

// TestChangeTunnelAuthModeAction_Invoke_APIErrorReadBody exercises ChangeTunnelAuthModeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestChangeTunnelAuthModeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ChangeTunnelAuthModeAction{client: newMockClientReadErrorBody(t, 501)}
	m := ChangeTunnelAuthModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
