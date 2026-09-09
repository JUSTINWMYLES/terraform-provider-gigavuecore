package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSnmpServerSystemConfigAction_Invoke_Happy exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSnmpServerSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSnmpServerSystemConfigAction_Invoke_NilClient exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSnmpServerSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSnmpServerSystemConfigAction_Invoke_BuildError exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSnmpServerSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSnmpServerSystemConfigAction_Invoke_SendError exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSnmpServerSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSnmpServerSystemConfigAction_Invoke_APIError exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSnmpServerSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_snmp_server_system_config")
}

// TestRedefineSnmpServerSystemConfigAction_Invoke_APIErrorReadBody exercises RedefineSnmpServerSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSnmpServerSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSnmpServerSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSnmpServerSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
