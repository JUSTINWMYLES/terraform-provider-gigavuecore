package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateTopologyLinkAction_Invoke_Happy exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateTopologyLinkAction_Invoke_Happy(t *testing.T) {
	r := &CreateTopologyLinkAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateTopologyLinkAction_Invoke_NilClient exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateTopologyLinkAction_Invoke_NilClient(t *testing.T) {
	r := &CreateTopologyLinkAction{}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateTopologyLinkAction_Invoke_BuildError exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateTopologyLinkAction_Invoke_BuildError(t *testing.T) {
	r := &CreateTopologyLinkAction{client: newMalformedBaseURLClient(t)}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateTopologyLinkAction_Invoke_SendError exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateTopologyLinkAction_Invoke_SendError(t *testing.T) {
	r := &CreateTopologyLinkAction{client: newTransportErrorClient(t)}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateTopologyLinkAction_Invoke_APIError exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateTopologyLinkAction_Invoke_APIError(t *testing.T) {
	r := &CreateTopologyLinkAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_topology_link")
}

// TestCreateTopologyLinkAction_Invoke_APIErrorReadBody exercises CreateTopologyLinkAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateTopologyLinkAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateTopologyLinkAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateTopologyLinkActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
