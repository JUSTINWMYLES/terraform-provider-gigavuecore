package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestValidateDeviceUpgradeSpecAction_Invoke_Happy exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestValidateDeviceUpgradeSpecAction_Invoke_Happy(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{client: newMockClientStatus(t, 200, "{}")}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestValidateDeviceUpgradeSpecAction_Invoke_NilClient exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestValidateDeviceUpgradeSpecAction_Invoke_NilClient(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestValidateDeviceUpgradeSpecAction_Invoke_BuildError exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestValidateDeviceUpgradeSpecAction_Invoke_BuildError(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{client: newMalformedBaseURLClient(t)}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestValidateDeviceUpgradeSpecAction_Invoke_SendError exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestValidateDeviceUpgradeSpecAction_Invoke_SendError(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{client: newTransportErrorClient(t)}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestValidateDeviceUpgradeSpecAction_Invoke_APIError exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestValidateDeviceUpgradeSpecAction_Invoke_APIError(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_validate_device_upgrade_spec")
}

// TestValidateDeviceUpgradeSpecAction_Invoke_APIErrorReadBody exercises ValidateDeviceUpgradeSpecAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestValidateDeviceUpgradeSpecAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ValidateDeviceUpgradeSpecAction{client: newMockClientReadErrorBody(t, 501)}
	m := ValidateDeviceUpgradeSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
