package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_Happy exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{client: newMockClientStatus(t, 202, "{}")}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_NilClient exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_BuildError exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_SendError exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{client: newTransportErrorClient(t)}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_APIError exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_acme_certificate_details_of_device")
}

// TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_APIErrorReadBody exercises ConfigureAcmeCertificateDetailsOfDeviceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureAcmeCertificateDetailsOfDeviceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureAcmeCertificateDetailsOfDeviceAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureAcmeCertificateDetailsOfDeviceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
