package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureVolumeUsageAlertEnableAction_Invoke_Happy exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{client: newMockClientStatus(t, 200, "{}")}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureVolumeUsageAlertEnableAction_Invoke_NilClient exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureVolumeUsageAlertEnableAction_Invoke_BuildError exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureVolumeUsageAlertEnableAction_Invoke_SendError exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{client: newTransportErrorClient(t)}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureVolumeUsageAlertEnableAction_Invoke_APIError exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_volume_usage_alert_enable")
}

// TestConfigureVolumeUsageAlertEnableAction_Invoke_APIErrorReadBody exercises ConfigureVolumeUsageAlertEnableAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureVolumeUsageAlertEnableAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureVolumeUsageAlertEnableAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureVolumeUsageAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
