package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResetDeviceUpgradeAction_Invoke_Happy exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResetDeviceUpgradeAction_Invoke_Happy(t *testing.T) {
	r := &ResetDeviceUpgradeAction{client: newMockClientStatus(t, 201, "{}")}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResetDeviceUpgradeAction_Invoke_NilClient exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResetDeviceUpgradeAction_Invoke_NilClient(t *testing.T) {
	r := &ResetDeviceUpgradeAction{}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResetDeviceUpgradeAction_Invoke_BuildError exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResetDeviceUpgradeAction_Invoke_BuildError(t *testing.T) {
	r := &ResetDeviceUpgradeAction{client: newMalformedBaseURLClient(t)}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResetDeviceUpgradeAction_Invoke_SendError exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResetDeviceUpgradeAction_Invoke_SendError(t *testing.T) {
	r := &ResetDeviceUpgradeAction{client: newTransportErrorClient(t)}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResetDeviceUpgradeAction_Invoke_APIError exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResetDeviceUpgradeAction_Invoke_APIError(t *testing.T) {
	r := &ResetDeviceUpgradeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reset_device_upgrade")
}

// TestResetDeviceUpgradeAction_Invoke_APIErrorReadBody exercises ResetDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResetDeviceUpgradeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResetDeviceUpgradeAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResetDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
