package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_Happy exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_NilClient exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_BuildError exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_SendError exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_APIError exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_diameter_whitelist")
}

// TestRedefineGsGroupDiameterWhitelistAction_Invoke_APIErrorReadBody exercises RedefineGsGroupDiameterWhitelistAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupDiameterWhitelistAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupDiameterWhitelistAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupDiameterWhitelistActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
