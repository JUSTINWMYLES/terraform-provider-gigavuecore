package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_Happy exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_Happy(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_NilClient exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_BuildError exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{client: newMalformedBaseURLClient(t)}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_SendError exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_SendError(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{client: newTransportErrorClient(t)}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_APIError exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_APIError(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_apply_gv_tap_precryption_policy_config")
}

// TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_APIErrorReadBody exercises ApplyGvTapPrecryptionPolicyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplyGvTapPrecryptionPolicyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ApplyGvTapPrecryptionPolicyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ApplyGvTapPrecryptionPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
