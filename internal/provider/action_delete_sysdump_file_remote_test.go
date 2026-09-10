package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSysdumpFileAction_Invoke_Happy exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSysdumpFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSysdumpFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSysdumpFileAction_Invoke_NilClient exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSysdumpFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSysdumpFileAction{}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSysdumpFileAction_Invoke_BuildError exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSysdumpFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSysdumpFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSysdumpFileAction_Invoke_SendError exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSysdumpFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSysdumpFileAction{client: newTransportErrorClient(t)}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSysdumpFileAction_Invoke_APIError exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSysdumpFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSysdumpFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_sysdump_file")
}

// TestDeleteSysdumpFileAction_Invoke_APIErrorReadBody exercises DeleteSysdumpFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSysdumpFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSysdumpFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSysdumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
