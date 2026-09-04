package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRenewBindingsAction_Invoke_Happy exercises RenewBindingsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRenewBindingsAction_Invoke_Happy(t *testing.T) {
	r := &RenewBindingsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRenewBindingsAction_Invoke_NilClient exercises RenewBindingsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRenewBindingsAction_Invoke_NilClient(t *testing.T) {
	r := &RenewBindingsAction{}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRenewBindingsAction_Invoke_BuildError exercises RenewBindingsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRenewBindingsAction_Invoke_BuildError(t *testing.T) {
	r := &RenewBindingsAction{client: newMalformedBaseURLClient(t)}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRenewBindingsAction_Invoke_SendError exercises RenewBindingsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRenewBindingsAction_Invoke_SendError(t *testing.T) {
	r := &RenewBindingsAction{client: newTransportErrorClient(t)}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRenewBindingsAction_Invoke_APIError exercises RenewBindingsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRenewBindingsAction_Invoke_APIError(t *testing.T) {
	r := &RenewBindingsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_renew_bindings")
}

// TestRenewBindingsAction_Invoke_APIErrorReadBody exercises RenewBindingsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRenewBindingsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RenewBindingsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RenewBindingsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
