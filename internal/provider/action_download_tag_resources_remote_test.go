package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDownloadTagResourcesAction_Invoke_Happy exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDownloadTagResourcesAction_Invoke_Happy(t *testing.T) {
	r := &DownloadTagResourcesAction{client: newMockClientStatus(t, 200, "{}")}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDownloadTagResourcesAction_Invoke_NilClient exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDownloadTagResourcesAction_Invoke_NilClient(t *testing.T) {
	r := &DownloadTagResourcesAction{}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDownloadTagResourcesAction_Invoke_BuildError exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDownloadTagResourcesAction_Invoke_BuildError(t *testing.T) {
	r := &DownloadTagResourcesAction{client: newMalformedBaseURLClient(t)}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDownloadTagResourcesAction_Invoke_SendError exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDownloadTagResourcesAction_Invoke_SendError(t *testing.T) {
	r := &DownloadTagResourcesAction{client: newTransportErrorClient(t)}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDownloadTagResourcesAction_Invoke_APIError exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDownloadTagResourcesAction_Invoke_APIError(t *testing.T) {
	r := &DownloadTagResourcesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_download_tag_resources")
}

// TestDownloadTagResourcesAction_Invoke_APIErrorReadBody exercises DownloadTagResourcesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDownloadTagResourcesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DownloadTagResourcesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DownloadTagResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
