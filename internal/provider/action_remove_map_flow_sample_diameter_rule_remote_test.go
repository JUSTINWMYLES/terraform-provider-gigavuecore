package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_Happy exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_NilClient exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_BuildError exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_SendError exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_APIError exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_flow_sample_diameter_rule")
}

// TestRemoveMapFlowSampleDiameterRuleAction_Invoke_APIErrorReadBody exercises RemoveMapFlowSampleDiameterRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapFlowSampleDiameterRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapFlowSampleDiameterRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapFlowSampleDiameterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
