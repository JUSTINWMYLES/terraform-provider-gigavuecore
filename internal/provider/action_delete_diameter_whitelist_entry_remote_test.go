package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteDiameterWhitelistEntryAction_Invoke_Happy exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteDiameterWhitelistEntryAction_Invoke_Happy(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteDiameterWhitelistEntryAction_Invoke_NilClient exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteDiameterWhitelistEntryAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteDiameterWhitelistEntryAction_Invoke_BuildError exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteDiameterWhitelistEntryAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{client: newMalformedBaseURLClient(t)}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteDiameterWhitelistEntryAction_Invoke_SendError exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteDiameterWhitelistEntryAction_Invoke_SendError(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{client: newTransportErrorClient(t)}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteDiameterWhitelistEntryAction_Invoke_APIError exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteDiameterWhitelistEntryAction_Invoke_APIError(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_diameter_whitelist_entry")
}

// TestDeleteDiameterWhitelistEntryAction_Invoke_APIErrorReadBody exercises DeleteDiameterWhitelistEntryAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteDiameterWhitelistEntryAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteDiameterWhitelistEntryAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteDiameterWhitelistEntryActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
