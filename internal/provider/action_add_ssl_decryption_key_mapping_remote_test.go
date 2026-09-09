package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddSslDecryptionKeyMappingAction_Invoke_Happy exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddSslDecryptionKeyMappingAction_Invoke_Happy(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddSslDecryptionKeyMappingAction_Invoke_NilClient exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddSslDecryptionKeyMappingAction_Invoke_NilClient(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddSslDecryptionKeyMappingAction_Invoke_BuildError exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddSslDecryptionKeyMappingAction_Invoke_BuildError(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{client: newMalformedBaseURLClient(t)}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddSslDecryptionKeyMappingAction_Invoke_SendError exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddSslDecryptionKeyMappingAction_Invoke_SendError(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{client: newTransportErrorClient(t)}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddSslDecryptionKeyMappingAction_Invoke_APIError exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddSslDecryptionKeyMappingAction_Invoke_APIError(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_ssl_decryption_key_mapping")
}

// TestAddSslDecryptionKeyMappingAction_Invoke_APIErrorReadBody exercises AddSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddSslDecryptionKeyMappingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddSslDecryptionKeyMappingAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
