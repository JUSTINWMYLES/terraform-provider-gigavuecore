package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_Happy exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_NilClient exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_BuildError exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_SendError exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{client: newTransportErrorClient(t)}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_APIError exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_circuit_tunnel_vxlan_groups")
}

// TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_APIErrorReadBody exercises DeleteCircuitTunnelVxlanGroupsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteCircuitTunnelVxlanGroupsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteCircuitTunnelVxlanGroupsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteCircuitTunnelVxlanGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
