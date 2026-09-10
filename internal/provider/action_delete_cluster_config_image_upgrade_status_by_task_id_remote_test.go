package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_Happy exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_Happy(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_NilClient exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_BuildError exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{client: newMalformedBaseURLClient(t)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_SendError exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_SendError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{client: newTransportErrorClient(t)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_APIError exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_APIError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_cluster_config_image_upgrade_status_by_task_id")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_APIErrorReadBody exercises DeleteClusterConfigImageUpgradeStatusByTaskIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteClusterConfigImageUpgradeStatusByTaskIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
