package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUndeployTrafficFlowsAction_Invoke_Happy exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUndeployTrafficFlowsAction_Invoke_Happy(t *testing.T) {
	r := &UndeployTrafficFlowsAction{client: newMockClientStatus(t, 204, "{}")}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUndeployTrafficFlowsAction_Invoke_NilClient exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUndeployTrafficFlowsAction_Invoke_NilClient(t *testing.T) {
	r := &UndeployTrafficFlowsAction{}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUndeployTrafficFlowsAction_Invoke_BuildError exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUndeployTrafficFlowsAction_Invoke_BuildError(t *testing.T) {
	r := &UndeployTrafficFlowsAction{client: newMalformedBaseURLClient(t)}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUndeployTrafficFlowsAction_Invoke_SendError exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUndeployTrafficFlowsAction_Invoke_SendError(t *testing.T) {
	r := &UndeployTrafficFlowsAction{client: newTransportErrorClient(t)}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUndeployTrafficFlowsAction_Invoke_APIError exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUndeployTrafficFlowsAction_Invoke_APIError(t *testing.T) {
	r := &UndeployTrafficFlowsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_undeploy_traffic_flows")
}

// TestUndeployTrafficFlowsAction_Invoke_APIErrorReadBody exercises UndeployTrafficFlowsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUndeployTrafficFlowsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UndeployTrafficFlowsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UndeployTrafficFlowsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
