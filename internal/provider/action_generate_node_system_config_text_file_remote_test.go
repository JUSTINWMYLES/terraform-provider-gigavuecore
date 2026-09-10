package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGenerateNodeSystemConfigTextFileAction_Invoke_Happy exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_Happy(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{client: newMockClientStatus(t, 201, "{}")}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGenerateNodeSystemConfigTextFileAction_Invoke_NilClient exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_NilClient(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGenerateNodeSystemConfigTextFileAction_Invoke_BuildError exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_BuildError(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{client: newMalformedBaseURLClient(t)}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGenerateNodeSystemConfigTextFileAction_Invoke_SendError exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_SendError(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{client: newTransportErrorClient(t)}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGenerateNodeSystemConfigTextFileAction_Invoke_APIError exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_APIError(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_generate_node_system_config_text_file")
}

// TestGenerateNodeSystemConfigTextFileAction_Invoke_APIErrorReadBody exercises GenerateNodeSystemConfigTextFileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGenerateNodeSystemConfigTextFileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GenerateNodeSystemConfigTextFileAction{client: newMockClientReadErrorBody(t, 501)}
	m := GenerateNodeSystemConfigTextFileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
