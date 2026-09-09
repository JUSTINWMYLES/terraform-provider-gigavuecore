package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllTunnelEndpointAction_Invoke_Happy exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllTunnelEndpointAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllTunnelEndpointAction_Invoke_NilClient exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllTunnelEndpointAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllTunnelEndpointAction_Invoke_BuildError exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllTunnelEndpointAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllTunnelEndpointAction_Invoke_SendError exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllTunnelEndpointAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{client: newTransportErrorClient(t)}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllTunnelEndpointAction_Invoke_APIError exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllTunnelEndpointAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_tunnel_endpoint")
}

// TestDeleteAllTunnelEndpointAction_Invoke_APIErrorReadBody exercises DeleteAllTunnelEndpointAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllTunnelEndpointAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllTunnelEndpointAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllTunnelEndpointActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
