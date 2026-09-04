package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateGtpWhitelistEntriesAction_Invoke_Happy exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateGtpWhitelistEntriesAction_Invoke_Happy(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateGtpWhitelistEntriesAction_Invoke_NilClient exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateGtpWhitelistEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateGtpWhitelistEntriesAction_Invoke_BuildError exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateGtpWhitelistEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{client: newMalformedBaseURLClient(t)}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateGtpWhitelistEntriesAction_Invoke_SendError exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateGtpWhitelistEntriesAction_Invoke_SendError(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{client: newTransportErrorClient(t)}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateGtpWhitelistEntriesAction_Invoke_APIError exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateGtpWhitelistEntriesAction_Invoke_APIError(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_gtp_whitelist_entries")
}

// TestCreateGtpWhitelistEntriesAction_Invoke_APIErrorReadBody exercises CreateGtpWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateGtpWhitelistEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateGtpWhitelistEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateGtpWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
