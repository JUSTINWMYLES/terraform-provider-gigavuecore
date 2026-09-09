package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSipSessionAction_Invoke_Happy exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSipSessionAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSipSessionAction_Invoke_NilClient exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSipSessionAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSipSessionAction_Invoke_BuildError exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSipSessionAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSipSessionAction_Invoke_SendError exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSipSessionAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSipSessionAction_Invoke_APIError exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSipSessionAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sip_session")
}

// TestRedefineGsGroupSipSessionAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSipSessionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSipSessionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSipSessionAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSipSessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
