package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_Happy exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_NilClient exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_BuildError exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_SendError exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_APIError exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_params_gtp_random_sampling")
}

// TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_APIErrorReadBody exercises RedefineGsGroupParamsGtpRandomSamplingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupParamsGtpRandomSamplingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupParamsGtpRandomSamplingAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupParamsGtpRandomSamplingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
