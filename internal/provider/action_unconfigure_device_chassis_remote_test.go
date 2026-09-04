package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUnconfigureDeviceChassisAction_Invoke_Happy exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUnconfigureDeviceChassisAction_Invoke_Happy(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{client: newMockClientStatus(t, 204, "{}")}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUnconfigureDeviceChassisAction_Invoke_NilClient exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUnconfigureDeviceChassisAction_Invoke_NilClient(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUnconfigureDeviceChassisAction_Invoke_BuildError exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUnconfigureDeviceChassisAction_Invoke_BuildError(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{client: newMalformedBaseURLClient(t)}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUnconfigureDeviceChassisAction_Invoke_SendError exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUnconfigureDeviceChassisAction_Invoke_SendError(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{client: newTransportErrorClient(t)}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUnconfigureDeviceChassisAction_Invoke_APIError exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUnconfigureDeviceChassisAction_Invoke_APIError(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_unconfigure_device_chassis")
}

// TestUnconfigureDeviceChassisAction_Invoke_APIErrorReadBody exercises UnconfigureDeviceChassisAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUnconfigureDeviceChassisAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UnconfigureDeviceChassisAction{client: newMockClientReadErrorBody(t, 501)}
	m := UnconfigureDeviceChassisActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
