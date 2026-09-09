package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceOldFmBindingsAction_Invoke_Happy exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceOldFmBindingsAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceOldFmBindingsAction_Invoke_NilClient exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceOldFmBindingsAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceOldFmBindingsAction_Invoke_BuildError exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceOldFmBindingsAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceOldFmBindingsAction_Invoke_SendError exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceOldFmBindingsAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{client: newTransportErrorClient(t)}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceOldFmBindingsAction_Invoke_APIError exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceOldFmBindingsAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_old_fm_bindings")
}

// TestReplaceOldFmBindingsAction_Invoke_APIErrorReadBody exercises ReplaceOldFmBindingsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceOldFmBindingsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceOldFmBindingsAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceOldFmBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
