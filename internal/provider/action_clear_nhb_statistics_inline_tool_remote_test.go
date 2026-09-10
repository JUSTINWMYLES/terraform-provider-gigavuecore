package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearNhbStatisticsInlineToolAction_Invoke_Happy exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearNhbStatisticsInlineToolAction_Invoke_Happy(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearNhbStatisticsInlineToolAction_Invoke_NilClient exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearNhbStatisticsInlineToolAction_Invoke_NilClient(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearNhbStatisticsInlineToolAction_Invoke_BuildError exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearNhbStatisticsInlineToolAction_Invoke_BuildError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{client: newMalformedBaseURLClient(t)}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearNhbStatisticsInlineToolAction_Invoke_SendError exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearNhbStatisticsInlineToolAction_Invoke_SendError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{client: newTransportErrorClient(t)}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearNhbStatisticsInlineToolAction_Invoke_APIError exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearNhbStatisticsInlineToolAction_Invoke_APIError(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_nhb_statistics_inline_tool")
}

// TestClearNhbStatisticsInlineToolAction_Invoke_APIErrorReadBody exercises ClearNhbStatisticsInlineToolAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearNhbStatisticsInlineToolAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearNhbStatisticsInlineToolAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearNhbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
