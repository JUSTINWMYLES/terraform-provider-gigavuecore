package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineSnmpThrottleConfigAction_Invoke_Happy exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineSnmpThrottleConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineSnmpThrottleConfigAction_Invoke_NilClient exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineSnmpThrottleConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineSnmpThrottleConfigAction_Invoke_BuildError exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineSnmpThrottleConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineSnmpThrottleConfigAction_Invoke_SendError exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineSnmpThrottleConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{client: newTransportErrorClient(t)}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineSnmpThrottleConfigAction_Invoke_APIError exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineSnmpThrottleConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_snmp_throttle_config")
}

// TestRedefineSnmpThrottleConfigAction_Invoke_APIErrorReadBody exercises RedefineSnmpThrottleConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineSnmpThrottleConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineSnmpThrottleConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
