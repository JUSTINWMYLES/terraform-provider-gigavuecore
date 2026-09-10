package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestGenerateSysdumpAction_Invoke_Happy exercises GenerateSysdumpAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestGenerateSysdumpAction_Invoke_Happy(t *testing.T) {
	r := &GenerateSysdumpAction{client: newMockClientStatus(t, 201, "{}")}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGenerateSysdumpAction_Invoke_NilClient exercises GenerateSysdumpAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGenerateSysdumpAction_Invoke_NilClient(t *testing.T) {
	r := &GenerateSysdumpAction{}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGenerateSysdumpAction_Invoke_BuildError exercises GenerateSysdumpAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGenerateSysdumpAction_Invoke_BuildError(t *testing.T) {
	r := &GenerateSysdumpAction{client: newMalformedBaseURLClient(t)}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGenerateSysdumpAction_Invoke_SendError exercises GenerateSysdumpAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestGenerateSysdumpAction_Invoke_SendError(t *testing.T) {
	r := &GenerateSysdumpAction{client: newTransportErrorClient(t)}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGenerateSysdumpAction_Invoke_APIError exercises GenerateSysdumpAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGenerateSysdumpAction_Invoke_APIError(t *testing.T) {
	r := &GenerateSysdumpAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_generate_sysdump")
}

// TestGenerateSysdumpAction_Invoke_APIErrorReadBody exercises GenerateSysdumpAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGenerateSysdumpAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &GenerateSysdumpAction{client: newMockClientReadErrorBody(t, 501)}
	m := GenerateSysdumpActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
