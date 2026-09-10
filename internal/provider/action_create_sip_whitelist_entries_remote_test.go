package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateSipWhitelistEntriesAction_Invoke_Happy exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateSipWhitelistEntriesAction_Invoke_Happy(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateSipWhitelistEntriesAction_Invoke_NilClient exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateSipWhitelistEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateSipWhitelistEntriesAction_Invoke_BuildError exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateSipWhitelistEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{client: newMalformedBaseURLClient(t)}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateSipWhitelistEntriesAction_Invoke_SendError exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateSipWhitelistEntriesAction_Invoke_SendError(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{client: newTransportErrorClient(t)}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateSipWhitelistEntriesAction_Invoke_APIError exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateSipWhitelistEntriesAction_Invoke_APIError(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_sip_whitelist_entries")
}

// TestCreateSipWhitelistEntriesAction_Invoke_APIErrorReadBody exercises CreateSipWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateSipWhitelistEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateSipWhitelistEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateSipWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
