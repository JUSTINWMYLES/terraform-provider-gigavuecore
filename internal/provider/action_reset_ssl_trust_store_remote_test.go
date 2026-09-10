package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetSslTrustStoreAction_Invoke_Happy exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetSslTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &ResetSslTrustStoreAction{client: newMockClientStatus(t, 201, "{}")}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetSslTrustStoreAction_Invoke_NilClient exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetSslTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &ResetSslTrustStoreAction{}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetSslTrustStoreAction_Invoke_BuildError exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetSslTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &ResetSslTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetSslTrustStoreAction_Invoke_SendError exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetSslTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &ResetSslTrustStoreAction{client: newTransportErrorClient(t)}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetSslTrustStoreAction_Invoke_APIError exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetSslTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &ResetSslTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_ssl_trust_store")
}

// TestResetSslTrustStoreAction_Invoke_APIErrorReadBody exercises ResetSslTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetSslTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetSslTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
