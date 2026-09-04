package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRecoverInlineToolAction_Invoke_Happy exercises RecoverInlineToolAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRecoverInlineToolAction_Invoke_Happy(t *testing.T) {
	r := &RecoverInlineToolAction{client: newMockClientStatus(t, 200, "{}")}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRecoverInlineToolAction_Invoke_NilClient exercises RecoverInlineToolAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRecoverInlineToolAction_Invoke_NilClient(t *testing.T) {
	r := &RecoverInlineToolAction{}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRecoverInlineToolAction_Invoke_BuildError exercises RecoverInlineToolAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRecoverInlineToolAction_Invoke_BuildError(t *testing.T) {
	r := &RecoverInlineToolAction{client: newMalformedBaseURLClient(t)}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRecoverInlineToolAction_Invoke_SendError exercises RecoverInlineToolAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRecoverInlineToolAction_Invoke_SendError(t *testing.T) {
	r := &RecoverInlineToolAction{client: newTransportErrorClient(t)}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRecoverInlineToolAction_Invoke_APIError exercises RecoverInlineToolAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRecoverInlineToolAction_Invoke_APIError(t *testing.T) {
	r := &RecoverInlineToolAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_recover_inline_tool")
}

// TestRecoverInlineToolAction_Invoke_APIErrorReadBody exercises RecoverInlineToolAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRecoverInlineToolAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RecoverInlineToolAction{client: newMockClientReadErrorBody(t, 501)}
	m := RecoverInlineToolActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
