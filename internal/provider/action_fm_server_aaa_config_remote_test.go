package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFmServerAaaConfigAction_Invoke_Happy exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFmServerAaaConfigAction_Invoke_Happy(t *testing.T) {
	r := &FmServerAaaConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFmServerAaaConfigAction_Invoke_NilClient exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFmServerAaaConfigAction_Invoke_NilClient(t *testing.T) {
	r := &FmServerAaaConfigAction{}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFmServerAaaConfigAction_Invoke_BuildError exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFmServerAaaConfigAction_Invoke_BuildError(t *testing.T) {
	r := &FmServerAaaConfigAction{client: newMalformedBaseURLClient(t)}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFmServerAaaConfigAction_Invoke_SendError exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFmServerAaaConfigAction_Invoke_SendError(t *testing.T) {
	r := &FmServerAaaConfigAction{client: newTransportErrorClient(t)}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFmServerAaaConfigAction_Invoke_APIError exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFmServerAaaConfigAction_Invoke_APIError(t *testing.T) {
	r := &FmServerAaaConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_fm_server_aaa_config")
}

// TestFmServerAaaConfigAction_Invoke_APIErrorReadBody exercises FmServerAaaConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFmServerAaaConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FmServerAaaConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := FmServerAaaConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
