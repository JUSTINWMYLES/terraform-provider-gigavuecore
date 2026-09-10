package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAlarmSuppressionAction_Invoke_Happy exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAlarmSuppressionAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAlarmSuppressionAction_Invoke_NilClient exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAlarmSuppressionAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAlarmSuppressionAction_Invoke_BuildError exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAlarmSuppressionAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAlarmSuppressionAction_Invoke_SendError exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAlarmSuppressionAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{client: newTransportErrorClient(t)}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAlarmSuppressionAction_Invoke_APIError exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAlarmSuppressionAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_alarm_suppression")
}

// TestDeleteAlarmSuppressionAction_Invoke_APIErrorReadBody exercises DeleteAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAlarmSuppressionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAlarmSuppressionAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
