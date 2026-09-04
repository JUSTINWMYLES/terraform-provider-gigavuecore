package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineRadiusSystemConfigAction_Invoke_Happy exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineRadiusSystemConfigAction_Invoke_Happy(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineRadiusSystemConfigAction_Invoke_NilClient exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineRadiusSystemConfigAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineRadiusSystemConfigAction_Invoke_BuildError exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineRadiusSystemConfigAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{client: newMalformedBaseURLClient(t)}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineRadiusSystemConfigAction_Invoke_SendError exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineRadiusSystemConfigAction_Invoke_SendError(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{client: newTransportErrorClient(t)}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineRadiusSystemConfigAction_Invoke_APIError exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineRadiusSystemConfigAction_Invoke_APIError(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_radius_system_config")
}

// TestRedefineRadiusSystemConfigAction_Invoke_APIErrorReadBody exercises RedefineRadiusSystemConfigAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineRadiusSystemConfigAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineRadiusSystemConfigAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineRadiusSystemConfigActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
