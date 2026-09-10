package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateSnmpThrottleConfigAction_Invoke_Happy exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateSnmpThrottleConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateSnmpThrottleConfigAction_Invoke_NilClient exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateSnmpThrottleConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateSnmpThrottleConfigAction_Invoke_BuildError exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateSnmpThrottleConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateSnmpThrottleConfigAction_Invoke_SendError exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateSnmpThrottleConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{client: newTransportErrorClient(t)}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateSnmpThrottleConfigAction_Invoke_APIError exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateSnmpThrottleConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_snmp_throttle_config")
}

// TestUpdateSnmpThrottleConfigAction_Invoke_APIErrorReadBody exercises UpdateSnmpThrottleConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateSnmpThrottleConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateSnmpThrottleConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateSnmpThrottleConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
