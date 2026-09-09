package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteInlineSslProfileRuleAction_Invoke_Happy exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteInlineSslProfileRuleAction_Invoke_Happy(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteInlineSslProfileRuleAction_Invoke_NilClient exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteInlineSslProfileRuleAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteInlineSslProfileRuleAction_Invoke_BuildError exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteInlineSslProfileRuleAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{client: newMalformedBaseURLClient(t)}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteInlineSslProfileRuleAction_Invoke_SendError exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteInlineSslProfileRuleAction_Invoke_SendError(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{client: newTransportErrorClient(t)}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteInlineSslProfileRuleAction_Invoke_APIError exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteInlineSslProfileRuleAction_Invoke_APIError(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_inline_ssl_profile_rule")
}

// TestDeleteInlineSslProfileRuleAction_Invoke_APIErrorReadBody exercises DeleteInlineSslProfileRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteInlineSslProfileRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteInlineSslProfileRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
