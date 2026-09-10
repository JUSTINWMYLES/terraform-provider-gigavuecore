package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSnmpServerCommunityConfigAction_Invoke_Happy exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSnmpServerCommunityConfigAction_Invoke_NilClient exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSnmpServerCommunityConfigAction_Invoke_BuildError exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSnmpServerCommunityConfigAction_Invoke_SendError exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSnmpServerCommunityConfigAction_Invoke_APIError exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_snmp_server_community_config")
}

// TestRedefineSnmpServerCommunityConfigAction_Invoke_APIErrorReadBody exercises RedefineSnmpServerCommunityConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSnmpServerCommunityConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSnmpServerCommunityConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSnmpServerCommunityConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
