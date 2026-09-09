package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_Happy exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_Happy(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_NilClient exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_BuildError exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{client: newMalformedBaseURLClient(t)}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_SendError exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_SendError(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{client: newTransportErrorClient(t)}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_APIError exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_APIError(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_stored_cluster_config_backup_snapshot")
}

// TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_APIErrorReadBody exercises DeleteStoredClusterConfigBackupSnapshotAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteStoredClusterConfigBackupSnapshotAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteStoredClusterConfigBackupSnapshotAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteStoredClusterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
