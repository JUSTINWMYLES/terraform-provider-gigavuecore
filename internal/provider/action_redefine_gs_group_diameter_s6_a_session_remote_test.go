package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_Happy exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_NilClient exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_BuildError exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_SendError exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_APIError exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_diameter_s6_a_session")
}

// TestRedefineGsGroupDiameterS6ASessionAction_Invoke_APIErrorReadBody exercises RedefineGsGroupDiameterS6ASessionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupDiameterS6ASessionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupDiameterS6ASessionAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupDiameterS6ASessionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
