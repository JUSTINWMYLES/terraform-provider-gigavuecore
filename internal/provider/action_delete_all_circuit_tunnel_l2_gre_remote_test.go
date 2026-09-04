package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_Happy exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_NilClient exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_BuildError exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_SendError exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{client: newTransportErrorClient(t)}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_APIError exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_circuit_tunnel_l2_gre")
}

// TestDeleteAllCircuitTunnelL2GreAction_Invoke_APIErrorReadBody exercises DeleteAllCircuitTunnelL2GreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllCircuitTunnelL2GreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllCircuitTunnelL2GreAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllCircuitTunnelL2GreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
