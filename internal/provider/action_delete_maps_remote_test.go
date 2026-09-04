package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeleteMapsAction_Invoke_Happy exercises DeleteMapsAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeleteMapsAction_Invoke_Happy(t *testing.T) {
	r := &DeleteMapsAction{client: newMockClientStatus(t, 204, "{}")}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeleteMapsAction_Invoke_NilClient exercises DeleteMapsAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeleteMapsAction_Invoke_NilClient(t *testing.T) {
	r := &DeleteMapsAction{}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeleteMapsAction_Invoke_BuildError exercises DeleteMapsAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeleteMapsAction_Invoke_BuildError(t *testing.T) {
	r := &DeleteMapsAction{client: newMalformedBaseURLClient(t)}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeleteMapsAction_Invoke_SendError exercises DeleteMapsAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeleteMapsAction_Invoke_SendError(t *testing.T) {
	r := &DeleteMapsAction{client: newTransportErrorClient(t)}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeleteMapsAction_Invoke_APIError exercises DeleteMapsAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeleteMapsAction_Invoke_APIError(t *testing.T) {
	r := &DeleteMapsAction{client: newMockClientStatus(t, 501, "{\"message\":\"boom\"}")}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking gigavuecore_delete_maps")
}

// TestDeleteMapsAction_Invoke_APIErrorReadBody exercises DeleteMapsAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeleteMapsAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeleteMapsAction{client: newMockClientReadErrorBody(t, 501)}
	m := DeleteMapsActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
