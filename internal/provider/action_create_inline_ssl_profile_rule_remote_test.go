package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateInlineSslProfileRuleAction_Invoke_Happy exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateInlineSslProfileRuleAction_Invoke_Happy(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateInlineSslProfileRuleAction_Invoke_NilClient exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateInlineSslProfileRuleAction_Invoke_NilClient(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateInlineSslProfileRuleAction_Invoke_BuildError exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateInlineSslProfileRuleAction_Invoke_BuildError(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{client: newMalformedBaseURLClient(t)}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateInlineSslProfileRuleAction_Invoke_SendError exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateInlineSslProfileRuleAction_Invoke_SendError(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{client: newTransportErrorClient(t)}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateInlineSslProfileRuleAction_Invoke_APIError exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateInlineSslProfileRuleAction_Invoke_APIError(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_create_inline_ssl_profile_rule")
}

// TestCreateInlineSslProfileRuleAction_Invoke_APIErrorReadBody exercises CreateInlineSslProfileRuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateInlineSslProfileRuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateInlineSslProfileRuleAction{client: newMockClientReadErrorBody(t, 501)}
	m := CreateInlineSslProfileRuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
