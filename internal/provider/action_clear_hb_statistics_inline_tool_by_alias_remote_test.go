package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_Happy exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_Happy(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_NilClient exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_NilClient(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_BuildError exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_BuildError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{client: newMalformedBaseURLClient(t)}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_SendError exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_SendError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{client: newTransportErrorClient(t)}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_APIError exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_APIError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_hb_statistics_inline_tool_by_alias")
}

// TestClearHbStatisticsInlineToolByAliasAction_Invoke_APIErrorReadBody exercises ClearHbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearHbStatisticsInlineToolByAliasAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearHbStatisticsInlineToolByAliasAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearHbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
