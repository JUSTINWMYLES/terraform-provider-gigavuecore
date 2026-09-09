package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestServiceStateChangeAction_Invoke_Happy exercises ServiceStateChangeAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestServiceStateChangeAction_Invoke_Happy(t *testing.T) {
	r := &ServiceStateChangeAction{client: newMockClientStatus(t, 201, "{}")}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestServiceStateChangeAction_Invoke_NilClient exercises ServiceStateChangeAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestServiceStateChangeAction_Invoke_NilClient(t *testing.T) {
	r := &ServiceStateChangeAction{}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestServiceStateChangeAction_Invoke_BuildError exercises ServiceStateChangeAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestServiceStateChangeAction_Invoke_BuildError(t *testing.T) {
	r := &ServiceStateChangeAction{client: newMalformedBaseURLClient(t)}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestServiceStateChangeAction_Invoke_SendError exercises ServiceStateChangeAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestServiceStateChangeAction_Invoke_SendError(t *testing.T) {
	r := &ServiceStateChangeAction{client: newTransportErrorClient(t)}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestServiceStateChangeAction_Invoke_APIError exercises ServiceStateChangeAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestServiceStateChangeAction_Invoke_APIError(t *testing.T) {
	r := &ServiceStateChangeAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_service_state_change")
}

// TestServiceStateChangeAction_Invoke_APIErrorReadBody exercises ServiceStateChangeAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestServiceStateChangeAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ServiceStateChangeAction{client: newMockClientReadErrorBody(t, 501)}
	m := ServiceStateChangeActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
