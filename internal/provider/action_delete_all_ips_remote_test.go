package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllIpsAction_Invoke_Happy exercises DeleteAllIpsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllIpsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllIpsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllIpsAction_Invoke_NilClient exercises DeleteAllIpsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllIpsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllIpsAction{}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllIpsAction_Invoke_BuildError exercises DeleteAllIpsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllIpsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllIpsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllIpsAction_Invoke_SendError exercises DeleteAllIpsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllIpsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllIpsAction{client: newTransportErrorClient(t)}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllIpsAction_Invoke_APIError exercises DeleteAllIpsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllIpsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllIpsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_ips")
}

// TestDeleteAllIpsAction_Invoke_APIErrorReadBody exercises DeleteAllIpsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllIpsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllIpsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllIpsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
