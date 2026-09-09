package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetInlineSslTrustStoreAction_Invoke_Happy exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetInlineSslTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{client: newMockClientStatus(t, 201, "{}")}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetInlineSslTrustStoreAction_Invoke_NilClient exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetInlineSslTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetInlineSslTrustStoreAction_Invoke_BuildError exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetInlineSslTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetInlineSslTrustStoreAction_Invoke_SendError exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetInlineSslTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{client: newTransportErrorClient(t)}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetInlineSslTrustStoreAction_Invoke_APIError exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetInlineSslTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_inline_ssl_trust_store")
}

// TestResetInlineSslTrustStoreAction_Invoke_APIErrorReadBody exercises ResetInlineSslTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetInlineSslTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetInlineSslTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
