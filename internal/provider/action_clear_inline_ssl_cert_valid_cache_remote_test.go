package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearInlineSslCertValidCacheAction_Invoke_Happy exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearInlineSslCertValidCacheAction_Invoke_Happy(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearInlineSslCertValidCacheAction_Invoke_NilClient exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearInlineSslCertValidCacheAction_Invoke_NilClient(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearInlineSslCertValidCacheAction_Invoke_BuildError exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearInlineSslCertValidCacheAction_Invoke_BuildError(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{client: newMalformedBaseURLClient(t)}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearInlineSslCertValidCacheAction_Invoke_SendError exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearInlineSslCertValidCacheAction_Invoke_SendError(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{client: newTransportErrorClient(t)}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearInlineSslCertValidCacheAction_Invoke_APIError exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearInlineSslCertValidCacheAction_Invoke_APIError(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_inline_ssl_cert_valid_cache")
}

// TestClearInlineSslCertValidCacheAction_Invoke_APIErrorReadBody exercises ClearInlineSslCertValidCacheAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearInlineSslCertValidCacheAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearInlineSslCertValidCacheAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearInlineSslCertValidCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
