package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRebootClusterNodesAction_Invoke_Happy exercises RebootClusterNodesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRebootClusterNodesAction_Invoke_Happy(t *testing.T) {
	r := &RebootClusterNodesAction{client: newMockClientStatus(t, 200, "{}")}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRebootClusterNodesAction_Invoke_NilClient exercises RebootClusterNodesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRebootClusterNodesAction_Invoke_NilClient(t *testing.T) {
	r := &RebootClusterNodesAction{}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRebootClusterNodesAction_Invoke_BuildError exercises RebootClusterNodesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRebootClusterNodesAction_Invoke_BuildError(t *testing.T) {
	r := &RebootClusterNodesAction{client: newMalformedBaseURLClient(t)}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRebootClusterNodesAction_Invoke_SendError exercises RebootClusterNodesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRebootClusterNodesAction_Invoke_SendError(t *testing.T) {
	r := &RebootClusterNodesAction{client: newTransportErrorClient(t)}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRebootClusterNodesAction_Invoke_APIError exercises RebootClusterNodesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRebootClusterNodesAction_Invoke_APIError(t *testing.T) {
	r := &RebootClusterNodesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reboot_cluster_nodes")
}

// TestRebootClusterNodesAction_Invoke_APIErrorReadBody exercises RebootClusterNodesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRebootClusterNodesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RebootClusterNodesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RebootClusterNodesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
