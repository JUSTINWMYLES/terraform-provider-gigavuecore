package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestEnableAlarmAutoSuppressionAction_Invoke_Happy exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestEnableAlarmAutoSuppressionAction_Invoke_Happy(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{client: newMockClientStatus(t, 200, "{}")}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestEnableAlarmAutoSuppressionAction_Invoke_NilClient exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestEnableAlarmAutoSuppressionAction_Invoke_NilClient(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestEnableAlarmAutoSuppressionAction_Invoke_BuildError exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestEnableAlarmAutoSuppressionAction_Invoke_BuildError(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{client: newMalformedBaseURLClient(t)}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestEnableAlarmAutoSuppressionAction_Invoke_SendError exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestEnableAlarmAutoSuppressionAction_Invoke_SendError(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{client: newTransportErrorClient(t)}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestEnableAlarmAutoSuppressionAction_Invoke_APIError exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestEnableAlarmAutoSuppressionAction_Invoke_APIError(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_enable_alarm_auto_suppression")
}

// TestEnableAlarmAutoSuppressionAction_Invoke_APIErrorReadBody exercises EnableAlarmAutoSuppressionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestEnableAlarmAutoSuppressionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &EnableAlarmAutoSuppressionAction{client: newMockClientReadErrorBody(t, 501)}
	m := EnableAlarmAutoSuppressionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
