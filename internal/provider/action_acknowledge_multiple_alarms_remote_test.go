package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAcknowledgeMultipleAlarmsAction_Invoke_Happy exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAcknowledgeMultipleAlarmsAction_Invoke_Happy(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{client: newMockClientStatus(t, 200, "{}")}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAcknowledgeMultipleAlarmsAction_Invoke_NilClient exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAcknowledgeMultipleAlarmsAction_Invoke_NilClient(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAcknowledgeMultipleAlarmsAction_Invoke_BuildError exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAcknowledgeMultipleAlarmsAction_Invoke_BuildError(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{client: newMalformedBaseURLClient(t)}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAcknowledgeMultipleAlarmsAction_Invoke_SendError exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAcknowledgeMultipleAlarmsAction_Invoke_SendError(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{client: newTransportErrorClient(t)}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAcknowledgeMultipleAlarmsAction_Invoke_APIError exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAcknowledgeMultipleAlarmsAction_Invoke_APIError(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_acknowledge_multiple_alarms")
}

// TestAcknowledgeMultipleAlarmsAction_Invoke_APIErrorReadBody exercises AcknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAcknowledgeMultipleAlarmsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AcknowledgeMultipleAlarmsAction{client: newMockClientReadErrorBody(t, 501)}
	m := AcknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
