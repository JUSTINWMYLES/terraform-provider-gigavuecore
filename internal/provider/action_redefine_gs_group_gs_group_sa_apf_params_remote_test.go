package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_Happy exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_NilClient exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_BuildError exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_SendError exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_APIError exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gs_group_sa_apf_params")
}

// TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGsGroupSaApfParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGsGroupSaApfParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGsGroupSaApfParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGsGroupSaApfParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
