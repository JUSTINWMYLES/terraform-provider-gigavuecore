package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnacknowledgeMultipleAlarmsAction_Invoke_Happy exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_Happy(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{client: newMockClientStatus(t, 200, "{}")}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnacknowledgeMultipleAlarmsAction_Invoke_NilClient exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_NilClient(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnacknowledgeMultipleAlarmsAction_Invoke_BuildError exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_BuildError(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{client: newMalformedBaseURLClient(t)}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnacknowledgeMultipleAlarmsAction_Invoke_SendError exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_SendError(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{client: newTransportErrorClient(t)}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnacknowledgeMultipleAlarmsAction_Invoke_APIError exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_APIError(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unacknowledge_multiple_alarms")
}

// TestUnacknowledgeMultipleAlarmsAction_Invoke_APIErrorReadBody exercises UnacknowledgeMultipleAlarmsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnacknowledgeMultipleAlarmsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnacknowledgeMultipleAlarmsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnacknowledgeMultipleAlarmsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
