package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_Happy exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_NilClient exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_BuildError exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_SendError exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{client: newTransportErrorClient(t)}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIError exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle")
}

// TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIErrorReadBody exercises RedefineSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
