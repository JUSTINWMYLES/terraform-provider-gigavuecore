package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_Happy exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_Happy(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_NilClient exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_BuildError exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{client: newMalformedBaseURLClient(t)}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_SendError exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_SendError(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{client: newTransportErrorClient(t)}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_APIError exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_APIError(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_traffic_flows_global_settings")
}

// TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_APIErrorReadBody exercises UpdateTrafficFlowsGlobalSettingsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateTrafficFlowsGlobalSettingsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateTrafficFlowsGlobalSettingsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateTrafficFlowsGlobalSettingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
