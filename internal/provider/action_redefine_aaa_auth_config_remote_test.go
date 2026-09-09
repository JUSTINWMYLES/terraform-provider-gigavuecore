package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineAaaAuthConfigAction_Invoke_Happy exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineAaaAuthConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineAaaAuthConfigAction_Invoke_NilClient exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineAaaAuthConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineAaaAuthConfigAction_Invoke_BuildError exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineAaaAuthConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineAaaAuthConfigAction_Invoke_SendError exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineAaaAuthConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{client: newTransportErrorClient(t)}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineAaaAuthConfigAction_Invoke_APIError exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineAaaAuthConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_aaa_auth_config")
}

// TestRedefineAaaAuthConfigAction_Invoke_APIErrorReadBody exercises RedefineAaaAuthConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineAaaAuthConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineAaaAuthConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
