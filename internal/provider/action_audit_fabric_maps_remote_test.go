package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAuditFabricMapsAction_Invoke_Happy exercises AuditFabricMapsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAuditFabricMapsAction_Invoke_Happy(t *testing.T) {
	r := &AuditFabricMapsAction{client: newMockClientStatus(t, 202, "{}")}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAuditFabricMapsAction_Invoke_NilClient exercises AuditFabricMapsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAuditFabricMapsAction_Invoke_NilClient(t *testing.T) {
	r := &AuditFabricMapsAction{}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAuditFabricMapsAction_Invoke_BuildError exercises AuditFabricMapsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAuditFabricMapsAction_Invoke_BuildError(t *testing.T) {
	r := &AuditFabricMapsAction{client: newMalformedBaseURLClient(t)}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAuditFabricMapsAction_Invoke_SendError exercises AuditFabricMapsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAuditFabricMapsAction_Invoke_SendError(t *testing.T) {
	r := &AuditFabricMapsAction{client: newTransportErrorClient(t)}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAuditFabricMapsAction_Invoke_APIError exercises AuditFabricMapsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAuditFabricMapsAction_Invoke_APIError(t *testing.T) {
	r := &AuditFabricMapsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_audit_fabric_maps")
}

// TestAuditFabricMapsAction_Invoke_APIErrorReadBody exercises AuditFabricMapsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAuditFabricMapsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AuditFabricMapsAction{client: newMockClientReadErrorBody(t, 501)}
	m := AuditFabricMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
