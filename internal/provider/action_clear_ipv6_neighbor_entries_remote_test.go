package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestClearIpv6NeighborEntriesAction_Invoke_Happy exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestClearIpv6NeighborEntriesAction_Invoke_Happy(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{client: newMockClientStatus(t, 204, "{}")}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestClearIpv6NeighborEntriesAction_Invoke_NilClient exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestClearIpv6NeighborEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestClearIpv6NeighborEntriesAction_Invoke_BuildError exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestClearIpv6NeighborEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{client: newMalformedBaseURLClient(t)}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestClearIpv6NeighborEntriesAction_Invoke_SendError exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestClearIpv6NeighborEntriesAction_Invoke_SendError(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{client: newTransportErrorClient(t)}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestClearIpv6NeighborEntriesAction_Invoke_APIError exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestClearIpv6NeighborEntriesAction_Invoke_APIError(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_clear_ipv6_neighbor_entries")
}

// TestClearIpv6NeighborEntriesAction_Invoke_APIErrorReadBody exercises ClearIpv6NeighborEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestClearIpv6NeighborEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ClearIpv6NeighborEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := ClearIpv6NeighborEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
