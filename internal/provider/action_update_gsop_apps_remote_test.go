package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateGsopAppsAction_Invoke_Happy exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateGsopAppsAction_Invoke_Happy(t *testing.T) {
	r := &UpdateGsopAppsAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateGsopAppsAction_Invoke_NilClient exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateGsopAppsAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateGsopAppsAction{}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateGsopAppsAction_Invoke_BuildError exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateGsopAppsAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateGsopAppsAction{client: newMalformedBaseURLClient(t)}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateGsopAppsAction_Invoke_SendError exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateGsopAppsAction_Invoke_SendError(t *testing.T) {
	r := &UpdateGsopAppsAction{client: newTransportErrorClient(t)}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateGsopAppsAction_Invoke_APIError exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateGsopAppsAction_Invoke_APIError(t *testing.T) {
	r := &UpdateGsopAppsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_gsop_apps")
}

// TestUpdateGsopAppsAction_Invoke_APIErrorReadBody exercises UpdateGsopAppsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateGsopAppsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateGsopAppsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateGsopAppsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
