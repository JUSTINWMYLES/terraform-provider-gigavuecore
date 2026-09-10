package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAppendSslClientTrustStoreCertAction_Invoke_Happy exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAppendSslClientTrustStoreCertAction_Invoke_Happy(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{client: newMockClientStatus(t, 201, "{}")}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAppendSslClientTrustStoreCertAction_Invoke_NilClient exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAppendSslClientTrustStoreCertAction_Invoke_NilClient(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAppendSslClientTrustStoreCertAction_Invoke_BuildError exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAppendSslClientTrustStoreCertAction_Invoke_BuildError(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{client: newMalformedBaseURLClient(t)}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAppendSslClientTrustStoreCertAction_Invoke_SendError exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAppendSslClientTrustStoreCertAction_Invoke_SendError(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{client: newTransportErrorClient(t)}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAppendSslClientTrustStoreCertAction_Invoke_APIError exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAppendSslClientTrustStoreCertAction_Invoke_APIError(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_append_ssl_client_trust_store_cert")
}

// TestAppendSslClientTrustStoreCertAction_Invoke_APIErrorReadBody exercises AppendSslClientTrustStoreCertAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAppendSslClientTrustStoreCertAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AppendSslClientTrustStoreCertAction{client: newMockClientReadErrorBody(t, 501)}
	m := AppendSslClientTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
