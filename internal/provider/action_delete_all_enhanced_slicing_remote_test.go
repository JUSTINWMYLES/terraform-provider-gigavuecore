package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllEnhancedSlicingAction_Invoke_Happy exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllEnhancedSlicingAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllEnhancedSlicingAction_Invoke_NilClient exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllEnhancedSlicingAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllEnhancedSlicingAction_Invoke_BuildError exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllEnhancedSlicingAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllEnhancedSlicingAction_Invoke_SendError exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllEnhancedSlicingAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{client: newTransportErrorClient(t)}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllEnhancedSlicingAction_Invoke_APIError exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllEnhancedSlicingAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_enhanced_slicing")
}

// TestDeleteAllEnhancedSlicingAction_Invoke_APIErrorReadBody exercises DeleteAllEnhancedSlicingAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllEnhancedSlicingAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllEnhancedSlicingAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllEnhancedSlicingActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
