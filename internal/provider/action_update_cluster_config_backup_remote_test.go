package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateClusterConfigBackupAction_Invoke_Happy exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateClusterConfigBackupAction_Invoke_Happy(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateClusterConfigBackupAction_Invoke_NilClient exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateClusterConfigBackupAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateClusterConfigBackupAction_Invoke_BuildError exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateClusterConfigBackupAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{client: newMalformedBaseURLClient(t)}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateClusterConfigBackupAction_Invoke_SendError exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateClusterConfigBackupAction_Invoke_SendError(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{client: newTransportErrorClient(t)}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateClusterConfigBackupAction_Invoke_APIError exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateClusterConfigBackupAction_Invoke_APIError(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_cluster_config_backup")
}

// TestUpdateClusterConfigBackupAction_Invoke_APIErrorReadBody exercises UpdateClusterConfigBackupAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateClusterConfigBackupAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateClusterConfigBackupAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateClusterConfigBackupActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
