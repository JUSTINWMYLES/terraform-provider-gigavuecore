package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_Happy exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_Happy(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_NilClient exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_NilClient(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_BuildError exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_BuildError(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{client: newMalformedBaseURLClient(t)}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_SendError exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_SendError(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{client: newTransportErrorClient(t)}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_APIError exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_APIError(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_gtp_whitelist_entries_from_url")
}

// TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_APIErrorReadBody exercises UploadGtpWhitelistEntriesFromUrlAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadGtpWhitelistEntriesFromUrlAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadGtpWhitelistEntriesFromUrlAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadGtpWhitelistEntriesFromUrlActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
