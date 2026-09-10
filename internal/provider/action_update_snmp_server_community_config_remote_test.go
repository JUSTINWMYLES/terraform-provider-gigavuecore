package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSnmpServerCommunityConfigAction_Invoke_Happy exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSnmpServerCommunityConfigAction_Invoke_NilClient exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSnmpServerCommunityConfigAction_Invoke_BuildError exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSnmpServerCommunityConfigAction_Invoke_SendError exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSnmpServerCommunityConfigAction_Invoke_APIError exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_snmp_server_community_config")
}

// TestUpdateSnmpServerCommunityConfigAction_Invoke_APIErrorReadBody exercises UpdateSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSnmpServerCommunityConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSnmpServerCommunityConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
