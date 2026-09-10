package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineNtpConfigAction_Invoke_Happy exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineNtpConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineNtpConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineNtpConfigAction_Invoke_NilClient exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineNtpConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineNtpConfigAction{}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineNtpConfigAction_Invoke_BuildError exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineNtpConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineNtpConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineNtpConfigAction_Invoke_SendError exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineNtpConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineNtpConfigAction{client: newTransportErrorClient(t)}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineNtpConfigAction_Invoke_APIError exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineNtpConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineNtpConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_ntp_config")
}

// TestRedefineNtpConfigAction_Invoke_APIErrorReadBody exercises RedefineNtpConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineNtpConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineNtpConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineNtpConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
