package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestReplaceAllInlineSslProfileRulesAction_Invoke_Happy exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_Happy(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestReplaceAllInlineSslProfileRulesAction_Invoke_NilClient exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_NilClient(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestReplaceAllInlineSslProfileRulesAction_Invoke_BuildError exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_BuildError(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{client: newMalformedBaseURLClient(t)}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestReplaceAllInlineSslProfileRulesAction_Invoke_SendError exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_SendError(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{client: newTransportErrorClient(t)}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestReplaceAllInlineSslProfileRulesAction_Invoke_APIError exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_APIError(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_replace_all_inline_ssl_profile_rules")
}

// TestReplaceAllInlineSslProfileRulesAction_Invoke_APIErrorReadBody exercises ReplaceAllInlineSslProfileRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestReplaceAllInlineSslProfileRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ReplaceAllInlineSslProfileRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := ReplaceAllInlineSslProfileRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
