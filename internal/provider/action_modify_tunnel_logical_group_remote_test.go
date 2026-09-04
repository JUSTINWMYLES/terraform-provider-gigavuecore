package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestModifyTunnelLogicalGroupAction_Invoke_Happy exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestModifyTunnelLogicalGroupAction_Invoke_Happy(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{client: newMockClientStatus(t, 200, "{}")}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestModifyTunnelLogicalGroupAction_Invoke_NilClient exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestModifyTunnelLogicalGroupAction_Invoke_NilClient(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestModifyTunnelLogicalGroupAction_Invoke_BuildError exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestModifyTunnelLogicalGroupAction_Invoke_BuildError(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{client: newMalformedBaseURLClient(t)}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestModifyTunnelLogicalGroupAction_Invoke_SendError exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestModifyTunnelLogicalGroupAction_Invoke_SendError(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{client: newTransportErrorClient(t)}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestModifyTunnelLogicalGroupAction_Invoke_APIError exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestModifyTunnelLogicalGroupAction_Invoke_APIError(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_modify_tunnel_logical_group")
}

// TestModifyTunnelLogicalGroupAction_Invoke_APIErrorReadBody exercises ModifyTunnelLogicalGroupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestModifyTunnelLogicalGroupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ModifyTunnelLogicalGroupAction{client: newMockClientReadErrorBody(t, 501)}
	m := ModifyTunnelLogicalGroupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
