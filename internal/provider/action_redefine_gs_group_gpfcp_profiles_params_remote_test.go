package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_Happy exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_NilClient exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_BuildError exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_SendError exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_APIError exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gpfcp_profiles_params")
}

// TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGpfcpProfilesParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGpfcpProfilesParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGpfcpProfilesParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGpfcpProfilesParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
