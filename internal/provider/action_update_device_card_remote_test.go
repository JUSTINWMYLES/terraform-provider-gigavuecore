package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateDeviceCardAction_Invoke_Happy exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateDeviceCardAction_Invoke_Happy(t *testing.T) {
	r := &UpdateDeviceCardAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateDeviceCardAction_Invoke_NilClient exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateDeviceCardAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateDeviceCardAction{}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateDeviceCardAction_Invoke_BuildError exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateDeviceCardAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateDeviceCardAction{client: newMalformedBaseURLClient(t)}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateDeviceCardAction_Invoke_SendError exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateDeviceCardAction_Invoke_SendError(t *testing.T) {
	r := &UpdateDeviceCardAction{client: newTransportErrorClient(t)}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateDeviceCardAction_Invoke_APIError exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateDeviceCardAction_Invoke_APIError(t *testing.T) {
	r := &UpdateDeviceCardAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_device_card")
}

// TestUpdateDeviceCardAction_Invoke_APIErrorReadBody exercises UpdateDeviceCardAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateDeviceCardAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateDeviceCardAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateDeviceCardActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
