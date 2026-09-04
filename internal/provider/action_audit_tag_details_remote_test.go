package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAuditTagDetailsAction_Invoke_Happy exercises AuditTagDetailsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAuditTagDetailsAction_Invoke_Happy(t *testing.T) {
	r := &AuditTagDetailsAction{client: newMockClientStatus(t, 200, "{}")}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAuditTagDetailsAction_Invoke_NilClient exercises AuditTagDetailsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAuditTagDetailsAction_Invoke_NilClient(t *testing.T) {
	r := &AuditTagDetailsAction{}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAuditTagDetailsAction_Invoke_BuildError exercises AuditTagDetailsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAuditTagDetailsAction_Invoke_BuildError(t *testing.T) {
	r := &AuditTagDetailsAction{client: newMalformedBaseURLClient(t)}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAuditTagDetailsAction_Invoke_SendError exercises AuditTagDetailsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAuditTagDetailsAction_Invoke_SendError(t *testing.T) {
	r := &AuditTagDetailsAction{client: newTransportErrorClient(t)}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAuditTagDetailsAction_Invoke_APIError exercises AuditTagDetailsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAuditTagDetailsAction_Invoke_APIError(t *testing.T) {
	r := &AuditTagDetailsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_audit_tag_details")
}

// TestAuditTagDetailsAction_Invoke_APIErrorReadBody exercises AuditTagDetailsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAuditTagDetailsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AuditTagDetailsAction{client: newMockClientReadErrorBody(t, 501)}
	m := AuditTagDetailsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
