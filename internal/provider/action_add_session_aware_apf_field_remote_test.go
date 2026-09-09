package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAddSessionAwareApfFieldAction_Invoke_Happy exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAddSessionAwareApfFieldAction_Invoke_Happy(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{client: newMockClientStatus(t, 201, "{}")}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAddSessionAwareApfFieldAction_Invoke_NilClient exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAddSessionAwareApfFieldAction_Invoke_NilClient(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAddSessionAwareApfFieldAction_Invoke_BuildError exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAddSessionAwareApfFieldAction_Invoke_BuildError(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{client: newMalformedBaseURLClient(t)}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAddSessionAwareApfFieldAction_Invoke_SendError exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAddSessionAwareApfFieldAction_Invoke_SendError(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{client: newTransportErrorClient(t)}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAddSessionAwareApfFieldAction_Invoke_APIError exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAddSessionAwareApfFieldAction_Invoke_APIError(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_add_session_aware_apf_field")
}

// TestAddSessionAwareApfFieldAction_Invoke_APIErrorReadBody exercises AddSessionAwareApfFieldAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAddSessionAwareApfFieldAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AddSessionAwareApfFieldAction{client: newMockClientReadErrorBody(t, 501)}
	m := AddSessionAwareApfFieldActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
