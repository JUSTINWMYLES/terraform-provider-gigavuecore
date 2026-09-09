package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClusterConfigEditLeaderPreferenceAction_Invoke_Happy exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_Happy(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{client: newMockClientStatus(t, 202, "{}")}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClusterConfigEditLeaderPreferenceAction_Invoke_NilClient exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_NilClient(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClusterConfigEditLeaderPreferenceAction_Invoke_BuildError exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_BuildError(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{client: newMalformedBaseURLClient(t)}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClusterConfigEditLeaderPreferenceAction_Invoke_SendError exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_SendError(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{client: newTransportErrorClient(t)}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClusterConfigEditLeaderPreferenceAction_Invoke_APIError exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_APIError(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_cluster_config_edit_leader_preference")
}

// TestClusterConfigEditLeaderPreferenceAction_Invoke_APIErrorReadBody exercises ClusterConfigEditLeaderPreferenceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClusterConfigEditLeaderPreferenceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClusterConfigEditLeaderPreferenceAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClusterConfigEditLeaderPreferenceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
