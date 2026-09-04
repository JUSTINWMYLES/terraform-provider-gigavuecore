package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSnmpServerSystemConfigAction_Invoke_Happy exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSnmpServerSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSnmpServerSystemConfigAction_Invoke_NilClient exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSnmpServerSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSnmpServerSystemConfigAction_Invoke_BuildError exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSnmpServerSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSnmpServerSystemConfigAction_Invoke_SendError exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSnmpServerSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSnmpServerSystemConfigAction_Invoke_APIError exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSnmpServerSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_snmp_server_system_config")
}

// TestUpdateSnmpServerSystemConfigAction_Invoke_APIErrorReadBody exercises UpdateSnmpServerSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSnmpServerSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSnmpServerSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
