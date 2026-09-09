package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllAppsTcpProfileAction_Invoke_Happy exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllAppsTcpProfileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllAppsTcpProfileAction_Invoke_NilClient exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllAppsTcpProfileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllAppsTcpProfileAction_Invoke_BuildError exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllAppsTcpProfileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllAppsTcpProfileAction_Invoke_SendError exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllAppsTcpProfileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{client: newTransportErrorClient(t)}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllAppsTcpProfileAction_Invoke_APIError exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllAppsTcpProfileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_apps_tcp_profile")
}

// TestDeleteAllAppsTcpProfileAction_Invoke_APIErrorReadBody exercises DeleteAllAppsTcpProfileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllAppsTcpProfileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllAppsTcpProfileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllAppsTcpProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
