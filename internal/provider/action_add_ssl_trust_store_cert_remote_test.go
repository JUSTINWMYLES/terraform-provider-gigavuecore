package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddSslTrustStoreCertAction_Invoke_Happy exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddSslTrustStoreCertAction_Invoke_Happy(t *testing.T) {
	r := &AddSslTrustStoreCertAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddSslTrustStoreCertAction_Invoke_NilClient exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddSslTrustStoreCertAction_Invoke_NilClient(t *testing.T) {
	r := &AddSslTrustStoreCertAction{}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddSslTrustStoreCertAction_Invoke_BuildError exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddSslTrustStoreCertAction_Invoke_BuildError(t *testing.T) {
	r := &AddSslTrustStoreCertAction{client: newMalformedBaseURLClient(t)}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddSslTrustStoreCertAction_Invoke_SendError exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddSslTrustStoreCertAction_Invoke_SendError(t *testing.T) {
	r := &AddSslTrustStoreCertAction{client: newTransportErrorClient(t)}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddSslTrustStoreCertAction_Invoke_APIError exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddSslTrustStoreCertAction_Invoke_APIError(t *testing.T) {
	r := &AddSslTrustStoreCertAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_ssl_trust_store_cert")
}

// TestAddSslTrustStoreCertAction_Invoke_APIErrorReadBody exercises AddSslTrustStoreCertAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddSslTrustStoreCertAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddSslTrustStoreCertAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
