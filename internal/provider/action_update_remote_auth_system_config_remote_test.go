package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateRemoteAuthSystemConfigAction_Invoke_Happy exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateRemoteAuthSystemConfigAction_Invoke_NilClient exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateRemoteAuthSystemConfigAction_Invoke_BuildError exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateRemoteAuthSystemConfigAction_Invoke_SendError exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{client: newTransportErrorClient(t)}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateRemoteAuthSystemConfigAction_Invoke_APIError exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_remote_auth_system_config")
}

// TestUpdateRemoteAuthSystemConfigAction_Invoke_APIErrorReadBody exercises UpdateRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateRemoteAuthSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateRemoteAuthSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
