package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateFabricResourceConfigAction_Invoke_Happy exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateFabricResourceConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateFabricResourceConfigAction_Invoke_NilClient exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateFabricResourceConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateFabricResourceConfigAction_Invoke_BuildError exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateFabricResourceConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateFabricResourceConfigAction_Invoke_SendError exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateFabricResourceConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{client: newTransportErrorClient(t)}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateFabricResourceConfigAction_Invoke_APIError exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateFabricResourceConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_fabric_resource_config")
}

// TestUpdateFabricResourceConfigAction_Invoke_APIErrorReadBody exercises UpdateFabricResourceConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateFabricResourceConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateFabricResourceConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateFabricResourceConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
