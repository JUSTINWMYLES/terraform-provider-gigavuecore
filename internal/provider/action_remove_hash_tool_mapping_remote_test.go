package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveHashToolMappingAction_Invoke_Happy exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveHashToolMappingAction_Invoke_Happy(t *testing.T) {
	r := &RemoveHashToolMappingAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveHashToolMappingAction_Invoke_NilClient exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveHashToolMappingAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveHashToolMappingAction{}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveHashToolMappingAction_Invoke_BuildError exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveHashToolMappingAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveHashToolMappingAction{client: newMalformedBaseURLClient(t)}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveHashToolMappingAction_Invoke_SendError exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveHashToolMappingAction_Invoke_SendError(t *testing.T) {
	r := &RemoveHashToolMappingAction{client: newTransportErrorClient(t)}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveHashToolMappingAction_Invoke_APIError exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveHashToolMappingAction_Invoke_APIError(t *testing.T) {
	r := &RemoveHashToolMappingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_hash_tool_mapping")
}

// TestRemoveHashToolMappingAction_Invoke_APIErrorReadBody exercises RemoveHashToolMappingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveHashToolMappingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveHashToolMappingAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveHashToolMappingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
