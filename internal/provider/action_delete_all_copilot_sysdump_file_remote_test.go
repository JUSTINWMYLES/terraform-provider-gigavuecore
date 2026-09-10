package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllCopilotSysdumpFileAction_Invoke_Happy exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllCopilotSysdumpFileAction_Invoke_NilClient exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllCopilotSysdumpFileAction_Invoke_BuildError exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllCopilotSysdumpFileAction_Invoke_SendError exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{client: newTransportErrorClient(t)}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllCopilotSysdumpFileAction_Invoke_APIError exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_copilot_sysdump_file")
}

// TestDeleteAllCopilotSysdumpFileAction_Invoke_APIErrorReadBody exercises DeleteAllCopilotSysdumpFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllCopilotSysdumpFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllCopilotSysdumpFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllCopilotSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
