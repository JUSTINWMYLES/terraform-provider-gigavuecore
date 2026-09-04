package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateDevicePortAction_Invoke_Happy exercises UpdateDevicePortAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateDevicePortAction_Invoke_Happy(t *testing.T) {
	r := &UpdateDevicePortAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateDevicePortAction_Invoke_NilClient exercises UpdateDevicePortAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateDevicePortAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateDevicePortAction{}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateDevicePortAction_Invoke_BuildError exercises UpdateDevicePortAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateDevicePortAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateDevicePortAction{client: newMalformedBaseURLClient(t)}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateDevicePortAction_Invoke_SendError exercises UpdateDevicePortAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateDevicePortAction_Invoke_SendError(t *testing.T) {
	r := &UpdateDevicePortAction{client: newTransportErrorClient(t)}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateDevicePortAction_Invoke_APIError exercises UpdateDevicePortAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateDevicePortAction_Invoke_APIError(t *testing.T) {
	r := &UpdateDevicePortAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_device_port")
}

// TestUpdateDevicePortAction_Invoke_APIErrorReadBody exercises UpdateDevicePortAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateDevicePortAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateDevicePortAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateDevicePortActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
