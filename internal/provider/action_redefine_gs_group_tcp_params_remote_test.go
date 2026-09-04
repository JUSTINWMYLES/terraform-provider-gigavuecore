package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupTcpParamsAction_Invoke_Happy exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupTcpParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupTcpParamsAction_Invoke_NilClient exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupTcpParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupTcpParamsAction_Invoke_BuildError exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupTcpParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupTcpParamsAction_Invoke_SendError exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupTcpParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupTcpParamsAction_Invoke_APIError exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupTcpParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_tcp_params")
}

// TestRedefineGsGroupTcpParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupTcpParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupTcpParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupTcpParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupTcpParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
