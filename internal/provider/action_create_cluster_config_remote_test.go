package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateClusterConfigAction_Invoke_Happy exercises CreateClusterConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateClusterConfigAction_Invoke_Happy(t *testing.T) {
	r := &CreateClusterConfigAction{client: newMockClientStatus(t, 202, "{}")}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateClusterConfigAction_Invoke_NilClient exercises CreateClusterConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateClusterConfigAction_Invoke_NilClient(t *testing.T) {
	r := &CreateClusterConfigAction{}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateClusterConfigAction_Invoke_BuildError exercises CreateClusterConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateClusterConfigAction_Invoke_BuildError(t *testing.T) {
	r := &CreateClusterConfigAction{client: newMalformedBaseURLClient(t)}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateClusterConfigAction_Invoke_SendError exercises CreateClusterConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateClusterConfigAction_Invoke_SendError(t *testing.T) {
	r := &CreateClusterConfigAction{client: newTransportErrorClient(t)}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateClusterConfigAction_Invoke_APIError exercises CreateClusterConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateClusterConfigAction_Invoke_APIError(t *testing.T) {
	r := &CreateClusterConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_cluster_config")
}

// TestCreateClusterConfigAction_Invoke_APIErrorReadBody exercises CreateClusterConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateClusterConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateClusterConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateClusterConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
