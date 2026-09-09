package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateSslClientTrustStoreAction_Invoke_Happy exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateSslClientTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateSslClientTrustStoreAction_Invoke_NilClient exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateSslClientTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateSslClientTrustStoreAction_Invoke_BuildError exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateSslClientTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateSslClientTrustStoreAction_Invoke_SendError exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateSslClientTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{client: newTransportErrorClient(t)}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateSslClientTrustStoreAction_Invoke_APIError exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateSslClientTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_ssl_client_trust_store")
}

// TestCreateSslClientTrustStoreAction_Invoke_APIErrorReadBody exercises CreateSslClientTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateSslClientTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateSslClientTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateSslClientTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
