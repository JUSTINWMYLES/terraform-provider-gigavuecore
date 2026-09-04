package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClusterConfigAddMemberAction_Invoke_Happy exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClusterConfigAddMemberAction_Invoke_Happy(t *testing.T) {
	r := &ClusterConfigAddMemberAction{client: newMockClientStatus(t, 202, "{}")}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClusterConfigAddMemberAction_Invoke_NilClient exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClusterConfigAddMemberAction_Invoke_NilClient(t *testing.T) {
	r := &ClusterConfigAddMemberAction{}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClusterConfigAddMemberAction_Invoke_BuildError exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClusterConfigAddMemberAction_Invoke_BuildError(t *testing.T) {
	r := &ClusterConfigAddMemberAction{client: newMalformedBaseURLClient(t)}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClusterConfigAddMemberAction_Invoke_SendError exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClusterConfigAddMemberAction_Invoke_SendError(t *testing.T) {
	r := &ClusterConfigAddMemberAction{client: newTransportErrorClient(t)}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClusterConfigAddMemberAction_Invoke_APIError exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClusterConfigAddMemberAction_Invoke_APIError(t *testing.T) {
	r := &ClusterConfigAddMemberAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_cluster_config_add_member")
}

// TestClusterConfigAddMemberAction_Invoke_APIErrorReadBody exercises ClusterConfigAddMemberAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClusterConfigAddMemberAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClusterConfigAddMemberAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClusterConfigAddMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
