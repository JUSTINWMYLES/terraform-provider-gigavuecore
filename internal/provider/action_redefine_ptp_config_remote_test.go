package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefinePtpConfigAction_Invoke_Happy exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefinePtpConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefinePtpConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefinePtpConfigAction_Invoke_NilClient exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefinePtpConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefinePtpConfigAction{}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefinePtpConfigAction_Invoke_BuildError exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefinePtpConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefinePtpConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefinePtpConfigAction_Invoke_SendError exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefinePtpConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefinePtpConfigAction{client: newTransportErrorClient(t)}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefinePtpConfigAction_Invoke_APIError exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefinePtpConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefinePtpConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_ptp_config")
}

// TestRedefinePtpConfigAction_Invoke_APIErrorReadBody exercises RedefinePtpConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefinePtpConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefinePtpConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefinePtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
