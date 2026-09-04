package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteByScrollIdAction_Invoke_Happy exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteByScrollIdAction_Invoke_Happy(t *testing.T) {
	r := &DeleteByScrollIdAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteByScrollIdAction_Invoke_NilClient exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteByScrollIdAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteByScrollIdAction{}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteByScrollIdAction_Invoke_BuildError exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteByScrollIdAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteByScrollIdAction{client: newMalformedBaseURLClient(t)}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteByScrollIdAction_Invoke_SendError exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteByScrollIdAction_Invoke_SendError(t *testing.T) {
	r := &DeleteByScrollIdAction{client: newTransportErrorClient(t)}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteByScrollIdAction_Invoke_APIError exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteByScrollIdAction_Invoke_APIError(t *testing.T) {
	r := &DeleteByScrollIdAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_by_scroll_id")
}

// TestDeleteByScrollIdAction_Invoke_APIErrorReadBody exercises DeleteByScrollIdAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteByScrollIdAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteByScrollIdAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteByScrollIdActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
