package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateL2CircuitResourceConfigAction_Invoke_Happy exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateL2CircuitResourceConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateL2CircuitResourceConfigAction_Invoke_NilClient exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateL2CircuitResourceConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateL2CircuitResourceConfigAction_Invoke_BuildError exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateL2CircuitResourceConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateL2CircuitResourceConfigAction_Invoke_SendError exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateL2CircuitResourceConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{client: newTransportErrorClient(t)}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateL2CircuitResourceConfigAction_Invoke_APIError exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateL2CircuitResourceConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_l2_circuit_resource_config")
}

// TestUpdateL2CircuitResourceConfigAction_Invoke_APIErrorReadBody exercises UpdateL2CircuitResourceConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateL2CircuitResourceConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateL2CircuitResourceConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateL2CircuitResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
