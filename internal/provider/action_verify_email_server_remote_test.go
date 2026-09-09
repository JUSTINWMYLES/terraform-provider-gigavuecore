package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestVerifyEmailServerAction_Invoke_Happy exercises VerifyEmailServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestVerifyEmailServerAction_Invoke_Happy(t *testing.T) {
	r := &VerifyEmailServerAction{client: newMockClientStatus(t, 200, "{}")}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestVerifyEmailServerAction_Invoke_NilClient exercises VerifyEmailServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestVerifyEmailServerAction_Invoke_NilClient(t *testing.T) {
	r := &VerifyEmailServerAction{}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestVerifyEmailServerAction_Invoke_BuildError exercises VerifyEmailServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestVerifyEmailServerAction_Invoke_BuildError(t *testing.T) {
	r := &VerifyEmailServerAction{client: newMalformedBaseURLClient(t)}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestVerifyEmailServerAction_Invoke_SendError exercises VerifyEmailServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestVerifyEmailServerAction_Invoke_SendError(t *testing.T) {
	r := &VerifyEmailServerAction{client: newTransportErrorClient(t)}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestVerifyEmailServerAction_Invoke_APIError exercises VerifyEmailServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestVerifyEmailServerAction_Invoke_APIError(t *testing.T) {
	r := &VerifyEmailServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_verify_email_server")
}

// TestVerifyEmailServerAction_Invoke_APIErrorReadBody exercises VerifyEmailServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestVerifyEmailServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &VerifyEmailServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := VerifyEmailServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
