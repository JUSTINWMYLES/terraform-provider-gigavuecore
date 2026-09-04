package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearNrtStatsCountersAction_Invoke_Happy exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearNrtStatsCountersAction_Invoke_Happy(t *testing.T) {
	r := &ClearNrtStatsCountersAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearNrtStatsCountersAction_Invoke_NilClient exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearNrtStatsCountersAction_Invoke_NilClient(t *testing.T) {
	r := &ClearNrtStatsCountersAction{}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearNrtStatsCountersAction_Invoke_BuildError exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearNrtStatsCountersAction_Invoke_BuildError(t *testing.T) {
	r := &ClearNrtStatsCountersAction{client: newMalformedBaseURLClient(t)}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearNrtStatsCountersAction_Invoke_SendError exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearNrtStatsCountersAction_Invoke_SendError(t *testing.T) {
	r := &ClearNrtStatsCountersAction{client: newTransportErrorClient(t)}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearNrtStatsCountersAction_Invoke_APIError exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearNrtStatsCountersAction_Invoke_APIError(t *testing.T) {
	r := &ClearNrtStatsCountersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_nrt_stats_counters")
}

// TestClearNrtStatsCountersAction_Invoke_APIErrorReadBody exercises ClearNrtStatsCountersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearNrtStatsCountersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearNrtStatsCountersAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
