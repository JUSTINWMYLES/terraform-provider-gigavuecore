package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_Happy exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_Happy(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{client: newMockClientStatus(t, 200, "{}")}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_NilClient exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_NilClient(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_BuildError exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_BuildError(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{client: newMalformedBaseURLClient(t)}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_SendError exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_SendError(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{client: newTransportErrorClient(t)}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_APIError exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_APIError(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_cluster_config_switch_i_pprotocol_spec")
}

// TestClusterConfigSwitchIPprotocolSpecAction_Invoke_APIErrorReadBody exercises ClusterConfigSwitchIPprotocolSpecAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClusterConfigSwitchIPprotocolSpecAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClusterConfigSwitchIPprotocolSpecAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClusterConfigSwitchIPprotocolSpecActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
