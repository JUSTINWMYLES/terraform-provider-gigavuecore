package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestModifyEmailServerAction_Invoke_Happy exercises ModifyEmailServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestModifyEmailServerAction_Invoke_Happy(t *testing.T) {
	r := &ModifyEmailServerAction{client: newMockClientStatus(t, 200, "{}")}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestModifyEmailServerAction_Invoke_NilClient exercises ModifyEmailServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestModifyEmailServerAction_Invoke_NilClient(t *testing.T) {
	r := &ModifyEmailServerAction{}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestModifyEmailServerAction_Invoke_BuildError exercises ModifyEmailServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestModifyEmailServerAction_Invoke_BuildError(t *testing.T) {
	r := &ModifyEmailServerAction{client: newMalformedBaseURLClient(t)}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestModifyEmailServerAction_Invoke_SendError exercises ModifyEmailServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestModifyEmailServerAction_Invoke_SendError(t *testing.T) {
	r := &ModifyEmailServerAction{client: newTransportErrorClient(t)}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestModifyEmailServerAction_Invoke_APIError exercises ModifyEmailServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestModifyEmailServerAction_Invoke_APIError(t *testing.T) {
	r := &ModifyEmailServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_modify_email_server")
}

// TestModifyEmailServerAction_Invoke_APIErrorReadBody exercises ModifyEmailServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestModifyEmailServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ModifyEmailServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := ModifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
