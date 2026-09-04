package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteGtpWhitelistEntryAction_Invoke_Happy exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteGtpWhitelistEntryAction_Invoke_Happy(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteGtpWhitelistEntryAction_Invoke_NilClient exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteGtpWhitelistEntryAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteGtpWhitelistEntryAction_Invoke_BuildError exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteGtpWhitelistEntryAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{client: newMalformedBaseURLClient(t)}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteGtpWhitelistEntryAction_Invoke_SendError exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteGtpWhitelistEntryAction_Invoke_SendError(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{client: newTransportErrorClient(t)}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteGtpWhitelistEntryAction_Invoke_APIError exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteGtpWhitelistEntryAction_Invoke_APIError(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_gtp_whitelist_entry")
}

// TestDeleteGtpWhitelistEntryAction_Invoke_APIErrorReadBody exercises DeleteGtpWhitelistEntryAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteGtpWhitelistEntryAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteGtpWhitelistEntryAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteGtpWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
