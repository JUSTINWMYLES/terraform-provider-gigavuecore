package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddNtpAuthKeyAction_Invoke_Happy exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddNtpAuthKeyAction_Invoke_Happy(t *testing.T) {
	r := &AddNtpAuthKeyAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddNtpAuthKeyAction_Invoke_NilClient exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddNtpAuthKeyAction_Invoke_NilClient(t *testing.T) {
	r := &AddNtpAuthKeyAction{}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddNtpAuthKeyAction_Invoke_BuildError exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddNtpAuthKeyAction_Invoke_BuildError(t *testing.T) {
	r := &AddNtpAuthKeyAction{client: newMalformedBaseURLClient(t)}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddNtpAuthKeyAction_Invoke_SendError exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddNtpAuthKeyAction_Invoke_SendError(t *testing.T) {
	r := &AddNtpAuthKeyAction{client: newTransportErrorClient(t)}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddNtpAuthKeyAction_Invoke_APIError exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddNtpAuthKeyAction_Invoke_APIError(t *testing.T) {
	r := &AddNtpAuthKeyAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_ntp_auth_key")
}

// TestAddNtpAuthKeyAction_Invoke_APIErrorReadBody exercises AddNtpAuthKeyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddNtpAuthKeyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddNtpAuthKeyAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddNtpAuthKeyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
