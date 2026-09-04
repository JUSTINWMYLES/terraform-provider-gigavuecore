package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_Happy exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_Happy(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{client: newMockClientStatus(t, 200, "{}")}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_NilClient exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_NilClient(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_BuildError exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_BuildError(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{client: newMalformedBaseURLClient(t)}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_SendError exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_SendError(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{client: newTransportErrorClient(t)}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_APIError exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_APIError(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_redefine_gs_group_gs_group_flow_mask_params")
}

// TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_APIErrorReadBody exercises RedefineGsGroupGsGroupFlowMaskParamsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRedefineGsGroupGsGroupFlowMaskParamsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RedefineGsGroupGsGroupFlowMaskParamsAction{client: newMockClientReadErrorBody(t, 501)}
	m := RedefineGsGroupGsGroupFlowMaskParamsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
