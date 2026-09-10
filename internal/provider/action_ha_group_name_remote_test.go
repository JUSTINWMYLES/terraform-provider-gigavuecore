package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestHaGroupNameAction_Invoke_Happy exercises HaGroupNameAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestHaGroupNameAction_Invoke_Happy(t *testing.T) {
	r := &HaGroupNameAction{client: newMockClientStatus(t, 201, "{}")}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestHaGroupNameAction_Invoke_NilClient exercises HaGroupNameAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestHaGroupNameAction_Invoke_NilClient(t *testing.T) {
	r := &HaGroupNameAction{}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestHaGroupNameAction_Invoke_BuildError exercises HaGroupNameAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestHaGroupNameAction_Invoke_BuildError(t *testing.T) {
	r := &HaGroupNameAction{client: newMalformedBaseURLClient(t)}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestHaGroupNameAction_Invoke_SendError exercises HaGroupNameAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestHaGroupNameAction_Invoke_SendError(t *testing.T) {
	r := &HaGroupNameAction{client: newTransportErrorClient(t)}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestHaGroupNameAction_Invoke_APIError exercises HaGroupNameAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestHaGroupNameAction_Invoke_APIError(t *testing.T) {
	r := &HaGroupNameAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_ha_group_name")
}

// TestHaGroupNameAction_Invoke_APIErrorReadBody exercises HaGroupNameAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestHaGroupNameAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &HaGroupNameAction{client: newMockClientReadErrorBody(t, 501)}
	m := HaGroupNameActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
