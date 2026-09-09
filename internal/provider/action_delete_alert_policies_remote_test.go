package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAlertPoliciesAction_Invoke_Happy exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAlertPoliciesAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAlertPoliciesAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAlertPoliciesAction_Invoke_NilClient exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAlertPoliciesAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAlertPoliciesAction{}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAlertPoliciesAction_Invoke_BuildError exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAlertPoliciesAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAlertPoliciesAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAlertPoliciesAction_Invoke_SendError exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAlertPoliciesAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAlertPoliciesAction{client: newTransportErrorClient(t)}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAlertPoliciesAction_Invoke_APIError exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAlertPoliciesAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAlertPoliciesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_alert_policies")
}

// TestDeleteAlertPoliciesAction_Invoke_APIErrorReadBody exercises DeleteAlertPoliciesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAlertPoliciesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAlertPoliciesAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAlertPoliciesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
