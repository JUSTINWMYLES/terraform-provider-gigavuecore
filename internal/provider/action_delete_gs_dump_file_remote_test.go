package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteGsDumpFileAction_Invoke_Happy exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteGsDumpFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteGsDumpFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteGsDumpFileAction_Invoke_NilClient exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteGsDumpFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteGsDumpFileAction{}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteGsDumpFileAction_Invoke_BuildError exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteGsDumpFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteGsDumpFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteGsDumpFileAction_Invoke_SendError exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteGsDumpFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteGsDumpFileAction{client: newTransportErrorClient(t)}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteGsDumpFileAction_Invoke_APIError exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteGsDumpFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteGsDumpFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_gs_dump_file")
}

// TestDeleteGsDumpFileAction_Invoke_APIErrorReadBody exercises DeleteGsDumpFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteGsDumpFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteGsDumpFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteGsDumpFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
