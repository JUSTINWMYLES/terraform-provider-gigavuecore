package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupParamsAction_Invoke_Happy exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupParamsAction_Invoke_NilClient exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupParamsAction{}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupParamsAction_Invoke_BuildError exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupParamsAction_Invoke_SendError exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupParamsAction_Invoke_APIError exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_params")
}

// TestRedefineGsGroupParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
