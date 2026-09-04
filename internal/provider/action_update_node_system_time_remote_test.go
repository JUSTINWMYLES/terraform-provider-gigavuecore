package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateNodeSystemTimeAction_Invoke_Happy exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateNodeSystemTimeAction_Invoke_Happy(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateNodeSystemTimeAction_Invoke_NilClient exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateNodeSystemTimeAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateNodeSystemTimeAction_Invoke_BuildError exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateNodeSystemTimeAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{client: newMalformedBaseURLClient(t)}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateNodeSystemTimeAction_Invoke_SendError exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateNodeSystemTimeAction_Invoke_SendError(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{client: newTransportErrorClient(t)}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateNodeSystemTimeAction_Invoke_APIError exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateNodeSystemTimeAction_Invoke_APIError(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_node_system_time")
}

// TestUpdateNodeSystemTimeAction_Invoke_APIErrorReadBody exercises UpdateNodeSystemTimeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateNodeSystemTimeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateNodeSystemTimeAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateNodeSystemTimeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
