package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAuditFmTemplateAction_Invoke_Happy exercises AuditFmTemplateAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAuditFmTemplateAction_Invoke_Happy(t *testing.T) {
	r := &AuditFmTemplateAction{client: newMockClientStatus(t, 200, "{}")}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAuditFmTemplateAction_Invoke_NilClient exercises AuditFmTemplateAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAuditFmTemplateAction_Invoke_NilClient(t *testing.T) {
	r := &AuditFmTemplateAction{}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAuditFmTemplateAction_Invoke_BuildError exercises AuditFmTemplateAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAuditFmTemplateAction_Invoke_BuildError(t *testing.T) {
	r := &AuditFmTemplateAction{client: newMalformedBaseURLClient(t)}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAuditFmTemplateAction_Invoke_SendError exercises AuditFmTemplateAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAuditFmTemplateAction_Invoke_SendError(t *testing.T) {
	r := &AuditFmTemplateAction{client: newTransportErrorClient(t)}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAuditFmTemplateAction_Invoke_APIError exercises AuditFmTemplateAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAuditFmTemplateAction_Invoke_APIError(t *testing.T) {
	r := &AuditFmTemplateAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_audit_fm_template")
}

// TestAuditFmTemplateAction_Invoke_APIErrorReadBody exercises AuditFmTemplateAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAuditFmTemplateAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AuditFmTemplateAction{client: newMockClientReadErrorBody(t, 501)}
	m := AuditFmTemplateActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
