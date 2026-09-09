package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteAllElbAction_Invoke_Happy exercises DeleteAllElbAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteAllElbAction_Invoke_Happy(t *testing.T) {
	r := &DeleteAllElbAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteAllElbAction_Invoke_NilClient exercises DeleteAllElbAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteAllElbAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteAllElbAction{}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteAllElbAction_Invoke_BuildError exercises DeleteAllElbAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteAllElbAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteAllElbAction{client: newMalformedBaseURLClient(t)}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteAllElbAction_Invoke_SendError exercises DeleteAllElbAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteAllElbAction_Invoke_SendError(t *testing.T) {
	r := &DeleteAllElbAction{client: newTransportErrorClient(t)}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteAllElbAction_Invoke_APIError exercises DeleteAllElbAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteAllElbAction_Invoke_APIError(t *testing.T) {
	r := &DeleteAllElbAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_all_elb")
}

// TestDeleteAllElbAction_Invoke_APIErrorReadBody exercises DeleteAllElbAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteAllElbAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteAllElbAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteAllElbActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
