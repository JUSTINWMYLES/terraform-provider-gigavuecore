package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_Happy exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_NilClient exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_BuildError exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_SendError exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{client: newTransportErrorClient(t)}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_APIError exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_gv_tap_precryption_policy_configuration")
}

// TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_APIErrorReadBody exercises ClearGvTapPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearGvTapPrecryptionPolicyConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearGvTapPrecryptionPolicyConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearGvTapPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
