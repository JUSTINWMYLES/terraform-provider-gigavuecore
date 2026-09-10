package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveFmLicenseAction_Invoke_Happy exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveFmLicenseAction_Invoke_Happy(t *testing.T) {
	r := &RemoveFmLicenseAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveFmLicenseAction_Invoke_NilClient exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveFmLicenseAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveFmLicenseAction{}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveFmLicenseAction_Invoke_BuildError exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveFmLicenseAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveFmLicenseAction{client: newMalformedBaseURLClient(t)}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveFmLicenseAction_Invoke_SendError exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveFmLicenseAction_Invoke_SendError(t *testing.T) {
	r := &RemoveFmLicenseAction{client: newTransportErrorClient(t)}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveFmLicenseAction_Invoke_APIError exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveFmLicenseAction_Invoke_APIError(t *testing.T) {
	r := &RemoveFmLicenseAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_fm_license")
}

// TestRemoveFmLicenseAction_Invoke_APIErrorReadBody exercises RemoveFmLicenseAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveFmLicenseAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveFmLicenseAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveFmLicenseActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
