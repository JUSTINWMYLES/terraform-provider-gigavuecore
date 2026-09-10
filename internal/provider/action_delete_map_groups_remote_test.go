package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteMapGroupsAction_Invoke_Happy exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteMapGroupsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteMapGroupsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteMapGroupsAction_Invoke_NilClient exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteMapGroupsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteMapGroupsAction{}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteMapGroupsAction_Invoke_BuildError exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteMapGroupsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteMapGroupsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteMapGroupsAction_Invoke_SendError exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteMapGroupsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteMapGroupsAction{client: newTransportErrorClient(t)}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteMapGroupsAction_Invoke_APIError exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteMapGroupsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteMapGroupsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_map_groups")
}

// TestDeleteMapGroupsAction_Invoke_APIErrorReadBody exercises DeleteMapGroupsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteMapGroupsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteMapGroupsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteMapGroupsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
