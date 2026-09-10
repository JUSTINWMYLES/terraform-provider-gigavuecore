package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddNetflowExporterFilterRuleAction_Invoke_Happy exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddNetflowExporterFilterRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddNetflowExporterFilterRuleAction_Invoke_NilClient exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddNetflowExporterFilterRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddNetflowExporterFilterRuleAction_Invoke_BuildError exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddNetflowExporterFilterRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddNetflowExporterFilterRuleAction_Invoke_SendError exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddNetflowExporterFilterRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{client: newTransportErrorClient(t)}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddNetflowExporterFilterRuleAction_Invoke_APIError exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddNetflowExporterFilterRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_netflow_exporter_filter_rule")
}

// TestAddNetflowExporterFilterRuleAction_Invoke_APIErrorReadBody exercises AddNetflowExporterFilterRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddNetflowExporterFilterRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddNetflowExporterFilterRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddNetflowExporterFilterRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
