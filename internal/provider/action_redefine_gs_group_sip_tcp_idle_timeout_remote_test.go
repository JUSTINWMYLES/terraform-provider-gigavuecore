package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_Happy exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_NilClient exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_BuildError exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_SendError exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_APIError exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sip_tcp_idle_timeout")
}

// TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSipTcpIdleTimeoutAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSipTcpIdleTimeoutAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSipTcpIdleTimeoutAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSipTcpIdleTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
