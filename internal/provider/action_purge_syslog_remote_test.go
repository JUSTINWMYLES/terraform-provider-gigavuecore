package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPurgeSyslogAction_Invoke_Happy exercises PurgeSyslogAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPurgeSyslogAction_Invoke_Happy(t *testing.T) {
	r := &PurgeSyslogAction{client: newMockClientStatus(t, 204, "{}")}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPurgeSyslogAction_Invoke_NilClient exercises PurgeSyslogAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPurgeSyslogAction_Invoke_NilClient(t *testing.T) {
	r := &PurgeSyslogAction{}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPurgeSyslogAction_Invoke_BuildError exercises PurgeSyslogAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPurgeSyslogAction_Invoke_BuildError(t *testing.T) {
	r := &PurgeSyslogAction{client: newMalformedBaseURLClient(t)}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPurgeSyslogAction_Invoke_SendError exercises PurgeSyslogAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPurgeSyslogAction_Invoke_SendError(t *testing.T) {
	r := &PurgeSyslogAction{client: newTransportErrorClient(t)}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPurgeSyslogAction_Invoke_APIError exercises PurgeSyslogAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPurgeSyslogAction_Invoke_APIError(t *testing.T) {
	r := &PurgeSyslogAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_purge_syslog")
}

// TestPurgeSyslogAction_Invoke_APIErrorReadBody exercises PurgeSyslogAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPurgeSyslogAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PurgeSyslogAction{client: newMockClientReadErrorBody(t, 501)}
	m := PurgeSyslogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
