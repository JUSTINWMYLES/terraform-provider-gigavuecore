package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestValidateUpdatedPolicyAction_Invoke_Happy exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestValidateUpdatedPolicyAction_Invoke_Happy(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{client: newMockClientStatus(t, 201, "{}")}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestValidateUpdatedPolicyAction_Invoke_NilClient exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestValidateUpdatedPolicyAction_Invoke_NilClient(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestValidateUpdatedPolicyAction_Invoke_BuildError exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestValidateUpdatedPolicyAction_Invoke_BuildError(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{client: newMalformedBaseURLClient(t)}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestValidateUpdatedPolicyAction_Invoke_SendError exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestValidateUpdatedPolicyAction_Invoke_SendError(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{client: newTransportErrorClient(t)}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestValidateUpdatedPolicyAction_Invoke_APIError exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestValidateUpdatedPolicyAction_Invoke_APIError(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_validate_updated_policy")
}

// TestValidateUpdatedPolicyAction_Invoke_APIErrorReadBody exercises ValidateUpdatedPolicyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestValidateUpdatedPolicyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ValidateUpdatedPolicyAction{client: newMockClientReadErrorBody(t, 500)}
	m := ValidateUpdatedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
