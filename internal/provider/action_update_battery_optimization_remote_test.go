package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateBatteryOptimizationAction_Invoke_Happy exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateBatteryOptimizationAction_Invoke_Happy(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateBatteryOptimizationAction_Invoke_NilClient exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateBatteryOptimizationAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateBatteryOptimizationAction_Invoke_BuildError exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateBatteryOptimizationAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{client: newMalformedBaseURLClient(t)}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateBatteryOptimizationAction_Invoke_SendError exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateBatteryOptimizationAction_Invoke_SendError(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{client: newTransportErrorClient(t)}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateBatteryOptimizationAction_Invoke_APIError exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateBatteryOptimizationAction_Invoke_APIError(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_battery_optimization")
}

// TestUpdateBatteryOptimizationAction_Invoke_APIErrorReadBody exercises UpdateBatteryOptimizationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateBatteryOptimizationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateBatteryOptimizationAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateBatteryOptimizationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
