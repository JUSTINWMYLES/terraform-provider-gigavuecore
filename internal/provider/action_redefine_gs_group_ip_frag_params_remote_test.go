package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupIpFragParamsAction_Invoke_Happy exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupIpFragParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupIpFragParamsAction_Invoke_NilClient exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupIpFragParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupIpFragParamsAction_Invoke_BuildError exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupIpFragParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupIpFragParamsAction_Invoke_SendError exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupIpFragParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupIpFragParamsAction_Invoke_APIError exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupIpFragParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_ip_frag_params")
}

// TestRedefineGsGroupIpFragParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupIpFragParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupIpFragParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupIpFragParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupIpFragParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
