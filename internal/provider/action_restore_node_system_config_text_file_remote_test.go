package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRestoreNodeSystemConfigTextFileAction_Invoke_Happy exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_Happy(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRestoreNodeSystemConfigTextFileAction_Invoke_NilClient exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_NilClient(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRestoreNodeSystemConfigTextFileAction_Invoke_BuildError exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_BuildError(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{client: newMalformedBaseURLClient(t)}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRestoreNodeSystemConfigTextFileAction_Invoke_SendError exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_SendError(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{client: newTransportErrorClient(t)}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRestoreNodeSystemConfigTextFileAction_Invoke_APIError exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_APIError(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_restore_node_system_config_text_file")
}

// TestRestoreNodeSystemConfigTextFileAction_Invoke_APIErrorReadBody exercises RestoreNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRestoreNodeSystemConfigTextFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RestoreNodeSystemConfigTextFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := RestoreNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
