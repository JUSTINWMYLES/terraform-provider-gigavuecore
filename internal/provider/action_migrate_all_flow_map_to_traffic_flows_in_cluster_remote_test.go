package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_Happy exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_Happy(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{client: newMockClientStatus(t, 200, "{}")}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_NilClient exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_NilClient(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_BuildError exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_BuildError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{client: newMalformedBaseURLClient(t)}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_SendError exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_SendError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{client: newTransportErrorClient(t)}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_APIError exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_APIError(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster")
}

// TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_APIErrorReadBody exercises MigrateAllFlowMapToTrafficFlowsInClusterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestMigrateAllFlowMapToTrafficFlowsInClusterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &MigrateAllFlowMapToTrafficFlowsInClusterAction{client: newMockClientReadErrorBody(t, 501)}
	m := MigrateAllFlowMapToTrafficFlowsInClusterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
