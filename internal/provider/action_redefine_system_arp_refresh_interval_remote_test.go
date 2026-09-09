package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSystemArpRefreshIntervalAction_Invoke_Happy exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSystemArpRefreshIntervalAction_Invoke_NilClient exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSystemArpRefreshIntervalAction_Invoke_BuildError exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSystemArpRefreshIntervalAction_Invoke_SendError exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{client: newTransportErrorClient(t)}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSystemArpRefreshIntervalAction_Invoke_APIError exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_system_arp_refresh_interval")
}

// TestRedefineSystemArpRefreshIntervalAction_Invoke_APIErrorReadBody exercises RedefineSystemArpRefreshIntervalAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSystemArpRefreshIntervalAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSystemArpRefreshIntervalAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSystemArpRefreshIntervalActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
