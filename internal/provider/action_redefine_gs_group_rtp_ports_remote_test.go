package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupRtpPortsAction_Invoke_Happy exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupRtpPortsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupRtpPortsAction_Invoke_NilClient exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupRtpPortsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupRtpPortsAction_Invoke_BuildError exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupRtpPortsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupRtpPortsAction_Invoke_SendError exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupRtpPortsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupRtpPortsAction_Invoke_APIError exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupRtpPortsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_rtp_ports")
}

// TestRedefineGsGroupRtpPortsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupRtpPortsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupRtpPortsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupRtpPortsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupRtpPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
