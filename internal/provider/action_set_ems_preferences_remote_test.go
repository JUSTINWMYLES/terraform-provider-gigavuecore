package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSetEmsPreferencesAction_Invoke_Happy exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSetEmsPreferencesAction_Invoke_Happy(t *testing.T) {
	r := &SetEmsPreferencesAction{client: newMockClientStatus(t, 200, "{}")}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSetEmsPreferencesAction_Invoke_NilClient exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSetEmsPreferencesAction_Invoke_NilClient(t *testing.T) {
	r := &SetEmsPreferencesAction{}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSetEmsPreferencesAction_Invoke_BuildError exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSetEmsPreferencesAction_Invoke_BuildError(t *testing.T) {
	r := &SetEmsPreferencesAction{client: newMalformedBaseURLClient(t)}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSetEmsPreferencesAction_Invoke_SendError exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSetEmsPreferencesAction_Invoke_SendError(t *testing.T) {
	r := &SetEmsPreferencesAction{client: newTransportErrorClient(t)}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSetEmsPreferencesAction_Invoke_APIError exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSetEmsPreferencesAction_Invoke_APIError(t *testing.T) {
	r := &SetEmsPreferencesAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_set_ems_preferences")
}

// TestSetEmsPreferencesAction_Invoke_APIErrorReadBody exercises SetEmsPreferencesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSetEmsPreferencesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SetEmsPreferencesAction{client: newMockClientReadErrorBody(t, 501)}
	m := SetEmsPreferencesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
