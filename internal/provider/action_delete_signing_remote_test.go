package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSigningAction_Invoke_Happy exercises DeleteSigningAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSigningAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSigningAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSigningAction_Invoke_NilClient exercises DeleteSigningAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSigningAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSigningAction{}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSigningAction_Invoke_BuildError exercises DeleteSigningAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSigningAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSigningAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSigningAction_Invoke_SendError exercises DeleteSigningAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSigningAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSigningAction{client: newTransportErrorClient(t)}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSigningAction_Invoke_APIError exercises DeleteSigningAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSigningAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSigningAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_signing")
}

// TestDeleteSigningAction_Invoke_APIErrorReadBody exercises DeleteSigningAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSigningAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSigningAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSigningActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
