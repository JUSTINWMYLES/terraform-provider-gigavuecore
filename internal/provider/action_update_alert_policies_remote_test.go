package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateAlertPoliciesAction_Invoke_Happy exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateAlertPoliciesAction_Invoke_Happy(t *testing.T) {
	r := &UpdateAlertPoliciesAction{client: newMockClientStatus(t, 207, "{}")}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateAlertPoliciesAction_Invoke_NilClient exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateAlertPoliciesAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateAlertPoliciesAction{}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateAlertPoliciesAction_Invoke_BuildError exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateAlertPoliciesAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateAlertPoliciesAction{client: newMalformedBaseURLClient(t)}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateAlertPoliciesAction_Invoke_SendError exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateAlertPoliciesAction_Invoke_SendError(t *testing.T) {
	r := &UpdateAlertPoliciesAction{client: newTransportErrorClient(t)}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateAlertPoliciesAction_Invoke_APIError exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateAlertPoliciesAction_Invoke_APIError(t *testing.T) {
	r := &UpdateAlertPoliciesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_alert_policies")
}

// TestUpdateAlertPoliciesAction_Invoke_APIErrorReadBody exercises UpdateAlertPoliciesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateAlertPoliciesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateAlertPoliciesAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
