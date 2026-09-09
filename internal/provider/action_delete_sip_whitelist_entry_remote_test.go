package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSipWhitelistEntryAction_Invoke_Happy exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSipWhitelistEntryAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSipWhitelistEntryAction_Invoke_NilClient exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSipWhitelistEntryAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSipWhitelistEntryAction_Invoke_BuildError exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSipWhitelistEntryAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSipWhitelistEntryAction_Invoke_SendError exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSipWhitelistEntryAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{client: newTransportErrorClient(t)}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSipWhitelistEntryAction_Invoke_APIError exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSipWhitelistEntryAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_sip_whitelist_entry")
}

// TestDeleteSipWhitelistEntryAction_Invoke_APIErrorReadBody exercises DeleteSipWhitelistEntryAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSipWhitelistEntryAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSipWhitelistEntryAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSipWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
