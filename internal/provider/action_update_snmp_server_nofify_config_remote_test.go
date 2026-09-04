package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSnmpServerNofifyConfigAction_Invoke_Happy exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSnmpServerNofifyConfigAction_Invoke_NilClient exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSnmpServerNofifyConfigAction_Invoke_BuildError exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSnmpServerNofifyConfigAction_Invoke_SendError exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSnmpServerNofifyConfigAction_Invoke_APIError exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_snmp_server_nofify_config")
}

// TestUpdateSnmpServerNofifyConfigAction_Invoke_APIErrorReadBody exercises UpdateSnmpServerNofifyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSnmpServerNofifyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSnmpServerNofifyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSnmpServerNofifyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
