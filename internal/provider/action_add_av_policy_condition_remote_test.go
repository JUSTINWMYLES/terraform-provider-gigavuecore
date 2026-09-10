package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddAvPolicyConditionAction_Invoke_Happy exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddAvPolicyConditionAction_Invoke_Happy(t *testing.T) {
	r := &AddAvPolicyConditionAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddAvPolicyConditionAction_Invoke_NilClient exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddAvPolicyConditionAction_Invoke_NilClient(t *testing.T) {
	r := &AddAvPolicyConditionAction{}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddAvPolicyConditionAction_Invoke_BuildError exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddAvPolicyConditionAction_Invoke_BuildError(t *testing.T) {
	r := &AddAvPolicyConditionAction{client: newMalformedBaseURLClient(t)}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddAvPolicyConditionAction_Invoke_SendError exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddAvPolicyConditionAction_Invoke_SendError(t *testing.T) {
	r := &AddAvPolicyConditionAction{client: newTransportErrorClient(t)}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddAvPolicyConditionAction_Invoke_APIError exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddAvPolicyConditionAction_Invoke_APIError(t *testing.T) {
	r := &AddAvPolicyConditionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_av_policy_condition")
}

// TestAddAvPolicyConditionAction_Invoke_APIErrorReadBody exercises AddAvPolicyConditionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddAvPolicyConditionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddAvPolicyConditionAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddAvPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
