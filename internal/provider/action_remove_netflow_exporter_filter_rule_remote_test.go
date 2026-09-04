package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveNetflowExporterFilterRuleAction_Invoke_Happy exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveNetflowExporterFilterRuleAction_Invoke_NilClient exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveNetflowExporterFilterRuleAction_Invoke_BuildError exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveNetflowExporterFilterRuleAction_Invoke_SendError exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{client: newTransportErrorClient(t)}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveNetflowExporterFilterRuleAction_Invoke_APIError exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_netflow_exporter_filter_rule")
}

// TestRemoveNetflowExporterFilterRuleAction_Invoke_APIErrorReadBody exercises RemoveNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveNetflowExporterFilterRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveNetflowExporterFilterRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
