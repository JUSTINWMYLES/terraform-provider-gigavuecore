package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDownloadSyslogArchiveAction_Invoke_Happy exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDownloadSyslogArchiveAction_Invoke_Happy(t *testing.T) {
	r := &DownloadSyslogArchiveAction{client: newMockClientStatus(t, 201, "{}")}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDownloadSyslogArchiveAction_Invoke_NilClient exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDownloadSyslogArchiveAction_Invoke_NilClient(t *testing.T) {
	r := &DownloadSyslogArchiveAction{}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDownloadSyslogArchiveAction_Invoke_BuildError exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDownloadSyslogArchiveAction_Invoke_BuildError(t *testing.T) {
	r := &DownloadSyslogArchiveAction{client: newMalformedBaseURLClient(t)}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDownloadSyslogArchiveAction_Invoke_SendError exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDownloadSyslogArchiveAction_Invoke_SendError(t *testing.T) {
	r := &DownloadSyslogArchiveAction{client: newTransportErrorClient(t)}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDownloadSyslogArchiveAction_Invoke_APIError exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDownloadSyslogArchiveAction_Invoke_APIError(t *testing.T) {
	r := &DownloadSyslogArchiveAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_download_syslog_archive")
}

// TestDownloadSyslogArchiveAction_Invoke_APIErrorReadBody exercises DownloadSyslogArchiveAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDownloadSyslogArchiveAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DownloadSyslogArchiveAction{client: newMockClientReadErrorBody(t, 501)}
	m := DownloadSyslogArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
