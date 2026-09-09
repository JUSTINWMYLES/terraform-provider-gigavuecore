package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_Happy exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_Happy(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_NilClient exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_BuildError exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{client: newMalformedBaseURLClient(t)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_SendError exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_SendError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{client: newTransportErrorClient(t)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_APIError exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_APIError(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_cluster_config_image_upgrade_status_by_task_group_id")
}

// TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_APIErrorReadBody exercises DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteClusterConfigImageUpgradeStatusByTaskGroupIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
