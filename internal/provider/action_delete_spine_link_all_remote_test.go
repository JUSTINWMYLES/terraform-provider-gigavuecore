package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteSpineLinkAllAction_Invoke_Happy exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteSpineLinkAllAction_Invoke_Happy(t *testing.T) {
	r := &DeleteSpineLinkAllAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteSpineLinkAllAction_Invoke_NilClient exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteSpineLinkAllAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteSpineLinkAllAction{}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteSpineLinkAllAction_Invoke_BuildError exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteSpineLinkAllAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteSpineLinkAllAction{client: newMalformedBaseURLClient(t)}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteSpineLinkAllAction_Invoke_SendError exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteSpineLinkAllAction_Invoke_SendError(t *testing.T) {
	r := &DeleteSpineLinkAllAction{client: newTransportErrorClient(t)}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteSpineLinkAllAction_Invoke_APIError exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteSpineLinkAllAction_Invoke_APIError(t *testing.T) {
	r := &DeleteSpineLinkAllAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_spine_link_all")
}

// TestDeleteSpineLinkAllAction_Invoke_APIErrorReadBody exercises DeleteSpineLinkAllAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteSpineLinkAllAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteSpineLinkAllAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteSpineLinkAllActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
