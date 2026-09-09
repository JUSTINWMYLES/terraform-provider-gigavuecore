package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapChainPriorityAction_Invoke_Happy exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapChainPriorityAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapChainPriorityAction{client: newMockClientStatus(t, 202, "{}")}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapChainPriorityAction_Invoke_NilClient exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapChainPriorityAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapChainPriorityAction{}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapChainPriorityAction_Invoke_BuildError exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapChainPriorityAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapChainPriorityAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapChainPriorityAction_Invoke_SendError exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapChainPriorityAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapChainPriorityAction{client: newTransportErrorClient(t)}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapChainPriorityAction_Invoke_APIError exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapChainPriorityAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapChainPriorityAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_chain_priority")
}

// TestUpdateMapChainPriorityAction_Invoke_APIErrorReadBody exercises UpdateMapChainPriorityAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapChainPriorityAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapChainPriorityAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapChainPriorityActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
