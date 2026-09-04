package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadInlineSslTrustStoreAction_Invoke_Happy exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadInlineSslTrustStoreAction_Invoke_Happy(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadInlineSslTrustStoreAction_Invoke_NilClient exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadInlineSslTrustStoreAction_Invoke_NilClient(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadInlineSslTrustStoreAction_Invoke_BuildError exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadInlineSslTrustStoreAction_Invoke_BuildError(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{client: newMalformedBaseURLClient(t)}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadInlineSslTrustStoreAction_Invoke_SendError exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadInlineSslTrustStoreAction_Invoke_SendError(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{client: newTransportErrorClient(t)}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadInlineSslTrustStoreAction_Invoke_APIError exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadInlineSslTrustStoreAction_Invoke_APIError(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_inline_ssl_trust_store")
}

// TestUploadInlineSslTrustStoreAction_Invoke_APIErrorReadBody exercises UploadInlineSslTrustStoreAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadInlineSslTrustStoreAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadInlineSslTrustStoreAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadInlineSslTrustStoreActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
