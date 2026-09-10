package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllAppsProxyServerProfileAction_Invoke_Happy exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllAppsProxyServerProfileAction_Invoke_NilClient exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllAppsProxyServerProfileAction_Invoke_BuildError exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllAppsProxyServerProfileAction_Invoke_SendError exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{client: newTransportErrorClient(t)}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllAppsProxyServerProfileAction_Invoke_APIError exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_apps_proxy_server_profile")
}

// TestDeleteAllAppsProxyServerProfileAction_Invoke_APIErrorReadBody exercises DeleteAllAppsProxyServerProfileAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllAppsProxyServerProfileAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllAppsProxyServerProfileAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllAppsProxyServerProfileActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
