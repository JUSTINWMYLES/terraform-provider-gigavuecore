package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteTopologyLinkAction_Invoke_Happy exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteTopologyLinkAction_Invoke_Happy(t *testing.T) {
	r := &DeleteTopologyLinkAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteTopologyLinkAction_Invoke_NilClient exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteTopologyLinkAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteTopologyLinkAction{}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteTopologyLinkAction_Invoke_BuildError exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteTopologyLinkAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteTopologyLinkAction{client: newMalformedBaseURLClient(t)}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteTopologyLinkAction_Invoke_SendError exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteTopologyLinkAction_Invoke_SendError(t *testing.T) {
	r := &DeleteTopologyLinkAction{client: newTransportErrorClient(t)}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteTopologyLinkAction_Invoke_APIError exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteTopologyLinkAction_Invoke_APIError(t *testing.T) {
	r := &DeleteTopologyLinkAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_topology_link")
}

// TestDeleteTopologyLinkAction_Invoke_APIErrorReadBody exercises DeleteTopologyLinkAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteTopologyLinkAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteTopologyLinkAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
