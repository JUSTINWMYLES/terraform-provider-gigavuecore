package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_Happy exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_NilClient exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_BuildError exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_SendError exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_APIError exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows")
}

// TestMigrateFabricMapsToTrafficFlowsAction_Invoke_APIErrorReadBody exercises MigrateFabricMapsToTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateFabricMapsToTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateFabricMapsToTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
