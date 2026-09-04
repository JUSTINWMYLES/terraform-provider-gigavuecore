package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllInlineSslProfilesAction_Invoke_Happy exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllInlineSslProfilesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllInlineSslProfilesAction_Invoke_NilClient exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllInlineSslProfilesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllInlineSslProfilesAction_Invoke_BuildError exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllInlineSslProfilesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllInlineSslProfilesAction_Invoke_SendError exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllInlineSslProfilesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{client: newTransportErrorClient(t)}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllInlineSslProfilesAction_Invoke_APIError exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllInlineSslProfilesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_inline_ssl_profiles")
}

// TestDeleteAllInlineSslProfilesAction_Invoke_APIErrorReadBody exercises DeleteAllInlineSslProfilesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllInlineSslProfilesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllInlineSslProfilesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllInlineSslProfilesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
