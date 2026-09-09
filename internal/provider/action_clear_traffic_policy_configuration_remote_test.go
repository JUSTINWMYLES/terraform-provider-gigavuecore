package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearTrafficPolicyConfigurationAction_Invoke_Happy exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearTrafficPolicyConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearTrafficPolicyConfigurationAction_Invoke_NilClient exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearTrafficPolicyConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearTrafficPolicyConfigurationAction_Invoke_BuildError exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearTrafficPolicyConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearTrafficPolicyConfigurationAction_Invoke_SendError exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearTrafficPolicyConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{client: newTransportErrorClient(t)}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearTrafficPolicyConfigurationAction_Invoke_APIError exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearTrafficPolicyConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_traffic_policy_configuration")
}

// TestClearTrafficPolicyConfigurationAction_Invoke_APIErrorReadBody exercises ClearTrafficPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearTrafficPolicyConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearTrafficPolicyConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearTrafficPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
