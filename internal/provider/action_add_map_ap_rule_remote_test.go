package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddMapApRuleAction_Invoke_Happy exercises AddMapApRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddMapApRuleAction_Invoke_Happy(t *testing.T) {
	r := &AddMapApRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddMapApRuleAction_Invoke_NilClient exercises AddMapApRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddMapApRuleAction_Invoke_NilClient(t *testing.T) {
	r := &AddMapApRuleAction{}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddMapApRuleAction_Invoke_BuildError exercises AddMapApRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddMapApRuleAction_Invoke_BuildError(t *testing.T) {
	r := &AddMapApRuleAction{client: newMalformedBaseURLClient(t)}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddMapApRuleAction_Invoke_SendError exercises AddMapApRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddMapApRuleAction_Invoke_SendError(t *testing.T) {
	r := &AddMapApRuleAction{client: newTransportErrorClient(t)}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddMapApRuleAction_Invoke_APIError exercises AddMapApRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddMapApRuleAction_Invoke_APIError(t *testing.T) {
	r := &AddMapApRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_map_ap_rule")
}

// TestAddMapApRuleAction_Invoke_APIErrorReadBody exercises AddMapApRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddMapApRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddMapApRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
