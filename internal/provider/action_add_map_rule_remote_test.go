package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapRuleAction_Invoke_Happy exercises AddMapRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapRuleAction_Invoke_NilClient exercises AddMapRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapRuleAction{}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapRuleAction_Invoke_BuildError exercises AddMapRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapRuleAction_Invoke_SendError exercises AddMapRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapRuleAction{client: newTransportErrorClient(t)}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapRuleAction_Invoke_APIError exercises AddMapRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_rule")
}

// TestAddMapRuleAction_Invoke_APIErrorReadBody exercises AddMapRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
