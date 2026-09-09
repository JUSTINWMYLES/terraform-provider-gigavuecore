package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_Happy exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_Happy(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_NilClient exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_NilClient(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_BuildError exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_BuildError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{client: newMalformedBaseURLClient(t)}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_SendError exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_SendError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{client: newTransportErrorClient(t)}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_APIError exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_APIError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_multiple_clusters")
}

// TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_APIErrorReadBody exercises MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateAllFlowMapToTrafficFlowsInMultipleClustersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateAllFlowMapToTrafficFlowsInMultipleClustersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
