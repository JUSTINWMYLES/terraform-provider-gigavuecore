package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteClientTrustStoreAction_Invoke_Happy exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteClientTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &DeleteClientTrustStoreAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteClientTrustStoreAction_Invoke_NilClient exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteClientTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteClientTrustStoreAction{}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteClientTrustStoreAction_Invoke_BuildError exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteClientTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteClientTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteClientTrustStoreAction_Invoke_SendError exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteClientTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &DeleteClientTrustStoreAction{client: newTransportErrorClient(t)}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteClientTrustStoreAction_Invoke_APIError exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteClientTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &DeleteClientTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_client_trust_store")
}

// TestDeleteClientTrustStoreAction_Invoke_APIErrorReadBody exercises DeleteClientTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteClientTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteClientTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
