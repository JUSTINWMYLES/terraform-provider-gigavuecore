package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSetRfsSyncServerAction_Invoke_Happy exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSetRfsSyncServerAction_Invoke_Happy(t *testing.T) {
	r := &SetRfsSyncServerAction{client: newMockClientStatus(t, 201, "{}")}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSetRfsSyncServerAction_Invoke_NilClient exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSetRfsSyncServerAction_Invoke_NilClient(t *testing.T) {
	r := &SetRfsSyncServerAction{}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSetRfsSyncServerAction_Invoke_BuildError exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSetRfsSyncServerAction_Invoke_BuildError(t *testing.T) {
	r := &SetRfsSyncServerAction{client: newMalformedBaseURLClient(t)}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSetRfsSyncServerAction_Invoke_SendError exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSetRfsSyncServerAction_Invoke_SendError(t *testing.T) {
	r := &SetRfsSyncServerAction{client: newTransportErrorClient(t)}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSetRfsSyncServerAction_Invoke_APIError exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSetRfsSyncServerAction_Invoke_APIError(t *testing.T) {
	r := &SetRfsSyncServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_set_rfs_sync_server")
}

// TestSetRfsSyncServerAction_Invoke_APIErrorReadBody exercises SetRfsSyncServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSetRfsSyncServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SetRfsSyncServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := SetRfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
