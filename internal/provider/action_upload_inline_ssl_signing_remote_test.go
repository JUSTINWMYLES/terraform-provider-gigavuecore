package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUploadInlineSslSigningAction_Invoke_Happy exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUploadInlineSslSigningAction_Invoke_Happy(t *testing.T) {
	r := &UploadInlineSslSigningAction{client: newMockClientStatus(t, 201, "{}")}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUploadInlineSslSigningAction_Invoke_NilClient exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUploadInlineSslSigningAction_Invoke_NilClient(t *testing.T) {
	r := &UploadInlineSslSigningAction{}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUploadInlineSslSigningAction_Invoke_BuildError exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUploadInlineSslSigningAction_Invoke_BuildError(t *testing.T) {
	r := &UploadInlineSslSigningAction{client: newMalformedBaseURLClient(t)}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUploadInlineSslSigningAction_Invoke_SendError exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUploadInlineSslSigningAction_Invoke_SendError(t *testing.T) {
	r := &UploadInlineSslSigningAction{client: newTransportErrorClient(t)}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUploadInlineSslSigningAction_Invoke_APIError exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUploadInlineSslSigningAction_Invoke_APIError(t *testing.T) {
	r := &UploadInlineSslSigningAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upload_inline_ssl_signing")
}

// TestUploadInlineSslSigningAction_Invoke_APIErrorReadBody exercises UploadInlineSslSigningAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUploadInlineSslSigningAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UploadInlineSslSigningAction{client: newMockClientReadErrorBody(t, 501)}
	m := UploadInlineSslSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
