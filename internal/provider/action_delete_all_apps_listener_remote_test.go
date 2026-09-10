package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllAppsListenerAction_Invoke_Happy exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllAppsListenerAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllAppsListenerAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllAppsListenerAction_Invoke_NilClient exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllAppsListenerAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllAppsListenerAction{}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllAppsListenerAction_Invoke_BuildError exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllAppsListenerAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllAppsListenerAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllAppsListenerAction_Invoke_SendError exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllAppsListenerAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllAppsListenerAction{client: newTransportErrorClient(t)}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllAppsListenerAction_Invoke_APIError exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllAppsListenerAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllAppsListenerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_apps_listener")
}

// TestDeleteAllAppsListenerAction_Invoke_APIErrorReadBody exercises DeleteAllAppsListenerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllAppsListenerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllAppsListenerAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllAppsListenerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
