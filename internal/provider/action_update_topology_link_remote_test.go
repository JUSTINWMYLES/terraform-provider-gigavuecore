package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateTopologyLinkAction_Invoke_Happy exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateTopologyLinkAction_Invoke_Happy(t *testing.T) {
	r := &UpdateTopologyLinkAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateTopologyLinkAction_Invoke_NilClient exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateTopologyLinkAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateTopologyLinkAction{}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateTopologyLinkAction_Invoke_BuildError exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateTopologyLinkAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateTopologyLinkAction{client: newMalformedBaseURLClient(t)}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateTopologyLinkAction_Invoke_SendError exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateTopologyLinkAction_Invoke_SendError(t *testing.T) {
	r := &UpdateTopologyLinkAction{client: newTransportErrorClient(t)}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateTopologyLinkAction_Invoke_APIError exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateTopologyLinkAction_Invoke_APIError(t *testing.T) {
	r := &UpdateTopologyLinkAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_topology_link")
}

// TestUpdateTopologyLinkAction_Invoke_APIErrorReadBody exercises UpdateTopologyLinkAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateTopologyLinkAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateTopologyLinkAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
