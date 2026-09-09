package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddSourceRulesAction_Invoke_Happy exercises AddSourceRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddSourceRulesAction_Invoke_Happy(t *testing.T) {
	r := &AddSourceRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddSourceRulesAction_Invoke_NilClient exercises AddSourceRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddSourceRulesAction_Invoke_NilClient(t *testing.T) {
	r := &AddSourceRulesAction{}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddSourceRulesAction_Invoke_BuildError exercises AddSourceRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddSourceRulesAction_Invoke_BuildError(t *testing.T) {
	r := &AddSourceRulesAction{client: newMalformedBaseURLClient(t)}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddSourceRulesAction_Invoke_SendError exercises AddSourceRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddSourceRulesAction_Invoke_SendError(t *testing.T) {
	r := &AddSourceRulesAction{client: newTransportErrorClient(t)}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddSourceRulesAction_Invoke_APIError exercises AddSourceRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddSourceRulesAction_Invoke_APIError(t *testing.T) {
	r := &AddSourceRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_source_rules")
}

// TestAddSourceRulesAction_Invoke_APIErrorReadBody exercises AddSourceRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddSourceRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddSourceRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddSourceRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
