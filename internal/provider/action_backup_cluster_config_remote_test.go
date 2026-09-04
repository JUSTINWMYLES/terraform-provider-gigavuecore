package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBackupClusterConfigAction_Invoke_Happy exercises BackupClusterConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBackupClusterConfigAction_Invoke_Happy(t *testing.T) {
	r := &BackupClusterConfigAction{client: newMockClientStatus(t, 201, "{}")}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBackupClusterConfigAction_Invoke_NilClient exercises BackupClusterConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBackupClusterConfigAction_Invoke_NilClient(t *testing.T) {
	r := &BackupClusterConfigAction{}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBackupClusterConfigAction_Invoke_BuildError exercises BackupClusterConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBackupClusterConfigAction_Invoke_BuildError(t *testing.T) {
	r := &BackupClusterConfigAction{client: newMalformedBaseURLClient(t)}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBackupClusterConfigAction_Invoke_SendError exercises BackupClusterConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBackupClusterConfigAction_Invoke_SendError(t *testing.T) {
	r := &BackupClusterConfigAction{client: newTransportErrorClient(t)}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBackupClusterConfigAction_Invoke_APIError exercises BackupClusterConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBackupClusterConfigAction_Invoke_APIError(t *testing.T) {
	r := &BackupClusterConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_backup_cluster_config")
}

// TestBackupClusterConfigAction_Invoke_APIErrorReadBody exercises BackupClusterConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBackupClusterConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BackupClusterConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := BackupClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
