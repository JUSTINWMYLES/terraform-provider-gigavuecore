package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllGtpBackupFileAction_Invoke_Happy exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllGtpBackupFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllGtpBackupFileAction_Invoke_NilClient exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllGtpBackupFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllGtpBackupFileAction_Invoke_BuildError exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllGtpBackupFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllGtpBackupFileAction_Invoke_SendError exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllGtpBackupFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{client: newTransportErrorClient(t)}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllGtpBackupFileAction_Invoke_APIError exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllGtpBackupFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_gtp_backup_file")
}

// TestDeleteAllGtpBackupFileAction_Invoke_APIErrorReadBody exercises DeleteAllGtpBackupFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllGtpBackupFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllGtpBackupFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
