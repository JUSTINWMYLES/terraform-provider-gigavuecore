package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_Happy exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_NilClient exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_BuildError exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_SendError exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_APIError exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gtp_flow_params")
}

// TestRedefineGsGroupGtpFlowParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGtpFlowParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGtpFlowParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGtpFlowParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGtpFlowParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
