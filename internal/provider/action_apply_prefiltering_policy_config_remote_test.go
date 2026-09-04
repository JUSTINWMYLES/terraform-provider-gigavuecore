package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestApplyPrefilteringPolicyConfigAction_Invoke_Happy exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestApplyPrefilteringPolicyConfigAction_Invoke_Happy(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestApplyPrefilteringPolicyConfigAction_Invoke_NilClient exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestApplyPrefilteringPolicyConfigAction_Invoke_NilClient(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestApplyPrefilteringPolicyConfigAction_Invoke_BuildError exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestApplyPrefilteringPolicyConfigAction_Invoke_BuildError(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{client: newMalformedBaseURLClient(t)}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestApplyPrefilteringPolicyConfigAction_Invoke_SendError exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestApplyPrefilteringPolicyConfigAction_Invoke_SendError(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{client: newTransportErrorClient(t)}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestApplyPrefilteringPolicyConfigAction_Invoke_APIError exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestApplyPrefilteringPolicyConfigAction_Invoke_APIError(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_apply_prefiltering_policy_config")
}

// TestApplyPrefilteringPolicyConfigAction_Invoke_APIErrorReadBody exercises ApplyPrefilteringPolicyConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestApplyPrefilteringPolicyConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ApplyPrefilteringPolicyConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := ApplyPrefilteringPolicyConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
