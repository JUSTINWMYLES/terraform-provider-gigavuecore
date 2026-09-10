package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearIcapNrtStatsCountersAction_Invoke_Happy exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearIcapNrtStatsCountersAction_Invoke_Happy(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearIcapNrtStatsCountersAction_Invoke_NilClient exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearIcapNrtStatsCountersAction_Invoke_NilClient(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearIcapNrtStatsCountersAction_Invoke_BuildError exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearIcapNrtStatsCountersAction_Invoke_BuildError(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{client: newMalformedBaseURLClient(t)}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearIcapNrtStatsCountersAction_Invoke_SendError exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearIcapNrtStatsCountersAction_Invoke_SendError(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{client: newTransportErrorClient(t)}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearIcapNrtStatsCountersAction_Invoke_APIError exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearIcapNrtStatsCountersAction_Invoke_APIError(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_icap_nrt_stats_counters")
}

// TestClearIcapNrtStatsCountersAction_Invoke_APIErrorReadBody exercises ClearIcapNrtStatsCountersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearIcapNrtStatsCountersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearIcapNrtStatsCountersAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearIcapNrtStatsCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
