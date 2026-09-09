package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteGtpBackupFileAction_Invoke_Happy exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteGtpBackupFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteGtpBackupFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteGtpBackupFileAction_Invoke_NilClient exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteGtpBackupFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteGtpBackupFileAction{}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteGtpBackupFileAction_Invoke_BuildError exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteGtpBackupFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteGtpBackupFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteGtpBackupFileAction_Invoke_SendError exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteGtpBackupFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteGtpBackupFileAction{client: newTransportErrorClient(t)}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteGtpBackupFileAction_Invoke_APIError exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteGtpBackupFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteGtpBackupFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_gtp_backup_file")
}

// TestDeleteGtpBackupFileAction_Invoke_APIErrorReadBody exercises DeleteGtpBackupFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteGtpBackupFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteGtpBackupFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteGtpBackupFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
