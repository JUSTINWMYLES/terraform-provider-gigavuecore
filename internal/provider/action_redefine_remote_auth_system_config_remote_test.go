package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineRemoteAuthSystemConfigAction_Invoke_Happy exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineRemoteAuthSystemConfigAction_Invoke_NilClient exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineRemoteAuthSystemConfigAction_Invoke_BuildError exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineRemoteAuthSystemConfigAction_Invoke_SendError exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{client: newTransportErrorClient(t)}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineRemoteAuthSystemConfigAction_Invoke_APIError exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_remote_auth_system_config")
}

// TestRedefineRemoteAuthSystemConfigAction_Invoke_APIErrorReadBody exercises RedefineRemoteAuthSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineRemoteAuthSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineRemoteAuthSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineRemoteAuthSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
