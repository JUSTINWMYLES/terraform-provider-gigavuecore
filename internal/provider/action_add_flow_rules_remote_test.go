package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddFlowRulesAction_Invoke_Happy exercises AddFlowRulesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddFlowRulesAction_Invoke_Happy(t *testing.T) {
	r := &AddFlowRulesAction{client: newMockClientStatus(t, 200, "{}")}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddFlowRulesAction_Invoke_NilClient exercises AddFlowRulesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddFlowRulesAction_Invoke_NilClient(t *testing.T) {
	r := &AddFlowRulesAction{}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddFlowRulesAction_Invoke_BuildError exercises AddFlowRulesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddFlowRulesAction_Invoke_BuildError(t *testing.T) {
	r := &AddFlowRulesAction{client: newMalformedBaseURLClient(t)}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddFlowRulesAction_Invoke_SendError exercises AddFlowRulesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddFlowRulesAction_Invoke_SendError(t *testing.T) {
	r := &AddFlowRulesAction{client: newTransportErrorClient(t)}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddFlowRulesAction_Invoke_APIError exercises AddFlowRulesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddFlowRulesAction_Invoke_APIError(t *testing.T) {
	r := &AddFlowRulesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_flow_rules")
}

// TestAddFlowRulesAction_Invoke_APIErrorReadBody exercises AddFlowRulesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddFlowRulesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddFlowRulesAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddFlowRulesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
