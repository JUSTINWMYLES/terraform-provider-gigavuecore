package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteTextConfigFileAction_Invoke_Happy exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteTextConfigFileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteTextConfigFileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteTextConfigFileAction_Invoke_NilClient exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteTextConfigFileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteTextConfigFileAction{}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteTextConfigFileAction_Invoke_BuildError exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteTextConfigFileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteTextConfigFileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteTextConfigFileAction_Invoke_SendError exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteTextConfigFileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteTextConfigFileAction{client: newTransportErrorClient(t)}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteTextConfigFileAction_Invoke_APIError exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteTextConfigFileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteTextConfigFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_text_config_file")
}

// TestDeleteTextConfigFileAction_Invoke_APIErrorReadBody exercises DeleteTextConfigFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteTextConfigFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteTextConfigFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteTextConfigFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
