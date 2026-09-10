package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineFmNtpServerAction_Invoke_Happy exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineFmNtpServerAction_Invoke_Happy(t *testing.T) {
	r := &RedefineFmNtpServerAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineFmNtpServerAction_Invoke_NilClient exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineFmNtpServerAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineFmNtpServerAction{}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineFmNtpServerAction_Invoke_BuildError exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineFmNtpServerAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineFmNtpServerAction{client: newMalformedBaseURLClient(t)}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineFmNtpServerAction_Invoke_SendError exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineFmNtpServerAction_Invoke_SendError(t *testing.T) {
	r := &RedefineFmNtpServerAction{client: newTransportErrorClient(t)}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineFmNtpServerAction_Invoke_APIError exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineFmNtpServerAction_Invoke_APIError(t *testing.T) {
	r := &RedefineFmNtpServerAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_fm_ntp_server")
}

// TestRedefineFmNtpServerAction_Invoke_APIErrorReadBody exercises RedefineFmNtpServerAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineFmNtpServerAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineFmNtpServerAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineFmNtpServerActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
