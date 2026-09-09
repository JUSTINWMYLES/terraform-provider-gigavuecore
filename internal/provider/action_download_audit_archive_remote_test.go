package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDownloadAuditArchiveAction_Invoke_Happy exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDownloadAuditArchiveAction_Invoke_Happy(t *testing.T) {
	r := &DownloadAuditArchiveAction{client: newMockClientStatus(t, 201, "{}")}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDownloadAuditArchiveAction_Invoke_NilClient exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDownloadAuditArchiveAction_Invoke_NilClient(t *testing.T) {
	r := &DownloadAuditArchiveAction{}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDownloadAuditArchiveAction_Invoke_BuildError exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDownloadAuditArchiveAction_Invoke_BuildError(t *testing.T) {
	r := &DownloadAuditArchiveAction{client: newMalformedBaseURLClient(t)}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDownloadAuditArchiveAction_Invoke_SendError exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDownloadAuditArchiveAction_Invoke_SendError(t *testing.T) {
	r := &DownloadAuditArchiveAction{client: newTransportErrorClient(t)}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDownloadAuditArchiveAction_Invoke_APIError exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDownloadAuditArchiveAction_Invoke_APIError(t *testing.T) {
	r := &DownloadAuditArchiveAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_download_audit_archive")
}

// TestDownloadAuditArchiveAction_Invoke_APIErrorReadBody exercises DownloadAuditArchiveAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDownloadAuditArchiveAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DownloadAuditArchiveAction{client: newMockClientReadErrorBody(t, 501)}
	m := DownloadAuditArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
