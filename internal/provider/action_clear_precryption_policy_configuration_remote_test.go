package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearPrecryptionPolicyConfigurationAction_Invoke_Happy exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_Happy(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearPrecryptionPolicyConfigurationAction_Invoke_NilClient exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_NilClient(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearPrecryptionPolicyConfigurationAction_Invoke_BuildError exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_BuildError(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{client: newMalformedBaseURLClient(t)}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearPrecryptionPolicyConfigurationAction_Invoke_SendError exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_SendError(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{client: newTransportErrorClient(t)}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearPrecryptionPolicyConfigurationAction_Invoke_APIError exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_APIError(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_precryption_policy_configuration")
}

// TestClearPrecryptionPolicyConfigurationAction_Invoke_APIErrorReadBody exercises ClearPrecryptionPolicyConfigurationAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearPrecryptionPolicyConfigurationAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearPrecryptionPolicyConfigurationAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearPrecryptionPolicyConfigurationActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
