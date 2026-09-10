package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_Happy exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_Happy(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_NilClient exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_BuildError exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{client: newMalformedBaseURLClient(t)}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_SendError exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_SendError(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{client: newTransportErrorClient(t)}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_APIError exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_APIError(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_apply_gv_tap_prefiltering_policy_config")
}

// TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_APIErrorReadBody exercises ApplyGvTapPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplyGvTapPrefilteringPolicyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ApplyGvTapPrefilteringPolicyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ApplyGvTapPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
