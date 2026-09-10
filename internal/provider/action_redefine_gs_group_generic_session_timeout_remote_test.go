package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_Happy exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_NilClient exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_BuildError exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_SendError exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_APIError exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_generic_session_timeout")
}

// TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGenericSessionTimeoutAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGenericSessionTimeoutAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGenericSessionTimeoutAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGenericSessionTimeoutActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
