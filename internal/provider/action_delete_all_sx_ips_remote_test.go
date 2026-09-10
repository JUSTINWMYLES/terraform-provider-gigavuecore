package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllSxIpsAction_Invoke_Happy exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllSxIpsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllSxIpsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllSxIpsAction_Invoke_NilClient exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllSxIpsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllSxIpsAction{}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllSxIpsAction_Invoke_BuildError exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllSxIpsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllSxIpsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllSxIpsAction_Invoke_SendError exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllSxIpsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllSxIpsAction{client: newTransportErrorClient(t)}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllSxIpsAction_Invoke_APIError exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllSxIpsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllSxIpsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_sx_ips")
}

// TestDeleteAllSxIpsAction_Invoke_APIErrorReadBody exercises DeleteAllSxIpsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllSxIpsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllSxIpsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllSxIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
