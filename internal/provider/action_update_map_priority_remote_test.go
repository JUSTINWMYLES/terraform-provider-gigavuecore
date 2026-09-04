package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapPriorityAction_Invoke_Happy exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapPriorityAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapPriorityAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapPriorityAction_Invoke_NilClient exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapPriorityAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapPriorityAction{}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapPriorityAction_Invoke_BuildError exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapPriorityAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapPriorityAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapPriorityAction_Invoke_SendError exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapPriorityAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapPriorityAction{client: newTransportErrorClient(t)}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapPriorityAction_Invoke_APIError exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapPriorityAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapPriorityAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_priority")
}

// TestUpdateMapPriorityAction_Invoke_APIErrorReadBody exercises UpdateMapPriorityAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapPriorityAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapPriorityAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
