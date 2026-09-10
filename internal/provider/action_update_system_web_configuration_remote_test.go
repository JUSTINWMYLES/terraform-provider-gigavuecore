package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSystemWebConfigurationAction_Invoke_Happy exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSystemWebConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSystemWebConfigurationAction_Invoke_NilClient exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSystemWebConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSystemWebConfigurationAction_Invoke_BuildError exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSystemWebConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSystemWebConfigurationAction_Invoke_SendError exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSystemWebConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{client: newTransportErrorClient(t)}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSystemWebConfigurationAction_Invoke_APIError exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSystemWebConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_system_web_configuration")
}

// TestUpdateSystemWebConfigurationAction_Invoke_APIErrorReadBody exercises UpdateSystemWebConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSystemWebConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSystemWebConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSystemWebConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
