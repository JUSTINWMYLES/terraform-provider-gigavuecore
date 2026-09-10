package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearPtpPortsCountersByPortIdAction_Invoke_Happy exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearPtpPortsCountersByPortIdAction_Invoke_Happy(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearPtpPortsCountersByPortIdAction_Invoke_NilClient exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearPtpPortsCountersByPortIdAction_Invoke_NilClient(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearPtpPortsCountersByPortIdAction_Invoke_BuildError exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearPtpPortsCountersByPortIdAction_Invoke_BuildError(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{client: newMalformedBaseURLClient(t)}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearPtpPortsCountersByPortIdAction_Invoke_SendError exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearPtpPortsCountersByPortIdAction_Invoke_SendError(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{client: newTransportErrorClient(t)}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearPtpPortsCountersByPortIdAction_Invoke_APIError exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearPtpPortsCountersByPortIdAction_Invoke_APIError(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_ptp_ports_counters_by_port_id")
}

// TestClearPtpPortsCountersByPortIdAction_Invoke_APIErrorReadBody exercises ClearPtpPortsCountersByPortIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearPtpPortsCountersByPortIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearPtpPortsCountersByPortIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearPtpPortsCountersByPortIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
