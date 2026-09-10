package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUndeployPolicyAction_Invoke_Happy exercises UndeployPolicyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUndeployPolicyAction_Invoke_Happy(t *testing.T) {
	r := &UndeployPolicyAction{client: newMockClientStatus(t, 204, "{}")}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUndeployPolicyAction_Invoke_NilClient exercises UndeployPolicyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUndeployPolicyAction_Invoke_NilClient(t *testing.T) {
	r := &UndeployPolicyAction{}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUndeployPolicyAction_Invoke_BuildError exercises UndeployPolicyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUndeployPolicyAction_Invoke_BuildError(t *testing.T) {
	r := &UndeployPolicyAction{client: newMalformedBaseURLClient(t)}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUndeployPolicyAction_Invoke_SendError exercises UndeployPolicyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUndeployPolicyAction_Invoke_SendError(t *testing.T) {
	r := &UndeployPolicyAction{client: newTransportErrorClient(t)}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUndeployPolicyAction_Invoke_APIError exercises UndeployPolicyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUndeployPolicyAction_Invoke_APIError(t *testing.T) {
	r := &UndeployPolicyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_undeploy_policy")
}

// TestUndeployPolicyAction_Invoke_APIErrorReadBody exercises UndeployPolicyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUndeployPolicyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UndeployPolicyAction{client: newMockClientReadErrorBody(t, 500)}
	m := UndeployPolicyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
