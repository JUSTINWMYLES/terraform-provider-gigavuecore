package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSystemHostBannerAction_Invoke_Happy exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSystemHostBannerAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSystemHostBannerAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSystemHostBannerAction_Invoke_NilClient exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSystemHostBannerAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSystemHostBannerAction{}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSystemHostBannerAction_Invoke_BuildError exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSystemHostBannerAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSystemHostBannerAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSystemHostBannerAction_Invoke_SendError exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSystemHostBannerAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSystemHostBannerAction{client: newTransportErrorClient(t)}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSystemHostBannerAction_Invoke_APIError exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSystemHostBannerAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSystemHostBannerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_system_host_banner")
}

// TestUpdateSystemHostBannerAction_Invoke_APIErrorReadBody exercises UpdateSystemHostBannerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSystemHostBannerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSystemHostBannerAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSystemHostBannerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
