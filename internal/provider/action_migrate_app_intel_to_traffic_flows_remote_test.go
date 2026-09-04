package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_Happy exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_NilClient exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_BuildError exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_SendError exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_APIError exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_app_intel_to_traffic_flows")
}

// TestMigrateAppIntelToTrafficFlowsAction_Invoke_APIErrorReadBody exercises MigrateAppIntelToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateAppIntelToTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateAppIntelToTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateAppIntelToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
