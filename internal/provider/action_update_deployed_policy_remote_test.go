package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateDeployedPolicyAction_Invoke_Happy exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateDeployedPolicyAction_Invoke_Happy(t *testing.T) {
	r := &UpdateDeployedPolicyAction{client: newMockClientStatus(t, 201, "{}")}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateDeployedPolicyAction_Invoke_NilClient exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateDeployedPolicyAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateDeployedPolicyAction{}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateDeployedPolicyAction_Invoke_BuildError exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateDeployedPolicyAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateDeployedPolicyAction{client: newMalformedBaseURLClient(t)}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateDeployedPolicyAction_Invoke_SendError exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateDeployedPolicyAction_Invoke_SendError(t *testing.T) {
	r := &UpdateDeployedPolicyAction{client: newTransportErrorClient(t)}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateDeployedPolicyAction_Invoke_APIError exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateDeployedPolicyAction_Invoke_APIError(t *testing.T) {
	r := &UpdateDeployedPolicyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_deployed_policy")
}

// TestUpdateDeployedPolicyAction_Invoke_APIErrorReadBody exercises UpdateDeployedPolicyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateDeployedPolicyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateDeployedPolicyAction{client: newMockClientReadErrorBody(t, 500)}
	m := UpdateDeployedPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
