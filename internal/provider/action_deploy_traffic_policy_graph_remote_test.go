package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeployTrafficPolicyGraphAction_Invoke_Happy exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeployTrafficPolicyGraphAction_Invoke_Happy(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{client: newMockClientStatus(t, 201, "{}")}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeployTrafficPolicyGraphAction_Invoke_NilClient exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeployTrafficPolicyGraphAction_Invoke_NilClient(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeployTrafficPolicyGraphAction_Invoke_BuildError exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeployTrafficPolicyGraphAction_Invoke_BuildError(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{client: newMalformedBaseURLClient(t)}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeployTrafficPolicyGraphAction_Invoke_SendError exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeployTrafficPolicyGraphAction_Invoke_SendError(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{client: newTransportErrorClient(t)}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeployTrafficPolicyGraphAction_Invoke_APIError exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeployTrafficPolicyGraphAction_Invoke_APIError(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_deploy_traffic_policy_graph")
}

// TestDeployTrafficPolicyGraphAction_Invoke_APIErrorReadBody exercises DeployTrafficPolicyGraphAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeployTrafficPolicyGraphAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeployTrafficPolicyGraphAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeployTrafficPolicyGraphActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
