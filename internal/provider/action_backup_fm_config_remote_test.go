package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBackupFmConfigAction_Invoke_Happy exercises BackupFmConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBackupFmConfigAction_Invoke_Happy(t *testing.T) {
	r := &BackupFmConfigAction{client: newMockClientStatus(t, 201, "{}")}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBackupFmConfigAction_Invoke_NilClient exercises BackupFmConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBackupFmConfigAction_Invoke_NilClient(t *testing.T) {
	r := &BackupFmConfigAction{}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBackupFmConfigAction_Invoke_BuildError exercises BackupFmConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBackupFmConfigAction_Invoke_BuildError(t *testing.T) {
	r := &BackupFmConfigAction{client: newMalformedBaseURLClient(t)}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBackupFmConfigAction_Invoke_SendError exercises BackupFmConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBackupFmConfigAction_Invoke_SendError(t *testing.T) {
	r := &BackupFmConfigAction{client: newTransportErrorClient(t)}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBackupFmConfigAction_Invoke_APIError exercises BackupFmConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBackupFmConfigAction_Invoke_APIError(t *testing.T) {
	r := &BackupFmConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_backup_fm_config")
}

// TestBackupFmConfigAction_Invoke_APIErrorReadBody exercises BackupFmConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBackupFmConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BackupFmConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := BackupFmConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
