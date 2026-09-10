package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateGsGroupParamsAction_Invoke_Happy exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateGsGroupParamsAction_Invoke_Happy(t *testing.T) {
	r := &UpdateGsGroupParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateGsGroupParamsAction_Invoke_NilClient exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateGsGroupParamsAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateGsGroupParamsAction{}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateGsGroupParamsAction_Invoke_BuildError exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateGsGroupParamsAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateGsGroupParamsAction{client: newMalformedBaseURLClient(t)}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateGsGroupParamsAction_Invoke_SendError exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateGsGroupParamsAction_Invoke_SendError(t *testing.T) {
	r := &UpdateGsGroupParamsAction{client: newTransportErrorClient(t)}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateGsGroupParamsAction_Invoke_APIError exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateGsGroupParamsAction_Invoke_APIError(t *testing.T) {
	r := &UpdateGsGroupParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_gs_group_params")
}

// TestUpdateGsGroupParamsAction_Invoke_APIErrorReadBody exercises UpdateGsGroupParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateGsGroupParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateGsGroupParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateGsGroupParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
