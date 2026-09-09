package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSnmpServerNofifyConfigAction_Invoke_Happy exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSnmpServerNofifyConfigAction_Invoke_NilClient exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSnmpServerNofifyConfigAction_Invoke_BuildError exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSnmpServerNofifyConfigAction_Invoke_SendError exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSnmpServerNofifyConfigAction_Invoke_APIError exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_snmp_server_nofify_config")
}

// TestRedefineSnmpServerNofifyConfigAction_Invoke_APIErrorReadBody exercises RedefineSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSnmpServerNofifyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSnmpServerNofifyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
