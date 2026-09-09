package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllGtaProfilesAction_Invoke_Happy exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllGtaProfilesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllGtaProfilesAction_Invoke_NilClient exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllGtaProfilesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllGtaProfilesAction_Invoke_BuildError exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllGtaProfilesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllGtaProfilesAction_Invoke_SendError exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllGtaProfilesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{client: newTransportErrorClient(t)}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllGtaProfilesAction_Invoke_APIError exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllGtaProfilesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_gta_profiles")
}

// TestDeleteAllGtaProfilesAction_Invoke_APIErrorReadBody exercises DeleteAllGtaProfilesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllGtaProfilesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllGtaProfilesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllGtaProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
