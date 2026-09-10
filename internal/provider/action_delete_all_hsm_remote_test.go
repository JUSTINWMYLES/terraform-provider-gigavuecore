package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllHsmAction_Invoke_Happy exercises DeleteAllHsmAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllHsmAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllHsmAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllHsmAction_Invoke_NilClient exercises DeleteAllHsmAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllHsmAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllHsmAction{}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllHsmAction_Invoke_BuildError exercises DeleteAllHsmAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllHsmAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllHsmAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllHsmAction_Invoke_SendError exercises DeleteAllHsmAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllHsmAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllHsmAction{client: newTransportErrorClient(t)}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllHsmAction_Invoke_APIError exercises DeleteAllHsmAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllHsmAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllHsmAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_hsm")
}

// TestDeleteAllHsmAction_Invoke_APIErrorReadBody exercises DeleteAllHsmAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllHsmAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllHsmAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllHsmActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
