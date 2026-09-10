package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteClusterPortFilterCountersAction_Invoke_Happy exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteClusterPortFilterCountersAction_Invoke_Happy(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteClusterPortFilterCountersAction_Invoke_NilClient exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteClusterPortFilterCountersAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteClusterPortFilterCountersAction_Invoke_BuildError exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteClusterPortFilterCountersAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{client: newMalformedBaseURLClient(t)}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteClusterPortFilterCountersAction_Invoke_SendError exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteClusterPortFilterCountersAction_Invoke_SendError(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{client: newTransportErrorClient(t)}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteClusterPortFilterCountersAction_Invoke_APIError exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteClusterPortFilterCountersAction_Invoke_APIError(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_cluster_port_filter_counters")
}

// TestDeleteClusterPortFilterCountersAction_Invoke_APIErrorReadBody exercises DeleteClusterPortFilterCountersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteClusterPortFilterCountersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteClusterPortFilterCountersAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteClusterPortFilterCountersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
