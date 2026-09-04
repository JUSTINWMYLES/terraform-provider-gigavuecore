package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSipPortsAction_Invoke_Happy exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSipPortsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSipPortsAction_Invoke_NilClient exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSipPortsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSipPortsAction_Invoke_BuildError exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSipPortsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSipPortsAction_Invoke_SendError exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSipPortsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSipPortsAction_Invoke_APIError exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSipPortsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sip_ports")
}

// TestRedefineGsGroupSipPortsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSipPortsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSipPortsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSipPortsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSipPortsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
