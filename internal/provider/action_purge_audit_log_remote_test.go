package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPurgeAuditLogAction_Invoke_Happy exercises PurgeAuditLogAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPurgeAuditLogAction_Invoke_Happy(t *testing.T) {
	r := &PurgeAuditLogAction{client: newMockClientStatus(t, 204, "{}")}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPurgeAuditLogAction_Invoke_NilClient exercises PurgeAuditLogAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPurgeAuditLogAction_Invoke_NilClient(t *testing.T) {
	r := &PurgeAuditLogAction{}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPurgeAuditLogAction_Invoke_BuildError exercises PurgeAuditLogAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPurgeAuditLogAction_Invoke_BuildError(t *testing.T) {
	r := &PurgeAuditLogAction{client: newMalformedBaseURLClient(t)}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPurgeAuditLogAction_Invoke_SendError exercises PurgeAuditLogAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPurgeAuditLogAction_Invoke_SendError(t *testing.T) {
	r := &PurgeAuditLogAction{client: newTransportErrorClient(t)}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPurgeAuditLogAction_Invoke_APIError exercises PurgeAuditLogAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPurgeAuditLogAction_Invoke_APIError(t *testing.T) {
	r := &PurgeAuditLogAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_purge_audit_log")
}

// TestPurgeAuditLogAction_Invoke_APIErrorReadBody exercises PurgeAuditLogAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPurgeAuditLogAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PurgeAuditLogAction{client: newMockClientReadErrorBody(t, 501)}
	m := PurgeAuditLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
