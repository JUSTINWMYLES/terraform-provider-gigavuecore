package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_Happy exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_NilClient exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_BuildError exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_SendError exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_APIError exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gtp_gpfcp_delay_params")
}

// TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGtpGpfcpDelayParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGtpGpfcpDelayParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGtpGpfcpDelayParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGtpGpfcpDelayParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
