package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevokeFmLicenseAction_Invoke_Happy exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevokeFmLicenseAction_Invoke_Happy(t *testing.T) {
	r := &RevokeFmLicenseAction{client: newMockClientStatus(t, 200, "{}")}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevokeFmLicenseAction_Invoke_NilClient exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevokeFmLicenseAction_Invoke_NilClient(t *testing.T) {
	r := &RevokeFmLicenseAction{}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevokeFmLicenseAction_Invoke_BuildError exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevokeFmLicenseAction_Invoke_BuildError(t *testing.T) {
	r := &RevokeFmLicenseAction{client: newMalformedBaseURLClient(t)}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevokeFmLicenseAction_Invoke_SendError exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevokeFmLicenseAction_Invoke_SendError(t *testing.T) {
	r := &RevokeFmLicenseAction{client: newTransportErrorClient(t)}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevokeFmLicenseAction_Invoke_APIError exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevokeFmLicenseAction_Invoke_APIError(t *testing.T) {
	r := &RevokeFmLicenseAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revoke_fm_license")
}

// TestRevokeFmLicenseAction_Invoke_APIErrorReadBody exercises RevokeFmLicenseAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevokeFmLicenseAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevokeFmLicenseAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevokeFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
