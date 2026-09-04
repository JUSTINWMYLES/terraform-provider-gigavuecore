package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_Happy exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_Happy(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_NilClient exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_BuildError exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{client: newMalformedBaseURLClient(t)}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_SendError exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_SendError(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{client: newTransportErrorClient(t)}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_APIError exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_APIError(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_all_map_flow_sample_diameter_rules")
}

// TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_APIErrorReadBody exercises RemoveAllMapFlowSampleDiameterRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveAllMapFlowSampleDiameterRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveAllMapFlowSampleDiameterRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveAllMapFlowSampleDiameterRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
