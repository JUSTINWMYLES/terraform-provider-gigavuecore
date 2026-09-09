package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateMapChainAction_Invoke_Happy exercises UpdateMapChainAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateMapChainAction_Invoke_Happy(t *testing.T) {
	r := &UpdateMapChainAction{client: newMockClientStatus(t, 202, "{}")}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateMapChainAction_Invoke_NilClient exercises UpdateMapChainAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateMapChainAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateMapChainAction{}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateMapChainAction_Invoke_BuildError exercises UpdateMapChainAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateMapChainAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateMapChainAction{client: newMalformedBaseURLClient(t)}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateMapChainAction_Invoke_SendError exercises UpdateMapChainAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateMapChainAction_Invoke_SendError(t *testing.T) {
	r := &UpdateMapChainAction{client: newTransportErrorClient(t)}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateMapChainAction_Invoke_APIError exercises UpdateMapChainAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateMapChainAction_Invoke_APIError(t *testing.T) {
	r := &UpdateMapChainAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_map_chain")
}

// TestUpdateMapChainAction_Invoke_APIErrorReadBody exercises UpdateMapChainAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateMapChainAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateMapChainAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateMapChainActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
