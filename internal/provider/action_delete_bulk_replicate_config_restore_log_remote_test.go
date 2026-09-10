package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_Happy exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_Happy(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_NilClient exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_BuildError exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{client: newMalformedBaseURLClient(t)}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_SendError exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_SendError(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{client: newTransportErrorClient(t)}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_APIError exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_APIError(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_bulk_replicate_config_restore_log")
}

// TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_APIErrorReadBody exercises DeleteBulkReplicateConfigRestoreLogAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteBulkReplicateConfigRestoreLogAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteBulkReplicateConfigRestoreLogAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteBulkReplicateConfigRestoreLogActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
