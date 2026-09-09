package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSetUserSelectedCiphersAction_Invoke_Happy exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSetUserSelectedCiphersAction_Invoke_Happy(t *testing.T) {
	r := &SetUserSelectedCiphersAction{client: newMockClientStatus(t, 200, "{}")}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSetUserSelectedCiphersAction_Invoke_NilClient exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSetUserSelectedCiphersAction_Invoke_NilClient(t *testing.T) {
	r := &SetUserSelectedCiphersAction{}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSetUserSelectedCiphersAction_Invoke_BuildError exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSetUserSelectedCiphersAction_Invoke_BuildError(t *testing.T) {
	r := &SetUserSelectedCiphersAction{client: newMalformedBaseURLClient(t)}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSetUserSelectedCiphersAction_Invoke_SendError exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSetUserSelectedCiphersAction_Invoke_SendError(t *testing.T) {
	r := &SetUserSelectedCiphersAction{client: newTransportErrorClient(t)}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSetUserSelectedCiphersAction_Invoke_APIError exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSetUserSelectedCiphersAction_Invoke_APIError(t *testing.T) {
	r := &SetUserSelectedCiphersAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_set_user_selected_ciphers")
}

// TestSetUserSelectedCiphersAction_Invoke_APIErrorReadBody exercises SetUserSelectedCiphersAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSetUserSelectedCiphersAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SetUserSelectedCiphersAction{client: newMockClientReadErrorBody(t, 501)}
	m := SetUserSelectedCiphersActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
