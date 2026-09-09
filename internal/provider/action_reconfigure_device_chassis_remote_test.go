package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReconfigureDeviceChassisAction_Invoke_Happy exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReconfigureDeviceChassisAction_Invoke_Happy(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReconfigureDeviceChassisAction_Invoke_NilClient exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReconfigureDeviceChassisAction_Invoke_NilClient(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReconfigureDeviceChassisAction_Invoke_BuildError exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReconfigureDeviceChassisAction_Invoke_BuildError(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{client: newMalformedBaseURLClient(t)}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReconfigureDeviceChassisAction_Invoke_SendError exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReconfigureDeviceChassisAction_Invoke_SendError(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{client: newTransportErrorClient(t)}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReconfigureDeviceChassisAction_Invoke_APIError exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReconfigureDeviceChassisAction_Invoke_APIError(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reconfigure_device_chassis")
}

// TestReconfigureDeviceChassisAction_Invoke_APIErrorReadBody exercises ReconfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReconfigureDeviceChassisAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReconfigureDeviceChassisAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
