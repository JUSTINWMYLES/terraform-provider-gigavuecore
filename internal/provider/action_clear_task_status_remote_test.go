package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearTaskStatusAction_Invoke_Happy exercises ClearTaskStatusAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearTaskStatusAction_Invoke_Happy(t *testing.T) {
	r := &ClearTaskStatusAction{client: newMockClientStatus(t, 200, "{}")}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearTaskStatusAction_Invoke_NilClient exercises ClearTaskStatusAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearTaskStatusAction_Invoke_NilClient(t *testing.T) {
	r := &ClearTaskStatusAction{}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearTaskStatusAction_Invoke_BuildError exercises ClearTaskStatusAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearTaskStatusAction_Invoke_BuildError(t *testing.T) {
	r := &ClearTaskStatusAction{client: newMalformedBaseURLClient(t)}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearTaskStatusAction_Invoke_SendError exercises ClearTaskStatusAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearTaskStatusAction_Invoke_SendError(t *testing.T) {
	r := &ClearTaskStatusAction{client: newTransportErrorClient(t)}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearTaskStatusAction_Invoke_APIError exercises ClearTaskStatusAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearTaskStatusAction_Invoke_APIError(t *testing.T) {
	r := &ClearTaskStatusAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_task_status")
}

// TestClearTaskStatusAction_Invoke_APIErrorReadBody exercises ClearTaskStatusAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearTaskStatusAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearTaskStatusAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearTaskStatusActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
