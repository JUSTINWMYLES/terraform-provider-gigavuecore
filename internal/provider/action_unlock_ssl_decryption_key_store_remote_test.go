package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnlockSslDecryptionKeyStoreAction_Invoke_Happy exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_Happy(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{client: newMockClientStatus(t, 200, "{}")}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnlockSslDecryptionKeyStoreAction_Invoke_NilClient exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_NilClient(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnlockSslDecryptionKeyStoreAction_Invoke_BuildError exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_BuildError(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{client: newMalformedBaseURLClient(t)}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnlockSslDecryptionKeyStoreAction_Invoke_SendError exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_SendError(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{client: newTransportErrorClient(t)}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnlockSslDecryptionKeyStoreAction_Invoke_APIError exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_APIError(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unlock_ssl_decryption_key_store")
}

// TestUnlockSslDecryptionKeyStoreAction_Invoke_APIErrorReadBody exercises UnlockSslDecryptionKeyStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnlockSslDecryptionKeyStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnlockSslDecryptionKeyStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnlockSslDecryptionKeyStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
