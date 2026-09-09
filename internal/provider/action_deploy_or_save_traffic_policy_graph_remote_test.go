package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_Happy exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_Happy(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{client: newMockClientStatus(t, 201, "{}")}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_NilClient exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_NilClient(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_BuildError exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_BuildError(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{client: newMalformedBaseURLClient(t)}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_SendError exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_SendError(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{client: newTransportErrorClient(t)}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_APIError exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_APIError(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_deploy_or_save_traffic_policy_graph")
}

// TestDeployOrSaveTrafficPolicyGraphAction_Invoke_APIErrorReadBody exercises DeployOrSaveTrafficPolicyGraphAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeployOrSaveTrafficPolicyGraphAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeployOrSaveTrafficPolicyGraphAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeployOrSaveTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
