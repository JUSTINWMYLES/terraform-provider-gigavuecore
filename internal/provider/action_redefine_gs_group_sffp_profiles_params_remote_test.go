package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_Happy exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_NilClient exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_BuildError exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_SendError exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_APIError exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sffp_profiles_params")
}

// TestRedefineGsGroupSffpProfilesParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSffpProfilesParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSffpProfilesParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSffpProfilesParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSffpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
