package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteKeystoreKeysAction_Invoke_Happy exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteKeystoreKeysAction_Invoke_Happy(t *testing.T) {
	r := &DeleteKeystoreKeysAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteKeystoreKeysAction_Invoke_NilClient exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteKeystoreKeysAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteKeystoreKeysAction{}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteKeystoreKeysAction_Invoke_BuildError exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteKeystoreKeysAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteKeystoreKeysAction{client: newMalformedBaseURLClient(t)}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteKeystoreKeysAction_Invoke_SendError exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteKeystoreKeysAction_Invoke_SendError(t *testing.T) {
	r := &DeleteKeystoreKeysAction{client: newTransportErrorClient(t)}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteKeystoreKeysAction_Invoke_APIError exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteKeystoreKeysAction_Invoke_APIError(t *testing.T) {
	r := &DeleteKeystoreKeysAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_keystore_keys")
}

// TestDeleteKeystoreKeysAction_Invoke_APIErrorReadBody exercises DeleteKeystoreKeysAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteKeystoreKeysAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteKeystoreKeysAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteKeystoreKeysActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
