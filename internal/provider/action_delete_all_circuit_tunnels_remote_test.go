package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllCircuitTunnelsAction_Invoke_Happy exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllCircuitTunnelsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllCircuitTunnelsAction_Invoke_NilClient exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllCircuitTunnelsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllCircuitTunnelsAction_Invoke_BuildError exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllCircuitTunnelsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllCircuitTunnelsAction_Invoke_SendError exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllCircuitTunnelsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{client: newTransportErrorClient(t)}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllCircuitTunnelsAction_Invoke_APIError exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllCircuitTunnelsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_circuit_tunnels")
}

// TestDeleteAllCircuitTunnelsAction_Invoke_APIErrorReadBody exercises DeleteAllCircuitTunnelsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllCircuitTunnelsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllCircuitTunnelsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllCircuitTunnelsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
