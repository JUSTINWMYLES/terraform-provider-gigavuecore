package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupDedupParamsAction_Invoke_Happy exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupDedupParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupDedupParamsAction_Invoke_NilClient exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupDedupParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupDedupParamsAction_Invoke_BuildError exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupDedupParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupDedupParamsAction_Invoke_SendError exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupDedupParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupDedupParamsAction_Invoke_APIError exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupDedupParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_dedup_params")
}

// TestRedefineGsGroupDedupParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupDedupParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupDedupParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupDedupParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupDedupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
