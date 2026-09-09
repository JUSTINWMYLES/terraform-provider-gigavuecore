package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateAlarmAction_Invoke_Happy exercises UpdateAlarmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateAlarmAction_Invoke_Happy(t *testing.T) {
	r := &UpdateAlarmAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateAlarmAction_Invoke_NilClient exercises UpdateAlarmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateAlarmAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateAlarmAction{}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateAlarmAction_Invoke_BuildError exercises UpdateAlarmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateAlarmAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateAlarmAction{client: newMalformedBaseURLClient(t)}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateAlarmAction_Invoke_SendError exercises UpdateAlarmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateAlarmAction_Invoke_SendError(t *testing.T) {
	r := &UpdateAlarmAction{client: newTransportErrorClient(t)}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateAlarmAction_Invoke_APIError exercises UpdateAlarmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateAlarmAction_Invoke_APIError(t *testing.T) {
	r := &UpdateAlarmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_alarm")
}

// TestUpdateAlarmAction_Invoke_APIErrorReadBody exercises UpdateAlarmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateAlarmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateAlarmAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateAlarmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
