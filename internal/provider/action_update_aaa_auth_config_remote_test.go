package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateAaaAuthConfigAction_Invoke_Happy exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateAaaAuthConfigAction_Invoke_Happy(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateAaaAuthConfigAction_Invoke_NilClient exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateAaaAuthConfigAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateAaaAuthConfigAction_Invoke_BuildError exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateAaaAuthConfigAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{client: newMalformedBaseURLClient(t)}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateAaaAuthConfigAction_Invoke_SendError exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateAaaAuthConfigAction_Invoke_SendError(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{client: newTransportErrorClient(t)}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateAaaAuthConfigAction_Invoke_APIError exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateAaaAuthConfigAction_Invoke_APIError(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_aaa_auth_config")
}

// TestUpdateAaaAuthConfigAction_Invoke_APIErrorReadBody exercises UpdateAaaAuthConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateAaaAuthConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateAaaAuthConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateAaaAuthConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
