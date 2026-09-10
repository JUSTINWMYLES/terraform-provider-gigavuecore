package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_Happy exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_NilClient exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_BuildError exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_SendError exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{client: newTransportErrorClient(t)}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIError exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_snmp_throttle_config_system_snmp_throttle")
}

// TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIErrorReadBody exercises UpdateSnmpThrottleConfigSystemSnmpThrottleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSnmpThrottleConfigSystemSnmpThrottleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSnmpThrottleConfigSystemSnmpThrottleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
