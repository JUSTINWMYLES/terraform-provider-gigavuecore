package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRetryDeviceUpgradeAction_Invoke_Happy exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRetryDeviceUpgradeAction_Invoke_Happy(t *testing.T) {
	r := &RetryDeviceUpgradeAction{client: newMockClientStatus(t, 201, "{}")}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRetryDeviceUpgradeAction_Invoke_NilClient exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRetryDeviceUpgradeAction_Invoke_NilClient(t *testing.T) {
	r := &RetryDeviceUpgradeAction{}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRetryDeviceUpgradeAction_Invoke_BuildError exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRetryDeviceUpgradeAction_Invoke_BuildError(t *testing.T) {
	r := &RetryDeviceUpgradeAction{client: newMalformedBaseURLClient(t)}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRetryDeviceUpgradeAction_Invoke_SendError exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRetryDeviceUpgradeAction_Invoke_SendError(t *testing.T) {
	r := &RetryDeviceUpgradeAction{client: newTransportErrorClient(t)}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRetryDeviceUpgradeAction_Invoke_APIError exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRetryDeviceUpgradeAction_Invoke_APIError(t *testing.T) {
	r := &RetryDeviceUpgradeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_retry_device_upgrade")
}

// TestRetryDeviceUpgradeAction_Invoke_APIErrorReadBody exercises RetryDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRetryDeviceUpgradeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RetryDeviceUpgradeAction{client: newMockClientReadErrorBody(t, 501)}
	m := RetryDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
