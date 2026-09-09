package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestUpdateEmailReceiversAction_Invoke_Happy exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestUpdateEmailReceiversAction_Invoke_Happy(t *testing.T) {
	r := &UpdateEmailReceiversAction{client: newMockClientStatus(t, 200, "{}")}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestUpdateEmailReceiversAction_Invoke_NilClient exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestUpdateEmailReceiversAction_Invoke_NilClient(t *testing.T) {
	r := &UpdateEmailReceiversAction{}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestUpdateEmailReceiversAction_Invoke_BuildError exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestUpdateEmailReceiversAction_Invoke_BuildError(t *testing.T) {
	r := &UpdateEmailReceiversAction{client: newMalformedBaseURLClient(t)}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestUpdateEmailReceiversAction_Invoke_SendError exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestUpdateEmailReceiversAction_Invoke_SendError(t *testing.T) {
	r := &UpdateEmailReceiversAction{client: newTransportErrorClient(t)}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestUpdateEmailReceiversAction_Invoke_APIError exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestUpdateEmailReceiversAction_Invoke_APIError(t *testing.T) {
	r := &UpdateEmailReceiversAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_update_email_receivers")
}

// TestUpdateEmailReceiversAction_Invoke_APIErrorReadBody exercises UpdateEmailReceiversAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestUpdateEmailReceiversAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &UpdateEmailReceiversAction{client: newMockClientReadErrorBody(t, 501)}
	m := UpdateEmailReceiversActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
