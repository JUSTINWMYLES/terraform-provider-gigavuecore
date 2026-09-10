package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_Happy exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_Happy(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{client: newMockClientStatus(t, 202, "{}")}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_NilClient exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_NilClient(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_BuildError exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_BuildError(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{client: newMalformedBaseURLClient(t)}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_SendError exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_SendError(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{client: newTransportErrorClient(t)}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_APIError exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_APIError(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_audit_traffic_flow_generated_f_map")
}

// TestAuditTrafficFlowGeneratedFMapAction_Invoke_APIErrorReadBody exercises AuditTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAuditTrafficFlowGeneratedFMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AuditTrafficFlowGeneratedFMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := AuditTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
