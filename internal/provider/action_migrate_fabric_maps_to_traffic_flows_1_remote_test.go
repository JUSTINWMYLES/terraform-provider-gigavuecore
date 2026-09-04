package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_Happy exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_Happy(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_NilClient exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_NilClient(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_BuildError exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_BuildError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{client: newMalformedBaseURLClient(t)}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_SendError exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_SendError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{client: newTransportErrorClient(t)}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_APIError exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_APIError(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_fabric_maps_to_traffic_flows_1")
}

// TestMigrateFabricMapsToTrafficFlows1Action_Invoke_APIErrorReadBody exercises MigrateFabricMapsToTrafficFlows1Action.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateFabricMapsToTrafficFlows1Action_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateFabricMapsToTrafficFlows1Action{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateFabricMapsToTrafficFlows1ActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
