package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_Happy exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_NilClient exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_BuildError exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_SendError exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_APIError exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_engine_watchdog_timer")
}

// TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_APIErrorReadBody exercises RedefineGsGroupEngineWatchdogTimerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupEngineWatchdogTimerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupEngineWatchdogTimerAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupEngineWatchdogTimerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
