package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSslDecryptionEndpointsAction_Invoke_Happy exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSslDecryptionEndpointsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSslDecryptionEndpointsAction_Invoke_NilClient exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSslDecryptionEndpointsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSslDecryptionEndpointsAction_Invoke_BuildError exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSslDecryptionEndpointsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSslDecryptionEndpointsAction_Invoke_SendError exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSslDecryptionEndpointsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{client: newTransportErrorClient(t)}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSslDecryptionEndpointsAction_Invoke_APIError exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSslDecryptionEndpointsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_ssl_decryption_endpoints")
}

// TestDeleteSslDecryptionEndpointsAction_Invoke_APIErrorReadBody exercises DeleteSslDecryptionEndpointsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSslDecryptionEndpointsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSslDecryptionEndpointsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSslDecryptionEndpointsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
