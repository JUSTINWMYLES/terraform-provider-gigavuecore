package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_Happy exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_Happy(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{client: newMockClientStatus(t, 202, "{}")}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_NilClient exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_NilClient(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_BuildError exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_BuildError(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{client: newMalformedBaseURLClient(t)}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_SendError exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_SendError(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{client: newTransportErrorClient(t)}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_APIError exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_APIError(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_re_apply_traffic_flow_generated_f_map")
}

// TestReApplyTrafficFlowGeneratedFMapAction_Invoke_APIErrorReadBody exercises ReApplyTrafficFlowGeneratedFMapAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReApplyTrafficFlowGeneratedFMapAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReApplyTrafficFlowGeneratedFMapAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReApplyTrafficFlowGeneratedFMapActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
