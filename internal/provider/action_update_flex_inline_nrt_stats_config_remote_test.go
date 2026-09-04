package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_Happy exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_NilClient exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_BuildError exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_SendError exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{client: newTransportErrorClient(t)}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_APIError exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_flex_inline_nrt_stats_config")
}

// TestUpdateFlexInlineNrtStatsConfigAction_Invoke_APIErrorReadBody exercises UpdateFlexInlineNrtStatsConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateFlexInlineNrtStatsConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateFlexInlineNrtStatsConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateFlexInlineNrtStatsConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
