package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestApplyPrecryptionPolicyConfigAction_Invoke_Happy exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestApplyPrecryptionPolicyConfigAction_Invoke_Happy(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplyPrecryptionPolicyConfigAction_Invoke_NilClient exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplyPrecryptionPolicyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplyPrecryptionPolicyConfigAction_Invoke_BuildError exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplyPrecryptionPolicyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{client: newMalformedBaseURLClient(t)}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplyPrecryptionPolicyConfigAction_Invoke_SendError exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplyPrecryptionPolicyConfigAction_Invoke_SendError(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{client: newTransportErrorClient(t)}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplyPrecryptionPolicyConfigAction_Invoke_APIError exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplyPrecryptionPolicyConfigAction_Invoke_APIError(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_apply_precryption_policy_config")
}

// TestApplyPrecryptionPolicyConfigAction_Invoke_APIErrorReadBody exercises ApplyPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplyPrecryptionPolicyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ApplyPrecryptionPolicyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ApplyPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
