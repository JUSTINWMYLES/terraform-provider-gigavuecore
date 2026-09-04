package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_Happy exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_NilClient exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_BuildError exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_SendError exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_APIError exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_node_role_params")
}

// TestRedefineGsGroupNodeRoleParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupNodeRoleParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupNodeRoleParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupNodeRoleParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupNodeRoleParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
