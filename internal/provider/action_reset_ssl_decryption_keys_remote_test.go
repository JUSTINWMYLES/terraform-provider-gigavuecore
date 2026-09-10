package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetSslDecryptionKeysAction_Invoke_Happy exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetSslDecryptionKeysAction_Invoke_Happy(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{client: newMockClientStatus(t, 200, "{}")}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetSslDecryptionKeysAction_Invoke_NilClient exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetSslDecryptionKeysAction_Invoke_NilClient(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetSslDecryptionKeysAction_Invoke_BuildError exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetSslDecryptionKeysAction_Invoke_BuildError(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{client: newMalformedBaseURLClient(t)}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetSslDecryptionKeysAction_Invoke_SendError exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetSslDecryptionKeysAction_Invoke_SendError(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{client: newTransportErrorClient(t)}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetSslDecryptionKeysAction_Invoke_APIError exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetSslDecryptionKeysAction_Invoke_APIError(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_ssl_decryption_keys")
}

// TestResetSslDecryptionKeysAction_Invoke_APIErrorReadBody exercises ResetSslDecryptionKeysAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetSslDecryptionKeysAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetSslDecryptionKeysAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetSslDecryptionKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
