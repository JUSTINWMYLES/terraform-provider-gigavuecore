package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_Happy exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_NilClient exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_BuildError exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_SendError exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{client: newTransportErrorClient(t)}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_APIError exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_stored_config_backup_snapshots")
}

// TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_APIErrorReadBody exercises DeleteAllStoredConfigBackupSnapshotsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllStoredConfigBackupSnapshotsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllStoredConfigBackupSnapshotsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllStoredConfigBackupSnapshotsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
