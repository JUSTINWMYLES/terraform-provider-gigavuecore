package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBackupClustersConfigAction_Invoke_Happy exercises BackupClustersConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBackupClustersConfigAction_Invoke_Happy(t *testing.T) {
	r := &BackupClustersConfigAction{client: newMockClientStatus(t, 201, "{}")}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBackupClustersConfigAction_Invoke_NilClient exercises BackupClustersConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBackupClustersConfigAction_Invoke_NilClient(t *testing.T) {
	r := &BackupClustersConfigAction{}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBackupClustersConfigAction_Invoke_BuildError exercises BackupClustersConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBackupClustersConfigAction_Invoke_BuildError(t *testing.T) {
	r := &BackupClustersConfigAction{client: newMalformedBaseURLClient(t)}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBackupClustersConfigAction_Invoke_SendError exercises BackupClustersConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBackupClustersConfigAction_Invoke_SendError(t *testing.T) {
	r := &BackupClustersConfigAction{client: newTransportErrorClient(t)}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBackupClustersConfigAction_Invoke_APIError exercises BackupClustersConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBackupClustersConfigAction_Invoke_APIError(t *testing.T) {
	r := &BackupClustersConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_backup_clusters_config")
}

// TestBackupClustersConfigAction_Invoke_APIErrorReadBody exercises BackupClustersConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBackupClustersConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BackupClustersConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := BackupClustersConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
