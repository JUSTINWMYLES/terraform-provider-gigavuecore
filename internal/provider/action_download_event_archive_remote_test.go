package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDownloadEventArchiveAction_Invoke_Happy exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDownloadEventArchiveAction_Invoke_Happy(t *testing.T) {
	r := &DownloadEventArchiveAction{client: newMockClientStatus(t, 201, "{}")}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDownloadEventArchiveAction_Invoke_NilClient exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDownloadEventArchiveAction_Invoke_NilClient(t *testing.T) {
	r := &DownloadEventArchiveAction{}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDownloadEventArchiveAction_Invoke_BuildError exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDownloadEventArchiveAction_Invoke_BuildError(t *testing.T) {
	r := &DownloadEventArchiveAction{client: newMalformedBaseURLClient(t)}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDownloadEventArchiveAction_Invoke_SendError exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDownloadEventArchiveAction_Invoke_SendError(t *testing.T) {
	r := &DownloadEventArchiveAction{client: newTransportErrorClient(t)}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDownloadEventArchiveAction_Invoke_APIError exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDownloadEventArchiveAction_Invoke_APIError(t *testing.T) {
	r := &DownloadEventArchiveAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_download_event_archive")
}

// TestDownloadEventArchiveAction_Invoke_APIErrorReadBody exercises DownloadEventArchiveAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDownloadEventArchiveAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DownloadEventArchiveAction{client: newMockClientReadErrorBody(t, 501)}
	m := DownloadEventArchiveActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
