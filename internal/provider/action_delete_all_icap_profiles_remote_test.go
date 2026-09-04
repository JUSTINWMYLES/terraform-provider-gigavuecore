package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllIcapProfilesAction_Invoke_Happy exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllIcapProfilesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllIcapProfilesAction_Invoke_NilClient exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllIcapProfilesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllIcapProfilesAction_Invoke_BuildError exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllIcapProfilesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllIcapProfilesAction_Invoke_SendError exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllIcapProfilesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{client: newTransportErrorClient(t)}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllIcapProfilesAction_Invoke_APIError exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllIcapProfilesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_icap_profiles")
}

// TestDeleteAllIcapProfilesAction_Invoke_APIErrorReadBody exercises DeleteAllIcapProfilesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllIcapProfilesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllIcapProfilesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllIcapProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
