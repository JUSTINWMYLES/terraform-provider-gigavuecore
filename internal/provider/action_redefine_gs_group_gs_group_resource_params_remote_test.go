package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_Happy exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_NilClient exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_BuildError exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_SendError exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_APIError exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gs_group_resource_params")
}

// TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGsGroupResourceParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGsGroupResourceParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGsGroupResourceParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGsGroupResourceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
