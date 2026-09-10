package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeployTrafficFlowsAction_Invoke_Happy exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeployTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &DeployTrafficFlowsAction{client: newMockClientStatus(t, 202, "{}")}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeployTrafficFlowsAction_Invoke_NilClient exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeployTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &DeployTrafficFlowsAction{}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeployTrafficFlowsAction_Invoke_BuildError exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeployTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &DeployTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeployTrafficFlowsAction_Invoke_SendError exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeployTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &DeployTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeployTrafficFlowsAction_Invoke_APIError exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeployTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &DeployTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_deploy_traffic_flows")
}

// TestDeployTrafficFlowsAction_Invoke_APIErrorReadBody exercises DeployTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeployTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeployTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
