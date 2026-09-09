package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEditAlarmSuppressionAction_Invoke_Happy exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEditAlarmSuppressionAction_Invoke_Happy(t *testing.T) {
	r := &EditAlarmSuppressionAction{client: newMockClientStatus(t, 200, "{}")}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEditAlarmSuppressionAction_Invoke_NilClient exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEditAlarmSuppressionAction_Invoke_NilClient(t *testing.T) {
	r := &EditAlarmSuppressionAction{}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEditAlarmSuppressionAction_Invoke_BuildError exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEditAlarmSuppressionAction_Invoke_BuildError(t *testing.T) {
	r := &EditAlarmSuppressionAction{client: newMalformedBaseURLClient(t)}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEditAlarmSuppressionAction_Invoke_SendError exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEditAlarmSuppressionAction_Invoke_SendError(t *testing.T) {
	r := &EditAlarmSuppressionAction{client: newTransportErrorClient(t)}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEditAlarmSuppressionAction_Invoke_APIError exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEditAlarmSuppressionAction_Invoke_APIError(t *testing.T) {
	r := &EditAlarmSuppressionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_edit_alarm_suppression")
}

// TestEditAlarmSuppressionAction_Invoke_APIErrorReadBody exercises EditAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEditAlarmSuppressionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EditAlarmSuppressionAction{client: newMockClientReadErrorBody(t, 501)}
	m := EditAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
