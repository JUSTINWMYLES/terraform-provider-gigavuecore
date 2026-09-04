package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_Happy exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_NilClient exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_BuildError exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_SendError exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_APIError exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_load_balance_params")
}

// TestRedefineGsGroupLoadBalanceParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupLoadBalanceParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupLoadBalanceParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupLoadBalanceParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupLoadBalanceParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
