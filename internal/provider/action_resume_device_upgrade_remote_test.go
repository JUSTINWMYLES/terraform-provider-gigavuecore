package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestResumeDeviceUpgradeAction_Invoke_Happy exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestResumeDeviceUpgradeAction_Invoke_Happy(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{client: newMockClientStatus(t, 201, "{}")}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestResumeDeviceUpgradeAction_Invoke_NilClient exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestResumeDeviceUpgradeAction_Invoke_NilClient(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestResumeDeviceUpgradeAction_Invoke_BuildError exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestResumeDeviceUpgradeAction_Invoke_BuildError(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{client: newMalformedBaseURLClient(t)}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestResumeDeviceUpgradeAction_Invoke_SendError exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestResumeDeviceUpgradeAction_Invoke_SendError(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{client: newTransportErrorClient(t)}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestResumeDeviceUpgradeAction_Invoke_APIError exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestResumeDeviceUpgradeAction_Invoke_APIError(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_resume_device_upgrade")
}

// TestResumeDeviceUpgradeAction_Invoke_APIErrorReadBody exercises ResumeDeviceUpgradeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestResumeDeviceUpgradeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ResumeDeviceUpgradeAction{client: newMockClientReadErrorBody(t, 501)}
	m := ResumeDeviceUpgradeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
