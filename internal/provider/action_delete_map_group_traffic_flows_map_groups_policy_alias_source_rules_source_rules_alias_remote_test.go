package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_Happy exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_Happy(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_NilClient exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_BuildError exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{client: newMalformedBaseURLClient(t)}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_SendError exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_SendError(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{client: newTransportErrorClient(t)}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_APIError exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_APIError(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_map_group_traffic_flows_map_groups_policy_alias_source_rules_source_rules_alias")
}

// TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_APIErrorReadBody exercises DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
