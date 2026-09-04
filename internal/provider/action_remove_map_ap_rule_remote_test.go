package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMapApRuleAction_Invoke_Happy exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMapApRuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMapApRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMapApRuleAction_Invoke_NilClient exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMapApRuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMapApRuleAction{}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMapApRuleAction_Invoke_BuildError exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMapApRuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMapApRuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMapApRuleAction_Invoke_SendError exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMapApRuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMapApRuleAction{client: newTransportErrorClient(t)}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMapApRuleAction_Invoke_APIError exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMapApRuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMapApRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_map_ap_rule")
}

// TestRemoveMapApRuleAction_Invoke_APIErrorReadBody exercises RemoveMapApRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMapApRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMapApRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveMapApRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
