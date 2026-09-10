package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_Happy exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_Happy(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{client: newMockClientStatus(t, 200, "{}")}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_NilClient exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_NilClient(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_BuildError exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_BuildError(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{client: newMalformedBaseURLClient(t)}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_SendError exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_SendError(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{client: newTransportErrorClient(t)}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_APIError exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_APIError(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_restore_cluater_config_backup_snapshot")
}

// TestRestoreCluaterConfigBackupSnapshotAction_Invoke_APIErrorReadBody exercises RestoreCluaterConfigBackupSnapshotAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRestoreCluaterConfigBackupSnapshotAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RestoreCluaterConfigBackupSnapshotAction{client: newMockClientReadErrorBody(t, 501)}
	m := RestoreCluaterConfigBackupSnapshotActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
