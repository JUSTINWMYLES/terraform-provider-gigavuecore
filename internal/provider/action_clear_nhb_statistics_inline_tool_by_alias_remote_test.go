package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_Happy exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_Happy(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_NilClient exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_NilClient(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_BuildError exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_BuildError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{client: newMalformedBaseURLClient(t)}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_SendError exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_SendError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{client: newTransportErrorClient(t)}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_APIError exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_APIError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_nhb_statistics_inline_tool_by_alias")
}

// TestClearNhbStatisticsInlineToolByAliasAction_Invoke_APIErrorReadBody exercises ClearNhbStatisticsInlineToolByAliasAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearNhbStatisticsInlineToolByAliasAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolByAliasAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearNhbStatisticsInlineToolByAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
