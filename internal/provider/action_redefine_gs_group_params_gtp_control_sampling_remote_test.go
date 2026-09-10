package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_Happy exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_NilClient exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_BuildError exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_SendError exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_APIError exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_params_gtp_control_sampling")
}

// TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_APIErrorReadBody exercises RedefineGsGroupParamsGtpControlSamplingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupParamsGtpControlSamplingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupParamsGtpControlSamplingAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupParamsGtpControlSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
