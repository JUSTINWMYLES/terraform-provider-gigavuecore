package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateIcapNrtStatsConfigAction_Invoke_Happy exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateIcapNrtStatsConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateIcapNrtStatsConfigAction_Invoke_NilClient exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateIcapNrtStatsConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateIcapNrtStatsConfigAction_Invoke_BuildError exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateIcapNrtStatsConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateIcapNrtStatsConfigAction_Invoke_SendError exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateIcapNrtStatsConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{client: newTransportErrorClient(t)}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateIcapNrtStatsConfigAction_Invoke_APIError exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateIcapNrtStatsConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_icap_nrt_stats_config")
}

// TestUpdateIcapNrtStatsConfigAction_Invoke_APIErrorReadBody exercises UpdateIcapNrtStatsConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateIcapNrtStatsConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateIcapNrtStatsConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateIcapNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
