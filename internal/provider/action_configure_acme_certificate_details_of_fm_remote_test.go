package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_Happy exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{client: newMockClientStatus(t, 202, "{}")}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_NilClient exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_BuildError exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_SendError exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{client: newTransportErrorClient(t)}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_APIError exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_acme_certificate_details_of_fm")
}

// TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_APIErrorReadBody exercises ConfigureAcmeCertificateDetailsOfFmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureAcmeCertificateDetailsOfFmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfFmAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureAcmeCertificateDetailsOfFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
