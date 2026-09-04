package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadSslTrustStoreAction_Invoke_Happy exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadSslTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &UploadSslTrustStoreAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadSslTrustStoreAction_Invoke_NilClient exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadSslTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &UploadSslTrustStoreAction{}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadSslTrustStoreAction_Invoke_BuildError exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadSslTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &UploadSslTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadSslTrustStoreAction_Invoke_SendError exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadSslTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &UploadSslTrustStoreAction{client: newTransportErrorClient(t)}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadSslTrustStoreAction_Invoke_APIError exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadSslTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &UploadSslTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_ssl_trust_store")
}

// TestUploadSslTrustStoreAction_Invoke_APIErrorReadBody exercises UploadSslTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadSslTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadSslTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
