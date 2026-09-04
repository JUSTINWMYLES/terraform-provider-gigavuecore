package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceSslClientTrustStoreCertAction_Invoke_Happy exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceSslClientTrustStoreCertAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{client: newMockClientStatus(t, 201, "{}")}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceSslClientTrustStoreCertAction_Invoke_NilClient exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceSslClientTrustStoreCertAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceSslClientTrustStoreCertAction_Invoke_BuildError exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceSslClientTrustStoreCertAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceSslClientTrustStoreCertAction_Invoke_SendError exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceSslClientTrustStoreCertAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{client: newTransportErrorClient(t)}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceSslClientTrustStoreCertAction_Invoke_APIError exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceSslClientTrustStoreCertAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_ssl_client_trust_store_cert")
}

// TestReplaceSslClientTrustStoreCertAction_Invoke_APIErrorReadBody exercises ReplaceSslClientTrustStoreCertAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceSslClientTrustStoreCertAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceSslClientTrustStoreCertAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
