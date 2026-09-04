package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSslDecryptionKeyMappingAction_Invoke_Happy exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSslDecryptionKeyMappingAction_Invoke_NilClient exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSslDecryptionKeyMappingAction_Invoke_BuildError exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSslDecryptionKeyMappingAction_Invoke_SendError exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{client: newTransportErrorClient(t)}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSslDecryptionKeyMappingAction_Invoke_APIError exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_ssl_decryption_key_mapping")
}

// TestDeleteSslDecryptionKeyMappingAction_Invoke_APIErrorReadBody exercises DeleteSslDecryptionKeyMappingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSslDecryptionKeyMappingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSslDecryptionKeyMappingAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSslDecryptionKeyMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
