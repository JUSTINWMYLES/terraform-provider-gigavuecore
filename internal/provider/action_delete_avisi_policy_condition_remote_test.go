package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAvisiPolicyConditionAction_Invoke_Happy exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAvisiPolicyConditionAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAvisiPolicyConditionAction_Invoke_NilClient exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAvisiPolicyConditionAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAvisiPolicyConditionAction_Invoke_BuildError exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAvisiPolicyConditionAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAvisiPolicyConditionAction_Invoke_SendError exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAvisiPolicyConditionAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{client: newTransportErrorClient(t)}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAvisiPolicyConditionAction_Invoke_APIError exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAvisiPolicyConditionAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_avisi_policy_condition")
}

// TestDeleteAvisiPolicyConditionAction_Invoke_APIErrorReadBody exercises DeleteAvisiPolicyConditionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAvisiPolicyConditionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAvisiPolicyConditionAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAvisiPolicyConditionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
