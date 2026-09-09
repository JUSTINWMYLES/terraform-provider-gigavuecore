package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddInlineSslTrustStoreCertAction_Invoke_Happy exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddInlineSslTrustStoreCertAction_Invoke_Happy(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddInlineSslTrustStoreCertAction_Invoke_NilClient exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddInlineSslTrustStoreCertAction_Invoke_NilClient(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddInlineSslTrustStoreCertAction_Invoke_BuildError exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddInlineSslTrustStoreCertAction_Invoke_BuildError(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{client: newMalformedBaseURLClient(t)}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddInlineSslTrustStoreCertAction_Invoke_SendError exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddInlineSslTrustStoreCertAction_Invoke_SendError(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{client: newTransportErrorClient(t)}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddInlineSslTrustStoreCertAction_Invoke_APIError exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddInlineSslTrustStoreCertAction_Invoke_APIError(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_inline_ssl_trust_store_cert")
}

// TestAddInlineSslTrustStoreCertAction_Invoke_APIErrorReadBody exercises AddInlineSslTrustStoreCertAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddInlineSslTrustStoreCertAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddInlineSslTrustStoreCertAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddInlineSslTrustStoreCertActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
