package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_Happy exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_NilClient exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_BuildError exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_SendError exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_APIError exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gs_group_system_params")
}

// TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGsGroupSystemParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGsGroupSystemParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGsGroupSystemParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGsGroupSystemParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
