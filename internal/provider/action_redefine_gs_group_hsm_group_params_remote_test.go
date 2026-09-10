package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_Happy exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_NilClient exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_BuildError exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_SendError exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_APIError exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_hsm_group_params")
}

// TestRedefineGsGroupHsmGroupParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupHsmGroupParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupHsmGroupParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupHsmGroupParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupHsmGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
