package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigDiskUsageThresholdAction_Invoke_Happy exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigDiskUsageThresholdAction_Invoke_Happy(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{client: newMockClientStatus(t, 200, "{}")}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigDiskUsageThresholdAction_Invoke_NilClient exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigDiskUsageThresholdAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigDiskUsageThresholdAction_Invoke_BuildError exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigDiskUsageThresholdAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{client: newMalformedBaseURLClient(t)}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigDiskUsageThresholdAction_Invoke_SendError exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigDiskUsageThresholdAction_Invoke_SendError(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{client: newTransportErrorClient(t)}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigDiskUsageThresholdAction_Invoke_APIError exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigDiskUsageThresholdAction_Invoke_APIError(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_config_disk_usage_threshold")
}

// TestConfigDiskUsageThresholdAction_Invoke_APIErrorReadBody exercises ConfigDiskUsageThresholdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigDiskUsageThresholdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigDiskUsageThresholdAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigDiskUsageThresholdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
