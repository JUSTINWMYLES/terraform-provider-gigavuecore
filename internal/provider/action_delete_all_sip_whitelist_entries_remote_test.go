package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllSipWhitelistEntriesAction_Invoke_Happy exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllSipWhitelistEntriesAction_Invoke_NilClient exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllSipWhitelistEntriesAction_Invoke_BuildError exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllSipWhitelistEntriesAction_Invoke_SendError exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{client: newTransportErrorClient(t)}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllSipWhitelistEntriesAction_Invoke_APIError exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_sip_whitelist_entries")
}

// TestDeleteAllSipWhitelistEntriesAction_Invoke_APIErrorReadBody exercises DeleteAllSipWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllSipWhitelistEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllSipWhitelistEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
