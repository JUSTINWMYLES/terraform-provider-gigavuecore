package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGenerateGigasmartDumpAction_Invoke_Happy exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGenerateGigasmartDumpAction_Invoke_Happy(t *testing.T) {
	r := &GenerateGigasmartDumpAction{client: newMockClientStatus(t, 201, "{}")}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGenerateGigasmartDumpAction_Invoke_NilClient exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGenerateGigasmartDumpAction_Invoke_NilClient(t *testing.T) {
	r := &GenerateGigasmartDumpAction{}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGenerateGigasmartDumpAction_Invoke_BuildError exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGenerateGigasmartDumpAction_Invoke_BuildError(t *testing.T) {
	r := &GenerateGigasmartDumpAction{client: newMalformedBaseURLClient(t)}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGenerateGigasmartDumpAction_Invoke_SendError exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGenerateGigasmartDumpAction_Invoke_SendError(t *testing.T) {
	r := &GenerateGigasmartDumpAction{client: newTransportErrorClient(t)}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGenerateGigasmartDumpAction_Invoke_APIError exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGenerateGigasmartDumpAction_Invoke_APIError(t *testing.T) {
	r := &GenerateGigasmartDumpAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_generate_gigasmart_dump")
}

// TestGenerateGigasmartDumpAction_Invoke_APIErrorReadBody exercises GenerateGigasmartDumpAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGenerateGigasmartDumpAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GenerateGigasmartDumpAction{client: newMockClientReadErrorBody(t, 501)}
	m := GenerateGigasmartDumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
