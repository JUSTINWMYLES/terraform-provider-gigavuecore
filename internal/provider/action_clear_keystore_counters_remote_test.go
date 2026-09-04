package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearKeystoreCountersAction_Invoke_Happy exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearKeystoreCountersAction_Invoke_Happy(t *testing.T) {
	r := &ClearKeystoreCountersAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearKeystoreCountersAction_Invoke_NilClient exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearKeystoreCountersAction_Invoke_NilClient(t *testing.T) {
	r := &ClearKeystoreCountersAction{}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearKeystoreCountersAction_Invoke_BuildError exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearKeystoreCountersAction_Invoke_BuildError(t *testing.T) {
	r := &ClearKeystoreCountersAction{client: newMalformedBaseURLClient(t)}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearKeystoreCountersAction_Invoke_SendError exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearKeystoreCountersAction_Invoke_SendError(t *testing.T) {
	r := &ClearKeystoreCountersAction{client: newTransportErrorClient(t)}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearKeystoreCountersAction_Invoke_APIError exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearKeystoreCountersAction_Invoke_APIError(t *testing.T) {
	r := &ClearKeystoreCountersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_keystore_counters")
}

// TestClearKeystoreCountersAction_Invoke_APIErrorReadBody exercises ClearKeystoreCountersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearKeystoreCountersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearKeystoreCountersAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearKeystoreCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
