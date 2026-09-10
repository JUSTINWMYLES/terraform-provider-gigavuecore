package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRebootFmAction_Invoke_Happy exercises RebootFmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRebootFmAction_Invoke_Happy(t *testing.T) {
	r := &RebootFmAction{client: newMockClientStatus(t, 200, "{}")}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRebootFmAction_Invoke_NilClient exercises RebootFmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRebootFmAction_Invoke_NilClient(t *testing.T) {
	r := &RebootFmAction{}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRebootFmAction_Invoke_BuildError exercises RebootFmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRebootFmAction_Invoke_BuildError(t *testing.T) {
	r := &RebootFmAction{client: newMalformedBaseURLClient(t)}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRebootFmAction_Invoke_SendError exercises RebootFmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRebootFmAction_Invoke_SendError(t *testing.T) {
	r := &RebootFmAction{client: newTransportErrorClient(t)}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRebootFmAction_Invoke_APIError exercises RebootFmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRebootFmAction_Invoke_APIError(t *testing.T) {
	r := &RebootFmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_reboot_fm")
}

// TestRebootFmAction_Invoke_APIErrorReadBody exercises RebootFmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRebootFmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RebootFmAction{client: newMockClientReadErrorBody(t, 501)}
	m := RebootFmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
