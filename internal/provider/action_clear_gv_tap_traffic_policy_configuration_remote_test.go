package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_Happy exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_NilClient exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_BuildError exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_SendError exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{client: newTransportErrorClient(t)}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_APIError exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_gv_tap_traffic_policy_configuration")
}

// TestClearGvTapTrafficPolicyConfigurationAction_Invoke_APIErrorReadBody exercises ClearGvTapTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearGvTapTrafficPolicyConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearGvTapTrafficPolicyConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearGvTapTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
