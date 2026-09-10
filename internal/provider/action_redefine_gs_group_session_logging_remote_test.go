package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSessionLoggingAction_Invoke_Happy exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSessionLoggingAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSessionLoggingAction_Invoke_NilClient exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSessionLoggingAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSessionLoggingAction_Invoke_BuildError exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSessionLoggingAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSessionLoggingAction_Invoke_SendError exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSessionLoggingAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSessionLoggingAction_Invoke_APIError exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSessionLoggingAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_session_logging")
}

// TestRedefineGsGroupSessionLoggingAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSessionLoggingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSessionLoggingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSessionLoggingAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSessionLoggingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
