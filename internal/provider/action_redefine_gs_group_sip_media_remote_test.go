package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSipMediaAction_Invoke_Happy exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSipMediaAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSipMediaAction_Invoke_NilClient exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSipMediaAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSipMediaAction_Invoke_BuildError exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSipMediaAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSipMediaAction_Invoke_SendError exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSipMediaAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSipMediaAction_Invoke_APIError exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSipMediaAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sip_media")
}

// TestRedefineGsGroupSipMediaAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSipMediaAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSipMediaAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSipMediaAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSipMediaActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
