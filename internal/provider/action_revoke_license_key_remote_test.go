package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRevokeLicenseKeyAction_Invoke_Happy exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRevokeLicenseKeyAction_Invoke_Happy(t *testing.T) {
	r := &RevokeLicenseKeyAction{client: newMockClientStatus(t, 200, "{}")}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRevokeLicenseKeyAction_Invoke_NilClient exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRevokeLicenseKeyAction_Invoke_NilClient(t *testing.T) {
	r := &RevokeLicenseKeyAction{}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRevokeLicenseKeyAction_Invoke_BuildError exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRevokeLicenseKeyAction_Invoke_BuildError(t *testing.T) {
	r := &RevokeLicenseKeyAction{client: newMalformedBaseURLClient(t)}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRevokeLicenseKeyAction_Invoke_SendError exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRevokeLicenseKeyAction_Invoke_SendError(t *testing.T) {
	r := &RevokeLicenseKeyAction{client: newTransportErrorClient(t)}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRevokeLicenseKeyAction_Invoke_APIError exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRevokeLicenseKeyAction_Invoke_APIError(t *testing.T) {
	r := &RevokeLicenseKeyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_revoke_license_key")
}

// TestRevokeLicenseKeyAction_Invoke_APIErrorReadBody exercises RevokeLicenseKeyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRevokeLicenseKeyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RevokeLicenseKeyAction{client: newMockClientReadErrorBody(t, 501)}
	m := RevokeLicenseKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
