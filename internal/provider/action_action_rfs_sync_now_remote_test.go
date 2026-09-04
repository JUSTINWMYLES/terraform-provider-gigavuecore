package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestActionRfsSyncNowAction_Invoke_Happy exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestActionRfsSyncNowAction_Invoke_Happy(t *testing.T) {
	r := &ActionRfsSyncNowAction{client: newMockClientStatus(t, 201, "{}")}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestActionRfsSyncNowAction_Invoke_NilClient exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestActionRfsSyncNowAction_Invoke_NilClient(t *testing.T) {
	r := &ActionRfsSyncNowAction{}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestActionRfsSyncNowAction_Invoke_BuildError exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestActionRfsSyncNowAction_Invoke_BuildError(t *testing.T) {
	r := &ActionRfsSyncNowAction{client: newMalformedBaseURLClient(t)}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestActionRfsSyncNowAction_Invoke_SendError exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestActionRfsSyncNowAction_Invoke_SendError(t *testing.T) {
	r := &ActionRfsSyncNowAction{client: newTransportErrorClient(t)}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestActionRfsSyncNowAction_Invoke_APIError exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestActionRfsSyncNowAction_Invoke_APIError(t *testing.T) {
	r := &ActionRfsSyncNowAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_action_rfs_sync_now")
}

// TestActionRfsSyncNowAction_Invoke_APIErrorReadBody exercises ActionRfsSyncNowAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestActionRfsSyncNowAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ActionRfsSyncNowAction{client: newMockClientReadErrorBody(t, 501)}
	m := ActionRfsSyncNowActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
