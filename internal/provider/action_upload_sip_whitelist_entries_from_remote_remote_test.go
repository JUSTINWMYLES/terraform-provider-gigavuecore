package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_Happy exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_Happy(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_NilClient exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_NilClient(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_BuildError exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_BuildError(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{client: newMalformedBaseURLClient(t)}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_SendError exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_SendError(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{client: newTransportErrorClient(t)}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_APIError exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_APIError(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_sip_whitelist_entries_from_remote")
}

// TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_APIErrorReadBody exercises UploadSipWhitelistEntriesFromRemoteAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadSipWhitelistEntriesFromRemoteAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadSipWhitelistEntriesFromRemoteAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadSipWhitelistEntriesFromRemoteActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
