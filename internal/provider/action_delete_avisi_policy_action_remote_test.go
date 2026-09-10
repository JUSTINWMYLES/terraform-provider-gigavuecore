package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAvisiPolicyActionAction_Invoke_Happy exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAvisiPolicyActionAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAvisiPolicyActionAction_Invoke_NilClient exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAvisiPolicyActionAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAvisiPolicyActionAction_Invoke_BuildError exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAvisiPolicyActionAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAvisiPolicyActionAction_Invoke_SendError exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAvisiPolicyActionAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{client: newTransportErrorClient(t)}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAvisiPolicyActionAction_Invoke_APIError exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAvisiPolicyActionAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_avisi_policy_action")
}

// TestDeleteAvisiPolicyActionAction_Invoke_APIErrorReadBody exercises DeleteAvisiPolicyActionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAvisiPolicyActionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAvisiPolicyActionAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAvisiPolicyActionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
