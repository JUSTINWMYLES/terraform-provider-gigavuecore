package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_Happy exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_NilClient exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_BuildError exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_SendError exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_APIError exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_sip_whitelist_params")
}

// TestRedefineGsGroupSipWhitelistParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupSipWhitelistParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupSipWhitelistParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupSipWhitelistParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupSipWhitelistParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
