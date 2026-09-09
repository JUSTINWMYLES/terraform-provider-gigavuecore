package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_Happy exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_Happy(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{client: newMockClientStatus(t, 201, "{}")}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_NilClient exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_NilClient(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_BuildError exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_BuildError(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{client: newMalformedBaseURLClient(t)}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_SendError exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_SendError(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{client: newTransportErrorClient(t)}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_APIError exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_APIError(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_install_node_locked_license_key_on_node")
}

// TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_APIErrorReadBody exercises InstallNodeLockedLicenseKeyOnNodeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInstallNodeLockedLicenseKeyOnNodeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &InstallNodeLockedLicenseKeyOnNodeAction{client: newMockClientReadErrorBody(t, 501)}
	m := InstallNodeLockedLicenseKeyOnNodeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
