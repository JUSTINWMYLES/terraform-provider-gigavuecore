package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearInlineSslUrlCacheAction_Invoke_Happy exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearInlineSslUrlCacheAction_Invoke_Happy(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearInlineSslUrlCacheAction_Invoke_NilClient exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearInlineSslUrlCacheAction_Invoke_NilClient(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearInlineSslUrlCacheAction_Invoke_BuildError exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearInlineSslUrlCacheAction_Invoke_BuildError(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{client: newMalformedBaseURLClient(t)}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearInlineSslUrlCacheAction_Invoke_SendError exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearInlineSslUrlCacheAction_Invoke_SendError(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{client: newTransportErrorClient(t)}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearInlineSslUrlCacheAction_Invoke_APIError exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearInlineSslUrlCacheAction_Invoke_APIError(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_inline_ssl_url_cache")
}

// TestClearInlineSslUrlCacheAction_Invoke_APIErrorReadBody exercises ClearInlineSslUrlCacheAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearInlineSslUrlCacheAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearInlineSslUrlCacheAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearInlineSslUrlCacheActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
