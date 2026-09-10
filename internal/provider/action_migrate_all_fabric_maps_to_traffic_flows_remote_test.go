package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_Happy exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_NilClient exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_BuildError exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_SendError exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_APIError exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_all_fabric_maps_to_traffic_flows")
}

// TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_APIErrorReadBody exercises MigrateAllFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateAllFabricMapsToTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateAllFabricMapsToTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateAllFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
