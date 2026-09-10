package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddFmLicenseAction_Invoke_Happy exercises AddFmLicenseAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddFmLicenseAction_Invoke_Happy(t *testing.T) {
	r := &AddFmLicenseAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddFmLicenseAction_Invoke_NilClient exercises AddFmLicenseAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddFmLicenseAction_Invoke_NilClient(t *testing.T) {
	r := &AddFmLicenseAction{}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddFmLicenseAction_Invoke_BuildError exercises AddFmLicenseAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddFmLicenseAction_Invoke_BuildError(t *testing.T) {
	r := &AddFmLicenseAction{client: newMalformedBaseURLClient(t)}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddFmLicenseAction_Invoke_SendError exercises AddFmLicenseAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddFmLicenseAction_Invoke_SendError(t *testing.T) {
	r := &AddFmLicenseAction{client: newTransportErrorClient(t)}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddFmLicenseAction_Invoke_APIError exercises AddFmLicenseAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddFmLicenseAction_Invoke_APIError(t *testing.T) {
	r := &AddFmLicenseAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_fm_license")
}

// TestAddFmLicenseAction_Invoke_APIErrorReadBody exercises AddFmLicenseAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddFmLicenseAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddFmLicenseAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
