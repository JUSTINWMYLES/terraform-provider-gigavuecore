package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveFmInstanceAction_Invoke_Happy exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveFmInstanceAction_Invoke_Happy(t *testing.T) {
	r := &RemoveFmInstanceAction{client: newMockClientStatus(t, 204, "{}")}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveFmInstanceAction_Invoke_NilClient exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveFmInstanceAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveFmInstanceAction{}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveFmInstanceAction_Invoke_BuildError exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveFmInstanceAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveFmInstanceAction{client: newMalformedBaseURLClient(t)}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveFmInstanceAction_Invoke_SendError exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveFmInstanceAction_Invoke_SendError(t *testing.T) {
	r := &RemoveFmInstanceAction{client: newTransportErrorClient(t)}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveFmInstanceAction_Invoke_APIError exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveFmInstanceAction_Invoke_APIError(t *testing.T) {
	r := &RemoveFmInstanceAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_remove_fm_instance")
}

// TestRemoveFmInstanceAction_Invoke_APIErrorReadBody exercises RemoveFmInstanceAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveFmInstanceAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveFmInstanceAction{client: newMockClientReadErrorBody(t, 501)}
	m := RemoveFmInstanceActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
