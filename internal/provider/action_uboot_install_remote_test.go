package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUbootInstallAction_Invoke_Happy exercises UbootInstallAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUbootInstallAction_Invoke_Happy(t *testing.T) {
	r := &UbootInstallAction{client: newMockClientStatus(t, 201, "{}")}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUbootInstallAction_Invoke_NilClient exercises UbootInstallAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUbootInstallAction_Invoke_NilClient(t *testing.T) {
	r := &UbootInstallAction{}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUbootInstallAction_Invoke_BuildError exercises UbootInstallAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUbootInstallAction_Invoke_BuildError(t *testing.T) {
	r := &UbootInstallAction{client: newMalformedBaseURLClient(t)}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUbootInstallAction_Invoke_SendError exercises UbootInstallAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUbootInstallAction_Invoke_SendError(t *testing.T) {
	r := &UbootInstallAction{client: newTransportErrorClient(t)}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUbootInstallAction_Invoke_APIError exercises UbootInstallAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUbootInstallAction_Invoke_APIError(t *testing.T) {
	r := &UbootInstallAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_uboot_install")
}

// TestUbootInstallAction_Invoke_APIErrorReadBody exercises UbootInstallAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUbootInstallAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UbootInstallAction{client: newMockClientReadErrorBody(t, 501)}
	m := UbootInstallActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
