package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClusterConfigDeleteMemberAction_Invoke_Happy exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClusterConfigDeleteMemberAction_Invoke_Happy(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{client: newMockClientStatus(t, 202, "{}")}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClusterConfigDeleteMemberAction_Invoke_NilClient exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClusterConfigDeleteMemberAction_Invoke_NilClient(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClusterConfigDeleteMemberAction_Invoke_BuildError exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClusterConfigDeleteMemberAction_Invoke_BuildError(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{client: newMalformedBaseURLClient(t)}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClusterConfigDeleteMemberAction_Invoke_SendError exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClusterConfigDeleteMemberAction_Invoke_SendError(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{client: newTransportErrorClient(t)}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClusterConfigDeleteMemberAction_Invoke_APIError exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClusterConfigDeleteMemberAction_Invoke_APIError(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_cluster_config_delete_member")
}

// TestClusterConfigDeleteMemberAction_Invoke_APIErrorReadBody exercises ClusterConfigDeleteMemberAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClusterConfigDeleteMemberAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClusterConfigDeleteMemberAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClusterConfigDeleteMemberActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
