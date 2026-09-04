package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_Happy exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_Happy(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{client: newMockClientStatus(t, 201, "{}")}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_NilClient exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_NilClient(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_BuildError exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_BuildError(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{client: newMalformedBaseURLClient(t)}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_SendError exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_SendError(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{client: newTransportErrorClient(t)}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_APIError exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_APIError(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_set_ssl_decryption_key_store_password")
}

// TestSetSslDecryptionKeyStorePasswordAction_Invoke_APIErrorReadBody exercises SetSslDecryptionKeyStorePasswordAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSetSslDecryptionKeyStorePasswordAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SetSslDecryptionKeyStorePasswordAction{client: newMockClientReadErrorBody(t, 501)}
	m := SetSslDecryptionKeyStorePasswordActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
