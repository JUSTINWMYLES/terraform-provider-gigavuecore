package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpgradeFmImageAction_Invoke_Happy exercises UpgradeFmImageAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpgradeFmImageAction_Invoke_Happy(t *testing.T) {
	r := &UpgradeFmImageAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpgradeFmImageAction_Invoke_NilClient exercises UpgradeFmImageAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpgradeFmImageAction_Invoke_NilClient(t *testing.T) {
	r := &UpgradeFmImageAction{}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpgradeFmImageAction_Invoke_BuildError exercises UpgradeFmImageAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpgradeFmImageAction_Invoke_BuildError(t *testing.T) {
	r := &UpgradeFmImageAction{client: newMalformedBaseURLClient(t)}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpgradeFmImageAction_Invoke_SendError exercises UpgradeFmImageAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpgradeFmImageAction_Invoke_SendError(t *testing.T) {
	r := &UpgradeFmImageAction{client: newTransportErrorClient(t)}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpgradeFmImageAction_Invoke_APIError exercises UpgradeFmImageAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpgradeFmImageAction_Invoke_APIError(t *testing.T) {
	r := &UpgradeFmImageAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_upgrade_fm_image")
}

// TestUpgradeFmImageAction_Invoke_APIErrorReadBody exercises UpgradeFmImageAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpgradeFmImageAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpgradeFmImageAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpgradeFmImageActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
