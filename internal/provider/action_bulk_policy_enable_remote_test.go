package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestBulkPolicyEnableAction_Invoke_Happy exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestBulkPolicyEnableAction_Invoke_Happy(t *testing.T) {
	r := &BulkPolicyEnableAction{client: newMockClientStatus(t, 200, "{}")}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestBulkPolicyEnableAction_Invoke_NilClient exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestBulkPolicyEnableAction_Invoke_NilClient(t *testing.T) {
	r := &BulkPolicyEnableAction{}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestBulkPolicyEnableAction_Invoke_BuildError exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestBulkPolicyEnableAction_Invoke_BuildError(t *testing.T) {
	r := &BulkPolicyEnableAction{client: newMalformedBaseURLClient(t)}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestBulkPolicyEnableAction_Invoke_SendError exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestBulkPolicyEnableAction_Invoke_SendError(t *testing.T) {
	r := &BulkPolicyEnableAction{client: newTransportErrorClient(t)}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestBulkPolicyEnableAction_Invoke_APIError exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestBulkPolicyEnableAction_Invoke_APIError(t *testing.T) {
	r := &BulkPolicyEnableAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_bulk_policy_enable")
}

// TestBulkPolicyEnableAction_Invoke_APIErrorReadBody exercises BulkPolicyEnableAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestBulkPolicyEnableAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &BulkPolicyEnableAction{client: newMockClientReadErrorBody(t, 501)}
	m := BulkPolicyEnableActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
