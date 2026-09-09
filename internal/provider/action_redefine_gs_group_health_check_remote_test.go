package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupHealthCheckAction_Invoke_Happy exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupHealthCheckAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupHealthCheckAction_Invoke_NilClient exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupHealthCheckAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupHealthCheckAction_Invoke_BuildError exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupHealthCheckAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupHealthCheckAction_Invoke_SendError exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupHealthCheckAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupHealthCheckAction_Invoke_APIError exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupHealthCheckAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_health_check")
}

// TestRedefineGsGroupHealthCheckAction_Invoke_APIErrorReadBody exercises RedefineGsGroupHealthCheckAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupHealthCheckAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupHealthCheckAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupHealthCheckActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
