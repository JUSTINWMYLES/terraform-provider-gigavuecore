package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureExpiryAlertEnableAction_Invoke_Happy exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureExpiryAlertEnableAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{client: newMockClientStatus(t, 200, "{}")}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureExpiryAlertEnableAction_Invoke_NilClient exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureExpiryAlertEnableAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureExpiryAlertEnableAction_Invoke_BuildError exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureExpiryAlertEnableAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureExpiryAlertEnableAction_Invoke_SendError exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureExpiryAlertEnableAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{client: newTransportErrorClient(t)}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureExpiryAlertEnableAction_Invoke_APIError exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureExpiryAlertEnableAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_expiry_alert_enable")
}

// TestConfigureExpiryAlertEnableAction_Invoke_APIErrorReadBody exercises ConfigureExpiryAlertEnableAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureExpiryAlertEnableAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureExpiryAlertEnableAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureExpiryAlertEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
