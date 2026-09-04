package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapGsRuleAction_Invoke_Happy exercises AddMapGsRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapGsRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapGsRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapGsRuleAction_Invoke_NilClient exercises AddMapGsRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapGsRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapGsRuleAction{}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapGsRuleAction_Invoke_BuildError exercises AddMapGsRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapGsRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapGsRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapGsRuleAction_Invoke_SendError exercises AddMapGsRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapGsRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapGsRuleAction{client: newTransportErrorClient(t)}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapGsRuleAction_Invoke_APIError exercises AddMapGsRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapGsRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapGsRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_gs_rule")
}

// TestAddMapGsRuleAction_Invoke_APIErrorReadBody exercises AddMapGsRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapGsRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapGsRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapGsRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
