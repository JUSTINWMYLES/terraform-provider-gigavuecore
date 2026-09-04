package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestConfigureDeviceChassisAction_Invoke_Happy exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestConfigureDeviceChassisAction_Invoke_Happy(t *testing.T) {
	r := &ConfigureDeviceChassisAction{client: newMockClientStatus(t, 201, "{}")}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestConfigureDeviceChassisAction_Invoke_NilClient exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestConfigureDeviceChassisAction_Invoke_NilClient(t *testing.T) {
	r := &ConfigureDeviceChassisAction{}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestConfigureDeviceChassisAction_Invoke_BuildError exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestConfigureDeviceChassisAction_Invoke_BuildError(t *testing.T) {
	r := &ConfigureDeviceChassisAction{client: newMalformedBaseURLClient(t)}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestConfigureDeviceChassisAction_Invoke_SendError exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestConfigureDeviceChassisAction_Invoke_SendError(t *testing.T) {
	r := &ConfigureDeviceChassisAction{client: newTransportErrorClient(t)}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestConfigureDeviceChassisAction_Invoke_APIError exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestConfigureDeviceChassisAction_Invoke_APIError(t *testing.T) {
	r := &ConfigureDeviceChassisAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_configure_device_chassis")
}

// TestConfigureDeviceChassisAction_Invoke_APIErrorReadBody exercises ConfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestConfigureDeviceChassisAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ConfigureDeviceChassisAction{client: newMockClientReadErrorBody(t, 501)}
	m := ConfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
