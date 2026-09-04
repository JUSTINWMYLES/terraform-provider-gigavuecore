package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateDiameterWhitelistEntriesAction_Invoke_Happy exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateDiameterWhitelistEntriesAction_Invoke_Happy(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateDiameterWhitelistEntriesAction_Invoke_NilClient exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateDiameterWhitelistEntriesAction_Invoke_NilClient(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateDiameterWhitelistEntriesAction_Invoke_BuildError exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateDiameterWhitelistEntriesAction_Invoke_BuildError(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{client: newMalformedBaseURLClient(t)}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateDiameterWhitelistEntriesAction_Invoke_SendError exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateDiameterWhitelistEntriesAction_Invoke_SendError(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{client: newTransportErrorClient(t)}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateDiameterWhitelistEntriesAction_Invoke_APIError exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateDiameterWhitelistEntriesAction_Invoke_APIError(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_diameter_whitelist_entries")
}

// TestCreateDiameterWhitelistEntriesAction_Invoke_APIErrorReadBody exercises CreateDiameterWhitelistEntriesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateDiameterWhitelistEntriesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateDiameterWhitelistEntriesAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateDiameterWhitelistEntriesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
