package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllPortThrottlesAction_Invoke_Happy exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllPortThrottlesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllPortThrottlesAction_Invoke_NilClient exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllPortThrottlesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllPortThrottlesAction_Invoke_BuildError exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllPortThrottlesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllPortThrottlesAction_Invoke_SendError exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllPortThrottlesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{client: newTransportErrorClient(t)}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllPortThrottlesAction_Invoke_APIError exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllPortThrottlesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_port_throttles")
}

// TestDeleteAllPortThrottlesAction_Invoke_APIErrorReadBody exercises DeleteAllPortThrottlesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllPortThrottlesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllPortThrottlesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllPortThrottlesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
