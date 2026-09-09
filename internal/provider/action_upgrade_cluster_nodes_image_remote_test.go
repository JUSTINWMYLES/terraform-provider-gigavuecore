package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpgradeClusterNodesImageAction_Invoke_Happy exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpgradeClusterNodesImageAction_Invoke_Happy(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpgradeClusterNodesImageAction_Invoke_NilClient exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpgradeClusterNodesImageAction_Invoke_NilClient(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpgradeClusterNodesImageAction_Invoke_BuildError exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpgradeClusterNodesImageAction_Invoke_BuildError(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{client: newMalformedBaseURLClient(t)}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpgradeClusterNodesImageAction_Invoke_SendError exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpgradeClusterNodesImageAction_Invoke_SendError(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{client: newTransportErrorClient(t)}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpgradeClusterNodesImageAction_Invoke_APIError exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpgradeClusterNodesImageAction_Invoke_APIError(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upgrade_cluster_nodes_image")
}

// TestUpgradeClusterNodesImageAction_Invoke_APIErrorReadBody exercises UpgradeClusterNodesImageAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpgradeClusterNodesImageAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpgradeClusterNodesImageAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpgradeClusterNodesImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
