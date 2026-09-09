package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEnableAlarmSuppressionAction_Invoke_Happy exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEnableAlarmSuppressionAction_Invoke_Happy(t *testing.T) {
	r := &EnableAlarmSuppressionAction{client: newMockClientStatus(t, 200, "{}")}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnableAlarmSuppressionAction_Invoke_NilClient exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnableAlarmSuppressionAction_Invoke_NilClient(t *testing.T) {
	r := &EnableAlarmSuppressionAction{}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnableAlarmSuppressionAction_Invoke_BuildError exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnableAlarmSuppressionAction_Invoke_BuildError(t *testing.T) {
	r := &EnableAlarmSuppressionAction{client: newMalformedBaseURLClient(t)}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnableAlarmSuppressionAction_Invoke_SendError exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnableAlarmSuppressionAction_Invoke_SendError(t *testing.T) {
	r := &EnableAlarmSuppressionAction{client: newTransportErrorClient(t)}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnableAlarmSuppressionAction_Invoke_APIError exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnableAlarmSuppressionAction_Invoke_APIError(t *testing.T) {
	r := &EnableAlarmSuppressionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_enable_alarm_suppression")
}

// TestEnableAlarmSuppressionAction_Invoke_APIErrorReadBody exercises EnableAlarmSuppressionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnableAlarmSuppressionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EnableAlarmSuppressionAction{client: newMockClientReadErrorBody(t, 501)}
	m := EnableAlarmSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
