package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSyslogConfigAction_Invoke_Happy exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSyslogConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSyslogConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSyslogConfigAction_Invoke_NilClient exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSyslogConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSyslogConfigAction{}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSyslogConfigAction_Invoke_BuildError exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSyslogConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSyslogConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSyslogConfigAction_Invoke_SendError exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSyslogConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSyslogConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSyslogConfigAction_Invoke_APIError exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSyslogConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSyslogConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_syslog_config")
}

// TestRedefineSyslogConfigAction_Invoke_APIErrorReadBody exercises RedefineSyslogConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSyslogConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSyslogConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSyslogConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
