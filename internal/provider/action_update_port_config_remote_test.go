package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdatePortConfigAction_Invoke_Happy exercises UpdatePortConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdatePortConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdatePortConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdatePortConfigAction_Invoke_NilClient exercises UpdatePortConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdatePortConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdatePortConfigAction{}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdatePortConfigAction_Invoke_BuildError exercises UpdatePortConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdatePortConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdatePortConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdatePortConfigAction_Invoke_SendError exercises UpdatePortConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdatePortConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdatePortConfigAction{client: newTransportErrorClient(t)}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdatePortConfigAction_Invoke_APIError exercises UpdatePortConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdatePortConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdatePortConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_port_config")
}

// TestUpdatePortConfigAction_Invoke_APIErrorReadBody exercises UpdatePortConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdatePortConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdatePortConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdatePortConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
