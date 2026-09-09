package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearHbStatisticsInlineToolAction_Invoke_Happy exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearHbStatisticsInlineToolAction_Invoke_Happy(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearHbStatisticsInlineToolAction_Invoke_NilClient exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearHbStatisticsInlineToolAction_Invoke_NilClient(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearHbStatisticsInlineToolAction_Invoke_BuildError exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearHbStatisticsInlineToolAction_Invoke_BuildError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{client: newMalformedBaseURLClient(t)}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearHbStatisticsInlineToolAction_Invoke_SendError exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearHbStatisticsInlineToolAction_Invoke_SendError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{client: newTransportErrorClient(t)}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearHbStatisticsInlineToolAction_Invoke_APIError exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearHbStatisticsInlineToolAction_Invoke_APIError(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_hb_statistics_inline_tool")
}

// TestClearHbStatisticsInlineToolAction_Invoke_APIErrorReadBody exercises ClearHbStatisticsInlineToolAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearHbStatisticsInlineToolAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearHbStatisticsInlineToolAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearHbStatisticsInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
