package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_Happy exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_NilClient exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_BuildError exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_SendError exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_APIError exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_port_throttle_sip_params")
}

// TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupPortThrottleSipParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupPortThrottleSipParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupPortThrottleSipParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupPortThrottleSipParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
