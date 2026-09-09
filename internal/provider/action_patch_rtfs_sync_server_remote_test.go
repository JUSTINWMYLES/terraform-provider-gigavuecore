package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPatchRtfsSyncServerAction_Invoke_Happy exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPatchRtfsSyncServerAction_Invoke_Happy(t *testing.T) {
	r := &PatchRtfsSyncServerAction{client: newMockClientStatus(t, 201, "{}")}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPatchRtfsSyncServerAction_Invoke_NilClient exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPatchRtfsSyncServerAction_Invoke_NilClient(t *testing.T) {
	r := &PatchRtfsSyncServerAction{}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPatchRtfsSyncServerAction_Invoke_BuildError exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPatchRtfsSyncServerAction_Invoke_BuildError(t *testing.T) {
	r := &PatchRtfsSyncServerAction{client: newMalformedBaseURLClient(t)}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPatchRtfsSyncServerAction_Invoke_SendError exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPatchRtfsSyncServerAction_Invoke_SendError(t *testing.T) {
	r := &PatchRtfsSyncServerAction{client: newTransportErrorClient(t)}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPatchRtfsSyncServerAction_Invoke_APIError exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPatchRtfsSyncServerAction_Invoke_APIError(t *testing.T) {
	r := &PatchRtfsSyncServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_patch_rtfs_sync_server")
}

// TestPatchRtfsSyncServerAction_Invoke_APIErrorReadBody exercises PatchRtfsSyncServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPatchRtfsSyncServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PatchRtfsSyncServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := PatchRtfsSyncServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
