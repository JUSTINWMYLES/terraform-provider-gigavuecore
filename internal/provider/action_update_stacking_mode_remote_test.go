package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateStackingModeAction_Invoke_Happy exercises UpdateStackingModeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateStackingModeAction_Invoke_Happy(t *testing.T) {
	r := &UpdateStackingModeAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateStackingModeAction_Invoke_NilClient exercises UpdateStackingModeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateStackingModeAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateStackingModeAction{}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateStackingModeAction_Invoke_BuildError exercises UpdateStackingModeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateStackingModeAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateStackingModeAction{client: newMalformedBaseURLClient(t)}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateStackingModeAction_Invoke_SendError exercises UpdateStackingModeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateStackingModeAction_Invoke_SendError(t *testing.T) {
	r := &UpdateStackingModeAction{client: newTransportErrorClient(t)}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateStackingModeAction_Invoke_APIError exercises UpdateStackingModeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateStackingModeAction_Invoke_APIError(t *testing.T) {
	r := &UpdateStackingModeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_stacking_mode")
}

// TestUpdateStackingModeAction_Invoke_APIErrorReadBody exercises UpdateStackingModeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateStackingModeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateStackingModeAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateStackingModeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
