package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSyslogConfigAction_Invoke_Happy exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSyslogConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSyslogConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSyslogConfigAction_Invoke_NilClient exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSyslogConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSyslogConfigAction{}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSyslogConfigAction_Invoke_BuildError exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSyslogConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSyslogConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSyslogConfigAction_Invoke_SendError exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSyslogConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSyslogConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSyslogConfigAction_Invoke_APIError exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSyslogConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSyslogConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_syslog_config")
}

// TestUpdateSyslogConfigAction_Invoke_APIErrorReadBody exercises UpdateSyslogConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSyslogConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSyslogConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
