package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_Happy exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_NilClient exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_BuildError exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_SendError exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_APIError exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_flow_sampling_params")
}

// TestRedefineGsGroupFlowSamplingParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupFlowSamplingParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupFlowSamplingParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupFlowSamplingParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupFlowSamplingParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
