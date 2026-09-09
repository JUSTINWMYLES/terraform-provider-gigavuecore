package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_Happy exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_NilClient exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_BuildError exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_SendError exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{client: newTransportErrorClient(t)}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_APIError exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_gtp_whitelist_entries")
}

// TestDeleteAllGtpWhitelistEntriesAction_Invoke_APIErrorReadBody exercises DeleteAllGtpWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllGtpWhitelistEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllGtpWhitelistEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
